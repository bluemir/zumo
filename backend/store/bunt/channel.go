package bunt

import (
	"github.com/bluemir/zumo/datatype"
	"github.com/sirupsen/logrus"
)

func (s *Store) FindChannels() ([]datatype.Channel, error) {
	list, err := s.list("/channel/*")
	if err != nil {
		return nil, err
	}
	result := []datatype.Channel{}
	for _, v := range list {
		channel := &datatype.Channel{}

		if err := bind(v, channel); err != nil {
			return nil, err
		}
		result = append(result, *channel)
	}
	return result, nil
}
func (s *Store) GetChannel(channelID string) (*datatype.Channel, error) {
	logrus.Debugf("[store:GetChannel] channel: %s", channelID)
	channel := &datatype.Channel{}
	err := s.get("/channel/"+channelID, channel)
	if err != nil {
		return nil, err
	}
	return channel, nil
}
func (s *Store) PutChannel(channel *datatype.Channel) (*datatype.Channel, error) {
	logrus.Debugf("[store:PutChannel] channel: %v", channel)
	err := s.put("/channel/"+channel.ID, channel)

	if err != nil {
		return nil, err
	}
	// bunt is local store, just pass it
	go s.sync.PutChannel(channel)

	return channel, nil
}
func (s *Store) DeleteChannel(channelID string) error {
	logrus.Debugf("[store:DeleteChannel] Id: %s", channelID)

	if err := s.remove("/channel/" + channelID); err != nil {
		return err
	}
	go s.sync.DeleteChannel(channelID)

	return nil
}
