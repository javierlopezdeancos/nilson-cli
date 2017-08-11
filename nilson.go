package main

import (
	"os"

	"github.com/devnilson/nilson-cli/championship"
	"github.com/urfave/cli"
)

// Following
// https://github.com/urfave/cli
// http://blog.dutchcoders.io/creating-cli-applications-in-golang/
func main() {
	app := cli.NewApp()
	app.Name = "Nilson Cli"
	app.Version = "0.0.1"
	app.Usage = "Make an explosive entrance"
	app.EnableBashCompletion = true

	app.Flags = []cli.Flag{}
	app.Flags = append(app.Flags, []cli.Flag{}...)

	app.Commands = []cli.Command{}
	app.Commands = append(app.Commands, championship.ChampionshipCommands()...)

	app.Before = func(context *cli.Context) error {
		return nil
	}
	app.BashComplete = func(c *cli.Context) {
	}

	app.Run(os.Args)
}
