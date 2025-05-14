package bookmarks

import (
	"fmt"
	"os/exec"

	"github.com/WesleyT4N/quick-open/internal/lib"
)

type Bookmark struct {
	Title string `json:"title"`
	URL   string `json:"url"`
	Alias string `json:"alias"`
}

func (b *Bookmark) Open() error {
	openCmd, err := lib.GetOpenCommand()
	if err != nil {
		return fmt.Errorf("failed to get command to open bookmark: %w", err)
	}
	osCmd := exec.Command(openCmd, b.URL)
	if err := osCmd.Run(); err != nil {
		return fmt.Errorf("failed to open bookmark on Run: %w", err)
	}
	return nil
}
