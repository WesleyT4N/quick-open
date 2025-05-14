package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/WesleyT4N/quick-open/internal/bookmarks"
	"github.com/urfave/cli/v2"
)

var BookmarkFilePath = filepath.Join(os.Getenv("HOME"), "/.config/quick-open/bookmarks.json")

func OpenBookmark(c *cli.Context, query string) error {
	bm, err := bookmarks.LoadBookmarkManager(BookmarkFilePath)
	if err != nil {
		return cli.Exit("Failed to load bookmarks: "+err.Error(), 1)
	}

	b, err := bm.FindBookmark(query)
	if err != nil {
		return cli.Exit("Bookmark not found: "+err.Error(), 1)
	}

	fmt.Printf("Opening bookmark: %s (%s) in your browser...", b.Title, b.URL)
	if err = b.Open(); err != nil {
		return cli.Exit("Error: "+err.Error(), 1)
	}
	return nil
}

var BookmarkCmd = &cli.Command{
	Name:    "bookmark",
	Aliases: []string{"bm"},
	Usage:   "Manage bookmarks",
	Subcommands: []*cli.Command{
		{
			Name:      "add",
			Usage:     "Add a bookmark",
			Flags:     []cli.Flag{&cli.StringFlag{Name: "alias", Aliases: []string{"a"}, Usage: "Alias for the bookmark"}},
			ArgsUsage: "<title> <url>",
			Action: func(c *cli.Context) error {
				if c.NArg() != 2 {
					cli.ShowSubcommandHelpAndExit(c, 1)
				}

				bm, err := bookmarks.LoadBookmarkManager(BookmarkFilePath)
				if err != nil {
					return cli.Exit("Failed to load bookmarks: "+err.Error(), 1)
				}

				title := c.Args().Get(0)
				url := c.Args().Get(1)
				alias := c.String("alias")

				b, err := bm.AddBookmark(title, url, alias)
				if err != nil {
					return cli.Exit("Failed to add bookmark: "+err.Error(), 1)
				}

				if err := bm.Save(BookmarkFilePath); err != nil {
					return cli.Exit("Failed to save bookmark: "+err.Error(), 1)
				}
				fmt.Printf("Bookmark added: %s (%s)\n", b.Title, b.URL)
				return nil
			},
		},
		{
			Name:    "list",
			Aliases: []string{"ls"},
			Usage:   "List bookmarks",
			Action: func(c *cli.Context) error {
				bm, err := bookmarks.LoadBookmarkManager(BookmarkFilePath)
				if err != nil {
					return cli.Exit("Failed to load bookmarks: "+err.Error(), 1)
				}
				bm.List()
				return nil
			},
		},
		{
			Name:      "remove",
			Aliases:   []string{"rm"},
			Usage:     "Remove a bookmark",
			ArgsUsage: "<title | url | alias>",
			Action: func(c *cli.Context) error {
				if c.NArg() != 1 {
					cli.ShowSubcommandHelpAndExit(c, 1)
				}
				bm, err := bookmarks.LoadBookmarkManager(BookmarkFilePath)
				if err != nil {
					return cli.Exit("Failed to load bookmarks: "+err.Error(), 1)
				}
				query := c.Args().Get(0)
				removed, err := bm.RemoveBookmark(query)
				if err != nil {
					return cli.Exit("Bookmark not found: "+err.Error(), 1)
				}
				if err := bm.Save(BookmarkFilePath); err != nil {
					return cli.Exit("Failed to save bookmarks: "+err.Error(), 1)
				}
				fmt.Printf("Bookmark removed: %s (%s)\n", removed.Title, removed.URL)
				return nil
			},
		},
		{
			Name:      "open",
			Aliases:   []string{"o"},
			Usage:     "Open a bookmark",
			ArgsUsage: "<title | url | alias>",
			Action: func(c *cli.Context) error {
				if c.NArg() != 1 {
					cli.ShowSubcommandHelpAndExit(c, 1)
				}
				query := c.Args().Get(0)
				return OpenBookmark(c, query)
			},
		},
	},
}
