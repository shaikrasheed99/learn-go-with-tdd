package mutex

import (
	"testing"

	"gotest.tools/assert"
)

func TestWebsiteChecker(t *testing.T) {
	t.Run("ShouldBeAbleToGetStatusCodesOfWebsites", func(t *testing.T) {
		websites := []string{
			"https://medium.com/@rasheed99/pagination-and-sorting-using-spring-boot-with-tdd-part-i-c6533aa57a0c",
			"https://medium.com/@rasheed99/pagination-and-sorting-using-spring-boot-with-tdd-part-ii-281958b2d6fb",
			"https://medium.com/@rasheed99/how-to-write-unit-tests-in-go-a0492e18aff2",
			"https://medium.com/@rasheed99/how-to-write-table-driven-tests-in-go-8e96ef048cca",
			"https://medium.com/@rasheed99/how-to-write-subtests-in-go-7cd9c066579d",
			"https://api.spotify.com/v1/albums",
		}

		want := "401 Unauthorized"

		got := WebsiteChecker(websites)

		assert.Equal(t, want, got["https://api.spotify.com/v1/albums"])
	})
}
