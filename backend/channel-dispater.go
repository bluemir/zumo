package backend

import (
	"sync"

	"github.com/sirupsen/logrus"

	"github.com/bluemir/zumo/datatype"
)

// ChannelListener is
type ChannelListener interface {
	OnMessage(channelID string, msg datatype.Message) error
}

// ChannelDispatcher is
type ChannelDispatcher struct {
	channel *datatype.Channel

	lock *sync.RWMutex

	listeners []ChannelListener
}

// NewChannelDispatcher is
func NewChannelDispatcher(channel datatype.Channel) (*ChannelDispatcher, error) {
	logrus.Debug("init channel Dispatcher")
	dispatcher := &ChannelDispatcher{
		channel: &channel,
		lock:    &sync.RWMutex{},
	}

	return dispatcher, nil
}

// AddListener is
func (d *ChannelDispatcher) AddListener(l ChannelListener) {
	d.lock.Lock()
	defer d.lock.Unlock()

	logrus.Debugf("[ChannelDispatcher:AddListener] %x", l)

	d.listeners = append(d.listeners, l)
}

// RemoveListener is
func (d *ChannelDispatcher) RemoveListener(l ChannelListener) {
	d.lock.Lock()
	defer d.lock.Unlock()

	logrus.Debugf("[ChannelDispatcher:RemoveListener] %x", l)

	for i, listener := range d.listeners {
		if l == listener {
			d.listeners = append(d.listeners[:i], d.listeners[i+1:]...)
		}
	}

	d.listeners = append(d.listeners, l)
}

// AppendMessage is
func (d *ChannelDispatcher) AppendMessage(msg datatype.Message) {
	d.lock.RLock()
	defer d.lock.RUnlock()

	for _, listener := range d.listeners {
		listener.OnMessage(d.channel.ID, msg)
	}
}

func (d *ChannelDispatcher) isMember(username string) bool {
	d.lock.RLock()
	defer d.lock.RUnlock()

	for _, name := range d.channel.Member {
		if name == username {
			return true
		}
	}
	return false
}
