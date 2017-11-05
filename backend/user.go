package backend

import (
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/bluemir/zumo/datatype"
)

func (b *backend) CreateUser(username string, labels map[string]string) (*datatype.User, error) {
	logrus.Debugf("[backend:CreateUser] username: '%s', labels: %+v", username, labels)
	user := &datatype.User{
		Name:   username,
		Labels: labels,
	}
	if len(user.Name) < 4 {
		return nil, fmt.Errorf("username is too short: %s", username)
	}

	if u, err := b.store.GetUser(user.Name); err != nil {
		return nil, err
	} else if u != nil {
		return nil, fmt.Errorf("user already exist")
	}

	user, err := b.store.PutUser(user)
	if err != nil {
		return nil, err
	}

	logrus.Debugf("[backend:CreateUser] '%s' created!", username, labels)
	return user, nil
}

func (b *backend) GetUser(username string) (*datatype.User, error) {
	return b.store.GetUser(username)
}
func (b *backend) ListUsers() ([]datatype.User, error) {
	return b.store.FindUser()
}
func (b *backend) Join(channelID, username string) error {
	// b.channels[channelID].Member.append(username)
	channel, err := b.store.GetChannel(channelID)
	if err != nil {
		return err
	}
	// TODO  duplicate check
	channel.Member = append(channel.Member, username)

	_, err = b.store.PutChannel(channel)
	if err != nil {
		return err
	}

	return nil
}
func (b *backend) Leave(channelID, username string) error {
	return nil
}
func (b *backend) JoinnedChannel(username string) ([]datatype.Channel, error) {
	result := []datatype.Channel{}

	// find in in-memory data
	for _, d := range b.channels {
		for _, m := range d.channel.Member {
			if m == username {
				result = append(result, *d.channel)
			}
		}
	}

	return result, nil
}
