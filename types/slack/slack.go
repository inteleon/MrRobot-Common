package slack

import (
	"context"
	"github.com/nlopes/slack"
)

// Client defines which methods are required by the Slack client.
type Client interface {
	OpenIMChannel(string) (bool, bool, string, error)
	GetUserByEmail(string) (*slack.User, error)
	JoinChannelContext(context.Context, string) (*slack.Channel, error)
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
