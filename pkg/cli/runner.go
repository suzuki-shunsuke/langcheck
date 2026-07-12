// Package cli provides the command-line interface layer for langcheck.
// This package serves as the main entry point for all CLI operations,
// handling command parsing, flag processing, and routing to appropriate subcommands.
// It orchestrates the overall CLI structure using urfave/cli framework and delegates
// actual business logic to controller packages.
package cli

import (
	"context"

	"github.com/suzuki-shunsuke/langcheck/pkg/cli/check"
	"github.com/suzuki-shunsuke/slog-util/slogutil"
	"github.com/suzuki-shunsuke/urfave-cli-v3-util/urfave"
	"github.com/urfave/cli/v3"
)

func Run(ctx context.Context, logger *slogutil.Logger, env *urfave.Env) error {
	cmd := urfave.Command(env, &cli.Command{
		Name:        "langcheck",
		Usage:       "",
		Description: ``,
		Commands: []*cli.Command{
			check.New(logger),
		},
	})
	if err := cmd.Run(ctx, env.Args); err != nil {
		return err //nolint:wrapcheck
	}
	return nil
}
