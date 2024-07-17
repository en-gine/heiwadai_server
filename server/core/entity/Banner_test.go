package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegenBanner(t *testing.T) {
	testCases := []struct {
		name     string
		imageURL string
		url      string
	}{
		{
			name:     "Valid Banner",
			imageURL: "https://example.com/image.jpg",
			url:      "https://example.com",
		},
		{
			name:     "Empty ImageURL",
			imageURL: "",
			url:      "https://example.com",
		},
		{
			name:     "Empty URL",
			imageURL: "https://example.com/image.jpg",
			url:      "",
		},
		{
			name:     "Both Empty",
			imageURL: "",
			url:      "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			banner := RegenBanner(tc.imageURL, tc.url)

			assert.NotNil(t, banner)
			assert.Equal(t, tc.imageURL, banner.ImageURL)
			assert.Equal(t, tc.url, banner.URL)
		})
	}
}
