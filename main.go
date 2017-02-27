package main

import (
	"bbb-future/crowler"
	"fmt"
	"os"

	cli "gopkg.in/urfave/cli.v2"
)

var runAction = func(c *cli.Context) (err error) {
	fmt.Printf("TODO")
	return
}

var crowlerAction = func(c *cli.Context) (err error) {
	err = crowler.SaveBrothersData()
	return
}

func main() {
	app := &cli.App{
		Name: "bbb-future",
		Commands: []cli.Command{
			{
				Name:        "run",
				Aliases:     []string{"r"},
				Usage:       "run the server",
				Description: "This start the server application",
				Action:      runAction,
			}, {
				Name:        "crowler",
				Aliases:     []string{"c"},
				Usage:       "run the crowler and populate db",
				Description: "This run a spider and populate the database",
				Action:      crowlerAction,
			},
		},
	}

	app.Run(os.Args)
}
