package main

import (
	"html/template"
	"io/ioutil"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

var generateCommand = cli.Command{
	Name:   "generate",
	Usage:  "generate a bulletin",
	Action: doGenerateCommand,
}

func doGenerateCommand(c *cli.Context) {
	if len(c.Args()) != 1 {
		log.Fatal("generate requires one argument")
	}

	f := c.Args()[0]
	b, err := ioutil.ReadFile(f)
	if err != nil {
		log.Fatalf("failed to read template file: %v", err)
	}

	tmpl, err := template.New(f).Funcs(templateFuncs).Parse(string(b))
	if err != nil {
		log.Fatalf("failed to parse template file: %v", err)
	}

	if err := tmpl.Execute(os.Stdout, nil); err != nil {
		log.Fatalf("error evaluating template: %v", err)
	}
}
