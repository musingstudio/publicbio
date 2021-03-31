package main

import (
	"github.com/abunchtell/publicbio"
	"github.com/urfave/cli/v2"
	"github.com/writeas/web-core/log"
	"os"
)

func main() {
	app := &cli.App{
		Name:    "Public Bio",
		Usage:   "A public bio and link page builder.",
		Version: publicbio.FormatVersion(),
		Action:  serveAction,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "u",
				Usage: "Configuration file for single-user site",
			},
		},
	}

	app.Commands = []*cli.Command{
		&cmdServe,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
}
