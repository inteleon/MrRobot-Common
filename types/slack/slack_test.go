package slack

import (
	"reflect"
	"testing"

	"github.com/nlopes/slack"
)

func TestWeComplyWithTheSlackClient(t *testing.T) {
	clientInterfaceType := reflect.TypeOf((*Client)(nil)).Elem()
	clientImplType := reflect.TypeOf((*slack.Client)(nil))

	if !clientImplType.Implements(clientInterfaceType) {
		t.Fatal("we do not comply with the slack client interface")
	}
}
