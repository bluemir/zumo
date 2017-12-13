package backend

import (
	"encoding/json"

	"github.com/bluemir/zumo/backend/store"
	"github.com/bluemir/zumo/datatype"
)

// Config is
type Config struct {
	Store struct {
		Driver   string
		Endpoint string
	}
}

// Backend is
type Backend interface {
	GetChannels() ([]datatype.Channel, error)
	CreateChannel(name string, labels map[string]string) (*datatype.Channel, error)
	DeleteChannel(channelID string) error

	GetMessages(channelID string, limit int) ([]datatype.Message, error)
	AppendMessage(username, channelID, text string, detail json.RawMessage) (*datatype.Message, error)

	CreateUser(username string, labels map[string]string) (*datatype.User, error)
	GetUser(username string) (*datatype.User, error)
	ListUsers() ([]datatype.User, error)
	JoinnedChannel(username string) ([]datatype.Channel, error)

	CreateToken(username, unhashedKey string) (*datatype.Token, error)
	Token(tokenString string) (*datatype.Token, error)

	CreateHook(channelID, username string) (*datatype.Hook, error)
	DoHook(hookID, text string, detail json.RawMessage) (*datatype.Message, error)

	Join(channeID, username string) error
	Leave(channelID, username string) error

	RequestDataStore(namespace string) DataStore

	RegisterUserAgent(username string, ua UserAgent) error
	// UnregisterUserAgent
}

// seprate backend
type backend struct {
	store store.Store

	channels         map[string]*ChannelDispatcher
	userAgentManager *UserAgentManager
}

// New is
func New(conf *Config) (Backend, error) {
	b := &backend{}

	// will be queued later
	events := NewSystemEvents()

	// store
	if store, err := store.New(conf.Store.Driver, conf.Store.Endpoint, &StoreEventHandler{b, events}, nil); err != nil {
		return nil, err
	} else {
		b.store = store
	}

	// channel dispatcher
	b.channels = map[string]*ChannelDispatcher{}
	if channels, err := b.store.FindChannels(); err != nil {
		return nil, err
	} else {
		for _, channel := range channels {
			dispacter, err := NewChannelDispatcher(channel)
			if err != nil {
				return nil, err
			}
			b.channels[channel.ID] = dispacter
		}
	}

	b.userAgentManager = NewUserAgentManager(b)

	// start main loop
	go b.runDispatcher(events)

	return b, nil
}
