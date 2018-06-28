package slack

import (
	"github.com/nlopes/slack"
)

// Client defines which methods are required by the Slack client.
type Client interface {
	OpenIMChannel(string) (bool, bool, string, error)
	GetUserByEmail(string) (*slack.User, error)
	SetDebug(bool)
}

// RTM defines which methods are required by the Slack RTM.
type RTM interface {
	ManageConnection()
	GetIncomingEvents() chan slack.RTMEvent
	NewOutgoingMessage(string, string) *slack.OutgoingMessage
	SendMessage(*slack.OutgoingMessage)
	GetInfo() *slack.Info
}
