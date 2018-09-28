package slack

import (
	"fmt"
	"regexp"
	"strings"
)

// IsCommand returns true if the message string is a direct command to the bot.
func IsCommand(userID, msg string) bool {
	return strings.HasPrefix(
		msg,
		fmt.Sprintf(
			"<@%s>",
			userID,
		),
	)
}

// ExtractCommand extracts the direct command to the bot from the message string.
func ExtractCommand(userID, msg string) string {
	re := regexp.MustCompile(
		fmt.Sprintf(
			"<@%s>\\W*",
			userID,
		),
	)

	return FormatCommand(re.ReplaceAllString(msg, ""))
}

// FormatCommand trims all whitespace from the command.
// It also makes sure the whole command is lowercase.
func FormatCommand(command string) string {
	return strings.TrimSpace(strings.ToLower(command))
}
