package slack

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	slacktypes "github.com/inteleon/MrRobot-Common/types/slack"
	"github.com/nlopes/slack"
)

// NewMessage creates and returns a new Slack Message helper.
func NewMessage(client slacktypes.Client) *Message {
	return &Message{
		client: client,
	}
}

// Message is a Slack message helper.
type Message struct {
	client slacktypes.Client
}

// Post posts a message to a Slack channel.
func (m *Message) Post(ctx context.Context, channelID, text string, options ...slack.MsgOption) (string, string, error) {
	options = append(
		options,
		slack.MsgOptionText(
			text,
			false,
		),
	)

	return m.client.PostMessageContext(
		ctx,
		channelID,
		options...,
	)
}

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
