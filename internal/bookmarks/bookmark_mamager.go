package bookmarks

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"slices"
)

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

	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	if _, err := file.Write(jsonData); err != nil {
		return fmt.Errorf("failed to write JSON to file: %w", err)
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

func (b *BookmarkManager) bookmarksByTitle() map[string]Bookmark {
	bookmarksByTitle := map[string]Bookmark{}
	for _, bookmark := range b.Bookmarks {
		bookmarksByTitle[bookmark.Title] = bookmark
	}
	return bookmarksByTitle
}

func (b *BookmarkManager) bookmarksByURL() map[string]Bookmark {
	bookmarksByURL := map[string]Bookmark{}
	for _, bookmark := range b.Bookmarks {
		bookmarksByURL[bookmark.URL] = bookmark
	}
	return bookmarksByURL
}

func (b *BookmarkManager) AddBookmark(title, urlStr, alias string) (*Bookmark, error) {
	parsedURL, err := url.Parse(path.Clean(urlStr))
	if err != nil {
		return nil, fmt.Errorf("invalid URL")
	}
	urlStr = parsedURL.String()

	bookmarksByTitle := b.bookmarksByTitle()
	if _, exists := bookmarksByTitle[title]; exists {
		return nil, fmt.Errorf("bookmark with title '%s' already exists", title)
	}
	bookmarksByURL := b.bookmarksByURL()
	if _, exists := bookmarksByURL[urlStr]; exists {
		return nil, fmt.Errorf("bookmark with URL '%s' already exists", urlStr)
	}

	newBookmark := Bookmark{
		Title: title,
		URL:   urlStr,
		Alias: alias,
	}
	b.Bookmarks = append(b.Bookmarks, newBookmark)
	return &newBookmark, nil
}

func (b *BookmarkManager) RemoveBookmark(query string) (*Bookmark, error) {
	for i, bookmark := range b.Bookmarks {
		if bookmark.Title == query || bookmark.URL == query || bookmark.Alias == query {
			removed := &bookmark
			b.Bookmarks = slices.Delete(b.Bookmarks, i, i+1)
			return removed, nil
		}
	}
	return nil, fmt.Errorf("bookmark not found")
}

func (b *BookmarkManager) FindBookmark(query string) (*Bookmark, error) {
	for _, bookmark := range b.Bookmarks {
		if bookmark.Title == query || bookmark.URL == query || bookmark.Alias == query {
			return &bookmark, nil
		}
	}
	return nil, fmt.Errorf("bookmark not found")
}
