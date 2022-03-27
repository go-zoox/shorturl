package shorturl

import "testing"

func TestShortURL(t *testing.T) {
	longURL := "https://www.google.com"
	shortURL := ShortURL(longURL)
	if shortURL != "30ORqw" {
		t.Errorf("ShortURL(%s) = %s, want %s", longURL, shortURL, longURL)
	}
}
