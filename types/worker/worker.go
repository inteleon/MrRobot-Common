package worker

import (
	"context"
	"github.com/inteleon/MrRobot-Common/types/slack"
	"github.com/julienschmidt/httprouter"
)

// Worker is the interface workers need to implement.
type Worker interface {
	Setup(context.Context, *httprouter.Router, slack.Client, slack.RTM) error
	Start(context.Context) error
	Listen(context.Context) (chan interface{}, error)
}
