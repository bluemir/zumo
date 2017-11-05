package bunt

import (
	"fmt"

	"github.com/bluemir/zumo/datatype"
	"github.com/sirupsen/logrus"
)

func (s *Store) FindMessages(channelId string, limit int) ([]datatype.Message, error) {
	list, err := s.find(fmt.Sprintf("/message/%s/*", channelId), limit)
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
func (s *Store) PutMessage(channelId string, msg *datatype.Message) (*datatype.Message, error) {
	// 만약 같은 시간에 두개가 들어오면?
	logrus.Debugf("[store:message:put] %s %+v", channelId, msg)
	// just pass

	err := s.put(fmt.Sprintf("/message/%s/%.16x", channelId, msg.Time.UnixNano()), msg)
	if err != nil {
		return nil, err
	}

	go s.sync.PutMessage(channelId, msg)
	return msg, nil
}
