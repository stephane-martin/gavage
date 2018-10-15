package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/storozhukBM/verifier"
	cli "gopkg.in/urfave/cli.v1"
)

type Arguments struct {
	EsPrefix string
	EsHost   string
	EsPort   int
	Year     int
	Size     int
}

func BuildArguments(c *cli.Context) Arguments {
	return Arguments{
		EsPrefix: strings.TrimSpace(c.GlobalString("esprefix")),
		EsHost:   strings.TrimSpace(c.GlobalString("eshost")),
		EsPort:   c.GlobalInt("esport"),
		Year:     c.GlobalInt("year"),
		Size:     c.GlobalInt("size"),
	}
}

func (a Arguments) Verify() error {
	_, parseErr := url.Parse(a.GetEsURL())
	verify := verifier.New()
	verify.
		That(len(a.EsHost) > 0, "Elasticsearch host is empty").
		That(a.EsPort > 0, "Elasticsearch port is not positive").
		That(a.Year >= 1900, "The year must be greater than 1900").
		That(a.Size > 0, "The size must be strictly positive").
		That(parseErr == nil, "Elasticsearch URL can not be parsed")
	return verify.GetError()
}

func (a Arguments) GetEsURL() string {
	return fmt.Sprintf("http://%s:%d", a.EsHost, a.EsPort)
}

func BuildCliApp() *cli.App {
	app := cli.NewApp()
	app.Name = "gavage"
	app.Authors = []cli.Author{
		{
			Email: "stephane.martin@soprasteria.com",
			Name:  "Stephane Martin",
		},
	}
	app.Copyright = "Apache 2 licence"
	app.Usage = "Inject access logs to Elasticsearch"
	app.Version = Version

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "esprefix",
			Value:  "big-httplogs",
			Usage:  "Elasticsearch index prefix",
			EnvVar: "GAVAGE_PREFIX",
		},
		cli.StringFlag{
			Name:   "eshost, e",
			Value:  "127.0.0.1",
			Usage:  "Elasticsearch host",
			EnvVar: "GAVAGE_ESHOST",
		},
		cli.IntFlag{
			Name:   "esport, p",
			Value:  9200,
			Usage:  "Elasticsearch port",
			EnvVar: "GAVAGE_ESPORT",
		},
		cli.IntFlag{
			Name:   "year",
			Value:  2018,
			Usage:  "Year of generated data",
			EnvVar: "GAVAGE_YEAR",
		},
		cli.IntFlag{
			Name:   "size",
			Value:  1000,
			Usage:  "Size of generated data in GB",
			EnvVar: "GAVAGE_SIZE",
		},
		cli.BoolFlag{
			Name:   "syslog",
			Usage:  "write logs to syslog instead of stderr",
			EnvVar: "GAVAGE_SYSLOG",
		},
		cli.StringFlag{
			Name:   "loglevel",
			Value:  "info",
			Usage:  "logging level",
			EnvVar: "GAVAGE_LOGLEVEL",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:   "feed",
			Usage:  "feed ES some random logs",
			Action: MainCommand,
		},
		{
			Name:  "conf",
			Usage: "print index creation configuration",
			Action: func(c *cli.Context) error {
				fmt.Println(GetEsOpts("example", 16, 0))
				return nil
			},
		},
		{
			Name:  "one",
			Usage: "print one fake JSON document",
			Action: func(c *cli.Context) error {
				var l *LogLine
				fmt.Println(l.toJSON())
				return nil
			},
		},
	}
	return app
}

func ExitCliWithError(msg string, err error) *cli.ExitError {
	if len(msg) == 0 && err == nil {
		return nil
	}
	if len(msg) == 0 {
		return cli.NewExitError(err.Error(), 1)
	}
	if err == nil {
		return cli.NewExitError(msg, 1)
	}
	return cli.NewExitError(fmt.Sprintf("%s => %s", msg, err.Error()), 1)
}
