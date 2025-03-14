package shortener

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func ShortenURL(inputURL string) (string, error) {
	escapedURL := url.QueryEscape(inputURL)
	apiURL := fmt.Sprintf("https://tinyurl.com/api-create.php?url=%s", escapedURL)

	resp, err := http.Get(apiURL)
	if err != nil {
		return "", fmt.Errorf("failed to shorten URL: %w", err)
	}
	defer resp.Body.Close()

	shortURL, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	return string(shortURL), nil
}
