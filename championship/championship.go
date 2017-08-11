package championship

import (
	"github.com/urfave/cli"
)

var championshipCommands = []cli.Command{
	{
		Name:        "championship:add",
		Aliases:     []string{"c++", "championship-add", "c-add"},
		Usage:       "",
		Description: `Retrieve document from the index.`,
		Action:      addChampionship,
		Flags:       []cli.Flag{},
	},
}

func ChampionshipCommands() []cli.Command {
	return championshipCommands
}
