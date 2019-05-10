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
	GetChannelsContext(context.Context, bool, ...slack.GetChannelsOption) ([]slack.Channel, error)
	GetChannelInfoContext(context.Context, string) (*slack.Channel, error)
	GetConversationInfoContext(context.Context, string, bool) (*slack.Channel, error)
	GetUserInfoContext(ctx context.Context, user string) (*slack.User, error)
	GetUserProfileContext(ctx context.Context, userId string, includeLabels bool) (*slack.UserProfile, error)
	PostMessageContext(context.Context, string, ...slack.MsgOption) (string, string, error)
}

// RTM defines which methods are required by the Slack RTM.
type RTM interface {
	ManageConnection()
	GetIncomingEvents() chan slack.RTMEvent
	NewOutgoingMessage(string, string, ...slack.RTMsgOption) *slack.OutgoingMessage
	SendMessage(*slack.OutgoingMessage)
	GetInfo() *slack.Info
}
