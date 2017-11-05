package backend

import (
	"encoding/json"
	"time"

	"github.com/bluemir/zumo/datatype"
)

func (b *backend) GetMessages(channelID string, limit int) ([]datatype.Message, error) {
	messages, err := b.store.FindMessages(channelID, limit)
	if err != nil {
		return nil, err
	}
	return messages, nil
}
func (b *backend) AppendMessage(username, channelID, text string, detail json.RawMessage) (*datatype.Message, error) {

	msg := &datatype.Message{
		Sender: username,
		Text:   text,
		Detail: detail,
		Time:   time.Now(),
	}

	if _, err := b.store.PutMessage(channelID, msg); err != nil {
		return nil, err
	}
	return msg, nil
}
