package championship

import (
	"github.com/urfave/cli"
	"fmt"
)

var chamionshipCmds = []cli.Command{
	cli.Command{
		Name: "championship:add",
		Usage: "",
		Description: `Retrieve document from the index.`,
		Action: addChampionship,
		BashComplete: func(context *cli.Context) {
			fmt.Println("Championship created")
		},
		Flags: []cli.Flag{
		},
	},
}

