package championship

import (
	"fmt"

	"github.com/urfave/cli"
)

func listChampionships(c *cli.Context) error {
	fmt.Println("Runned: chamiponship:list")
	return nil
}
