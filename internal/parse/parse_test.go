package parse

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseURLs(t *testing.T) {
	urls := ParseURLs("https://a.com, https://b.com , ,https://c.com")
	if len(urls) != 3 {
		t.Fatalf("expected 3 URLs, got %d", len(urls))
	}
}

func TestParseURLsFromFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "urls.txt")
	content := "# comment\nhttps://a.com\n\nhttps://b.com\n"
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}

	urls, err := ParseURLsFromFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if len(urls) != 2 {
		t.Fatalf("expected 2 URLs, got %d", len(urls))
	}
}
