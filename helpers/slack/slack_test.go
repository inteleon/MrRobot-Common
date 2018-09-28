package slack

import (
	"testing"
)

func TestHighlightedSuccess(t *testing.T) {
	userID := "ABCD"

	highlightedTexts := []string{
		"<@ABCD> yoyo",
		"blabla <@ABCD> yoyo",
		"blabla yoyo <@ABCD>",
		"<@ABDC> blabla yoyo <@ABCD>",
	}

	for _, text := range highlightedTexts {
		if !Highlighted(userID, text) {
			t.Error(text, "should be highlighted")
		}
	}
}

func TestHighlightedFailure(t *testing.T) {
	userID := "ABCD"

	highlightedTexts := []string{
		"<ABCD> yoyo",
		"blabla @ABCD yoyo",
		"blabla yoyo <@ABCD",
		"<@ABDC > blabla yoyo @ABCD>",
	}

	for _, text := range highlightedTexts {
		if Highlighted(userID, text) {
			t.Error(text, "shouldn't be highlighted")
		}
	}
}
