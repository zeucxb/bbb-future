package main

import (
	"bbb-future/crowler"
	"fmt"
	"os"

	cli "gopkg.in/urfave/cli.v2"
)

func main() {
	app := &cli.App{
		Name: "bbb-future",
		Commands: []cli.Command{
			{
				Name:        "run",
				Aliases:     []string{"r"},
				Usage:       "run the server",
				Description: "This start the server application",
				Action: func(c *cli.Context) error {
					fmt.Printf("TODO")
					return nil
				},
			}, {
				Name:        "crowler",
				Aliases:     []string{"c"},
				Usage:       "run the crowler and populate db",
				Description: "This run a spider and populate the database",
				Action: func(c *cli.Context) error {
					crowler.SaveBrothersData()
					return nil
				},
			},
		},
	}

	app.Run(os.Args)
}
