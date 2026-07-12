package check

import (
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func discardLogger() *slog.Logger {
	return slog.New(slog.DiscardHandler)
}

// cjk builds a string from code points so no CJK rune appears in the source
// (which gosmopolitan would otherwise flag).
func cjk(cps ...rune) string {
	return string(cps)
}

func TestController_checkText(t *testing.T) {
	t.Parallel()
	data := []struct {
		name string
		text string
		want bool
	}{
		{
			name: "english only",
			text: "all english",
			want: false,
		},
		{
			name: "empty",
			text: "",
			want: false,
		},
		{
			name: "han",
			text: "hello " + cjk(0x6F22, 0x5B57) + " world",
			want: true,
		},
		{
			name: "hiragana",
			text: cjk(0x3053, 0x3093, 0x306B, 0x3061, 0x306F),
			want: true,
		},
		{
			name: "katakana",
			text: cjk(0x30C6, 0x30B9, 0x30C8),
			want: true,
		},
		{
			name: "disallowed on later line",
			text: "hello\nworld\n" + cjk(0x3055, 0x3088, 0x3046, 0x306A, 0x3089),
			want: true,
		},
	}
	c := New()
	logger := discardLogger()
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			t.Parallel()
			if got := c.checkText(logger, "0", strings.NewReader(d.text)); got != d.want {
				t.Fatalf("checkText() = %v, want %v", got, d.want)
			}
		})
	}
}

func TestController_checkFile(t *testing.T) {
	t.Parallel()
	data := []struct {
		name    string
		content string
		missing bool
		want    bool
	}{
		{
			name:    "english only",
			content: "all english\nsecond line\n",
			want:    false,
		},
		{
			name:    "disallowed characters",
			content: "hello\n" + cjk(0x3053, 0x3093, 0x306B, 0x3061, 0x306F) + " world\n",
			want:    true,
		},
		{
			name:    "missing file",
			missing: true,
			want:    true,
		},
	}
	c := New()
	logger := discardLogger()
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			t.Parallel()
			file := filepath.Join(t.TempDir(), "input.txt")
			if !d.missing {
				if err := os.WriteFile(file, []byte(d.content), 0o600); err != nil {
					t.Fatal(err)
				}
			}
			if got := c.checkFile(logger, file); got != d.want {
				t.Fatalf("checkFile() = %v, want %v", got, d.want)
			}
		})
	}
}
