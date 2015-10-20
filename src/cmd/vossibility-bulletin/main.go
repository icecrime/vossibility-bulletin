package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/mattbaird/elastigo/api"
)

func main() {
	app := cli.NewApp()
	app.Name = "vossibility-bulleting"
	app.Usage = "generate bulletins from vossibility data"
	app.Version = "0.1.0"

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug",
			Usage: "enable debug output",
		},
		cli.StringFlag{
			Name:  "source",
			Usage: "url of the Elastic Search store",
		},
	}

	app.Before = func(c *cli.Context) error {
		if c.GlobalBool("debug") {
			log.SetLevel(log.DebugLevel)
		}
		if s := c.GlobalString("source"); s != "" {
			api.Hosts = append(api.Hosts, s)
		}
		return nil
	}
	app.Commands = []cli.Command{
		generateCommand,
	}

	app.Run(os.Args)
}
