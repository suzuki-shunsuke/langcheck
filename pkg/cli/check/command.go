package check

import (
	"context"

	"github.com/suzuki-shunsuke/langcheck/pkg/controller/check"
	"github.com/suzuki-shunsuke/slog-util/slogutil"
	"github.com/urfave/cli/v3"
)

type Args struct {
	Files []string
	Texts []string
}

func New(logger *slogutil.Logger) *cli.Command {
	args := &Args{}
	return &cli.Command{
		Name:        "check",
		Usage:       "",
		Description: ``,
		Action: func(ctx context.Context, _ *cli.Command) error {
			return action(ctx, logger, args)
		},
		Flags: []cli.Flag{
			&cli.StringSliceFlag{
				Name:        "text",
				Usage:       "text to check",
				Destination: &args.Texts,
			},
		},
		Arguments: []cli.Argument{
			&cli.StringArgs{
				Name:        "files",
				Max:         -1,
				Destination: &args.Files,
			},
		},
	}
}

func action(_ context.Context, logger *slogutil.Logger, args *Args) error {
	ctrl := check.New()
	return ctrl.Check(logger.Logger, args.Files, args.Texts) //nolint:wrapcheck
}
