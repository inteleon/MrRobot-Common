package slack

import (
	"context"
	"github.com/nlopes/slack"
)

// Client defines which methods are required by the Slack client.
type Client interface {
	OpenIMChannel(string) (bool, bool, string, error)
	GetUserByEmail(string) (*slack.User, error)
	GetUserIdentityContext(context.Context) (*slack.UserIdentityResponse, error)
	GetChannelsContext(context.Context, bool) ([]slack.Channel, error)
	GetChannelInfoContext(context.Context, string) (*slack.Channel, error)
	GetConversationInfoContext(context.Context, string, bool) (*slack.Channel, error)
	PostMessageContext(context.Context, string, string, slack.PostMessageParameters) (string, string, error)
	SetDebug(bool)
}

// RTM defines which methods are required by the Slack RTM.
type RTM interface {
	ManageConnection()
	GetIncomingEvents() chan slack.RTMEvent
	NewOutgoingMessage(string, string, ...slack.RTMsgOption) *slack.OutgoingMessage
	SendMessage(*slack.OutgoingMessage)
	GetInfo() *slack.Info
}
