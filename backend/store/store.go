package store

import (
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/bluemir/zumo/datatype"
)

// Store is
type Store interface {
	FindChannels() ([]datatype.Channel, error)
	GetChannel(channelID string) (*datatype.Channel, error)
	PutChannel(channel *datatype.Channel) (*datatype.Channel, error)
	DeleteChannel(ID string) error

	FindMessages(channelID string, limit int) ([]datatype.Message, error)
	PutMessage(channelID string, msg *datatype.Message) (*datatype.Message, error)

	FindUser() ([]datatype.User, error)
	GetUser(username string) (*datatype.User, error)
	PutUser(user *datatype.User) (*datatype.User, error)

	GetToken(username, hashedKey string) (*datatype.Token, error)
	PutToken(token *datatype.Token) (*datatype.Token, error)

	GetHook(hookID string) (*datatype.Hook, error)
	PutHook(*datatype.Hook) (*datatype.Hook, error)

	// genaral data, for pod or bots.
	GetData(namespace, key string, data interface{}) error
	PutData(namespace, key string, data interface{}) error
}

// Sync is
type Sync interface {
	PutChannel(channel *datatype.Channel)
	DeleteChannel(channelID string)
	PutMessage(channelID string, msg *datatype.Message)
}

var drivers = map[string]InitFunc{}

// InitFunc is
type InitFunc func(path string, sync Sync, opt map[string]string) (Store, error)

// Register is
func Register(name string, f InitFunc) {
	drivers[name] = f
}

// New is
func New(name, path string, sync Sync, opt map[string]string) (Store, error) {
	if d, ok := drivers[name]; !ok {
		logrus.Debugf("%+v", drivers)
		return nil, fmt.Errorf("'%s' is not found in store drivers", name)
	} else {
		return d(path, sync, opt)
	}
}
