package main

import (
	"github.com/suzuki-shunsuke/langcheck/pkg/cli"
	"github.com/suzuki-shunsuke/urfave-cli-v3-util/urfave"
)

var version = ""

func main() {
	urfave.Main("langcheck", version, cli.Run)
}
