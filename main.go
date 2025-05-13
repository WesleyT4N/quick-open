package main

import (
	"log"
	"os"

	"github.com/WesleyT4N/quick-open/cmd"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "qo",
		Usage: "Quicly open anything from your command line",
		Commands: []*cli.Command{
			cmd.BookmarkCmd,
		},
		ArgsUsage: "<alias-of-thing-to-open",
		Action: func(c *cli.Context) error {
			if c.NArg() != 1 {
				cli.ShowAppHelp(c)
				return nil
			}
			alias := c.Args().Get(0)
			err := c.App.Run([]string{c.App.Name, "bookmark", "open", alias})
			if err != nil {
				log.Fatalf("Failed to run command: %v", err)
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
