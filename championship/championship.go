package championship

import (
	"github.com/urfave/cli"
)

var championshipCommands = []cli.Command{
	{
		Name:        "championship:add",
		Aliases:     []string{"c++", "championship-add", "c-add"},
		Usage:       "",
		Description: `Creates a championship`,
		Action:      addChampionship,
		Flags:       []cli.Flag{},
	},
	{
		Name:        "championship:list",
		Aliases:     []string{"c?", "championship-list", "c-list"},
		Usage:       "",
		Description: `Lists available championships`,
		Action:      listChampionships,
		Flags:       []cli.Flag{},
	},
}

func ChampionshipCommands() []cli.Command {
	return championshipCommands
}
