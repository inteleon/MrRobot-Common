package slack

import (
	"fmt"
	"strings"
)

func Highlighted(userID, msg string) bool {
	return strings.Contains(
		msg,
		fmt.Sprintf(
			"<@%s>",
			userID,
		),
	)
}
