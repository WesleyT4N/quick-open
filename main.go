package main

import (
	"log"
	"os"

	"github.com/WesleyT4N/quick-open/cmd"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:                 "qo",
		EnableBashCompletion: true,
		Usage:                "Quicly open anything from your command line",
		Commands: []*cli.Command{
			cmd.BookmarkCmd,
		},
		UsageText: `qo [global options] [command] [options]
qo <title|alias>`,
		Action: func(c *cli.Context) error {
			if c.NArg() != 1 {
				cli.ShowAppHelp(c)
				return nil
			}
			query := c.Args().Get(0)
			return cmd.OpenBookmark(c, query)
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
