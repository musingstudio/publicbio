package main

import (
	"github.com/abunchtell/publicbio"
	"github.com/urfave/cli/v2"
)

var (
	cmdServe = cli.Command{
		Name:    "serve",
		Aliases: []string{"web"},
		Usage:   "Run web application",
		Action:  serveAction,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "host",
				Usage: "Site's base URL",
			},
			&cli.IntFlag{
				Name:  "p",
				Value: 8080,
				Usage: "Port to start server on",
			},
		},
	}
)

func serveAction(c *cli.Context) error {
	cfg := &publicbio.Config{
		Host:     c.String("host"),
		Port:     c.Int("p"),
		UserFile: c.String("u"),
	}
	publicbio.Serve(cfg)
	return nil
}
