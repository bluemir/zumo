package bunt

import (
	"fmt"

	"github.com/bluemir/zumo/datatype"
	"github.com/sirupsen/logrus"
)

// FindMessages is
func (s *Store) FindMessages(channelID string, limit int) ([]datatype.Message, error) {
	list, err := s.find(fmt.Sprintf("/message/%s/*", channelID), limit)
	if err != nil {
		return nil, err
	}
	result := []datatype.Message{}
	for _, v := range list {
		msg := &datatype.Message{}

		if err := bind(v, msg); err != nil {
			return nil, err
		}
		result = append(result, *msg)
	}
	return result, nil
}

// PutMessage is
func (s *Store) PutMessage(channelID string, msg *datatype.Message) (*datatype.Message, error) {
	// 만약 같은 시간에 두개가 들어오면? 그냥 무시한다.
	logrus.Debugf("[store:message:put] %s %s - %s", channelID, msg.Sender, msg.Text)
	// just pass

	err := s.put(fmt.Sprintf("/message/%s/%.16x", channelID, msg.Time.UnixNano()), msg)
	if err != nil {
		return nil, err
	}

	go s.sync.PutMessage(channelID, msg)
	return msg, nil
}
