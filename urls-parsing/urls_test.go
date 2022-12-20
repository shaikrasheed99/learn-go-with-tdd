package urlsparsing

import (
	"testing"

	"gotest.tools/assert"
)

func TestURLs(t *testing.T) {
	t.Run("ShouldBeAbleToParseTheURL", func(t *testing.T) {
		url := "https://dummywebsite.com/learn?backend=go&frontend=reactjs"

		want := ParsedURL{
			Scheme: "https",
			Host:   "dummywebsite.com",
			Path:   "/learn",
			Query: map[string]string{
				"backend":  "go",
				"frontend": "reactjs",
			},
		}

		got := Parse(url)

		assert.DeepEqual(t, want, got)
	})

	t.Run("ShouldBeAbleToConstructTheURL", func(t *testing.T) {
		parsedURL := ParsedURL{
			Scheme: "https",
			Host:   "dummywebsite.com",
			Path:   "/learn",
			Query: map[string]string{
				"backend": "go",
			},
		}

		want := "https://dummywebsite.com/learn?backend=go"

		got := Construct(parsedURL)

		assert.DeepEqual(t, want, got)
	})
}
