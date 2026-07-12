package check

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/suzuki-shunsuke/slog-error/slogerr"
)

var (
	pattern  = regexp.MustCompile(`[\p{Han}\p{Hiragana}\p{Katakana}]`)
	errCheck = errors.New("check failed")
)

// maxLineSize is the maximum size of a single line the scanner reads.
// It is enlarged from bufio.Scanner's default so long lines (e.g. minified
// files) don't abort the scan.
const maxLineSize = 1024 * 1024

func (c *Controller) Check(logger *slog.Logger, files, texts []string) error {
	failed := false
	for _, file := range files {
		if c.checkFile(logger, file) {
			failed = true
		}
	}
	for i, text := range texts {
		if c.checkText(logger, strconv.Itoa(i), strings.NewReader(text)) {
			failed = true
		}
	}
	if failed {
		return errCheck
	}
	return nil
}

func (c *Controller) checkFile(logger *slog.Logger, file string) bool {
	f, err := os.Open(file)
	if err != nil {
		logger.Error("open a file", "file", file, "error", err)
		return true
	}
	defer f.Close()

	return c.checkText(logger, file, f)
}

func (c *Controller) checkText(logger *slog.Logger, id string, text io.Reader) bool {
	failed := false
	scanner := bufio.NewScanner(text)
	scanner.Buffer(make([]byte, 0, bufio.MaxScanTokenSize), maxLineSize)
	lineNum := 0
	for scanner.Scan() {
		lineNum++
		line := scanner.Text()
		if pattern.MatchString(line) {
			report(fmt.Sprintf("%s:%d", id, lineNum), line)
			failed = true
		}
	}
	if err := scanner.Err(); err != nil {
		slogerr.WithError(logger, err).Error("read input", "id", id)
		return true
	}
	return failed
}

// report warns about one offending line in the plain format
// "<location>\n<line>", where location is "<file>:<line number>" or
// "<text index>:<line number>".
func report(location, line string) {
	fmt.Fprintf(os.Stderr, "%s\n%s\n", location, line)
}
