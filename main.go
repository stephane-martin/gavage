package main

import (
	"os"
	"strings"

	cli "gopkg.in/urfave/cli.v1"
)

var Version string

func main() {
	BuildCliApp().Run(os.Args)
}

func MainCommand(c *cli.Context) error {
	args := BuildArguments(c)
	loglevel := strings.TrimSpace(c.GlobalString("loglevel"))
	toSyslog := c.GlobalBool("syslog")

	err := args.Verify()
	if err != nil {
		return ExitCliWithError("Arguments validation failed", err)
	}

	if len(loglevel) == 0 {
		loglevel = "info"
	}
	logger := getLogger(loglevel, toSyslog)
	logger.Debug(
		"Arguments",
		"eshost", args.EsHost,
		"esport", args.EsPort,
	)

	return nil
}
