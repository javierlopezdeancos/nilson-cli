package main

import (
	"fmt"
	"os"

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

	app.Flags = []cli.Flag{
		{
			Name:  "debug",
			Usage: "Enable debug mode",
		},
	}
	app.Flags = append(app.Flags, []cli.Flag{}...)

	app.Commands = []cli.Command{
		{
			Name:    "championship",
			Aliases: []string{"c"},
			Usage:   "-- Championship --",
			Subcommands: []cli.Command{
				{
					Name:  "add",
					Usage: "add a new championship",
					Action: func(c *cli.Context) error {
						fmt.Println("Newwwww chamiponship!: ", c.Args().First())
						return nil
					},
				},
			},
		},
	}
	app.EnableBashCompletion = true

	app.Before = func(context *cli.Context) error {
		return nil
	}
	app.BashComplete = func(c *cli.Context) {
	}

	app.Run(os.Args)
}
