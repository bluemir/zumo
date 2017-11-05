package backend

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"

	"github.com/bluemir/zumo/datatype"
)

func (b *backend) GetChannels() ([]datatype.Channel, error) {
	return b.store.FindChannels()
}
func (b *backend) CreateChannel(name string, labels map[string]string) (*datatype.Channel, error) {
	logrus.Debugf("[backend:CreateChannel] name: %s, labels: %v", name, labels)
	name = strings.Trim(name, " \t\r\n")
	//validate
	if len(name) < 4 {
		return nil, fmt.Errorf("Too short channel name")
	}

	if labels == nil {
		labels = map[string]string{}
	}

	id := uuid.New()
	channels, err := b.store.FindChannels()
	if err != nil {
		return nil, err
	}
	for _, channel := range channels {
		if channel.Name == name {
			return nil, fmt.Errorf("channel '%s' already exist", name)
		}
	}
	channel, err := b.store.PutChannel(&datatype.Channel{
		ID:     id.String(),
		Name:   name,
		Labels: labels,
	})
	if err != nil {
		return nil, err
	}

	return channel, nil
}
func (b *backend) DeleteChannel(channelID string) error {
	return b.store.DeleteChannel(channelID)
}
