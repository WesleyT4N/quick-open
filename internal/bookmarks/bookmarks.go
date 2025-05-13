package bookmarks

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"slices"

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

type BookmarkManager struct {
	Bookmarks []Bookmark `json:"bookmarks"`
}

func initBookmarkManager(filePath string) (*BookmarkManager, error) {
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create directory: %s, %w", dir, err)
	}

	bm := &BookmarkManager{}
	bm.Save(filePath)
	return bm, nil
}

func LoadBookmarkManager(filePath string) (*BookmarkManager, error) {
	// create the file if it doesn't exist
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		bm, err := initBookmarkManager(filePath)
		if err != nil {
			return nil, fmt.Errorf("failed to initialize bookmark manager: %w", err)
		}
		return bm, nil
	}
	// laod from existing file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	bm := &BookmarkManager{
		Bookmarks: []Bookmark{},
	}
	if err := json.NewDecoder(file).Decode(bm); err != nil {
		return nil, fmt.Errorf("failed to decode JSON: %w", err)
	}

	return bm, nil
}

func (b *BookmarkManager) Save(filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()
	if err := json.NewEncoder(file).Encode(b); err != nil {
		return fmt.Errorf("failed to encode JSON: %w", err)
	}
	if err := file.Sync(); err != nil {
		return fmt.Errorf("failed to sync file: %w", err)
	}
	return nil
}

func (b *BookmarkManager) List() {
	if len(b.Bookmarks) == 0 {
		fmt.Println("No bookmarks found. Use 'qo bookmark add' to add a bookmark.")
	}

	for _, bookmark := range b.Bookmarks {
		fmt.Printf("Title: %s\nURL: %s\nAlias: %s\n\n", bookmark.Title, bookmark.URL, bookmark.Alias)
	}
}

func (b *BookmarkManager) AddBookmark(title, urlStr, alias string) (*Bookmark, error) {
	parsedURL, err := url.Parse(path.Clean(urlStr))
	if err != nil {
		return nil, fmt.Errorf("invalid URL")
	}
	newBookmark := Bookmark{
		Title: title,
		URL:   parsedURL.String(),
		Alias: alias,
	}
	b.Bookmarks = append(b.Bookmarks, newBookmark)
	return &newBookmark, nil
}

func (b *BookmarkManager) RemoveBookmark(title, urlStr, alias string) error {
	for i, bookmark := range b.Bookmarks {
		if bookmark.Title == title || bookmark.URL == urlStr || bookmark.Alias == alias {
			b.Bookmarks = slices.Delete(b.Bookmarks, i, i+1)
			return nil
		}
	}
	return fmt.Errorf("bookmark not found")
}

func (b *BookmarkManager) FindBookmark(query string) (*Bookmark, error) {
	for _, bookmark := range b.Bookmarks {
		if bookmark.Title == query || bookmark.URL == query || bookmark.Alias == query {
			return &bookmark, nil
		}
	}
	return nil, fmt.Errorf("bookmark not found")
}
