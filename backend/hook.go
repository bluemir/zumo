package backend

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/bluemir/zumo/backend/lib"
	"github.com/bluemir/zumo/datatype"
)

func (b *backend) CreateHook(channelID, username string) (*datatype.Hook, error) {
	logrus.Debugf("[backend:CreateHook] channelID: %s, Username: %s", channelID, username)
	return b.store.PutHook(&datatype.Hook{
		ID:        lib.RandomString(16),
		ChannelID: channelID,
		Username:  username,
	})
}
func (b *backend) DoHook(hookID, text string, detail json.RawMessage) (*datatype.Message, error) {
	logrus.Debugf("[backend:DoHook] HookID: %s, %s", hookID, text)

	if strings.Trim(text, " \r\n\t") == "" {
		return nil, errors.New("hook must have text")
	}

	hook, err := b.store.GetHook(hookID)
	if err != nil {
		return nil, errors.Wrap(err, "fail to find hook")
	}

	if detail == nil {
		detail = []byte("{}")
	}

	msg := &datatype.Message{
		Sender: hook.Username,
		Text:   text,
		Detail: detail,
		Time:   time.Now(),
	}

	if _, err := b.store.PutMessage(hook.ChannelID, msg); err != nil {
		return nil, errors.Wrap(err, "put message failed")
	}
	return msg, nil
}
