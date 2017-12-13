package backend

import (
	"github.com/sirupsen/logrus"

	"github.com/bluemir/zumo/datatype"
)

// ChannelListener is
type ChannelListener func(channelId string, msg datatype.Message) error

// ChannelDispatcher is
type ChannelDispatcher struct {
	channel *datatype.Channel

	quit      chan struct{}
	messageQ  chan datatype.Message
	appendQ   chan ChannelListener
	listeners []ChannelListener
}

// NewChannelDispatcher is
func NewChannelDispatcher(channel datatype.Channel) (*ChannelDispatcher, error) {
	logrus.Debug("init channel Dispatcher")
	dispatcher := &ChannelDispatcher{
		channel:  &channel,
		quit:     make(chan struct{}),
		messageQ: make(chan datatype.Message),
		appendQ:  make(chan ChannelListener),
	}
	go dispatcher.runDispacter()
	return dispatcher, nil
}

func (d *ChannelDispatcher) runDispacter() {
	for {
		select {
		case msg := <-d.messageQ:
			for _, l := range d.listeners {
				l(d.channel.ID, msg)
			}
		case l := <-d.appendQ:
			logrus.Debugf("[ChannelDispatcher:%s] append to listener", d.channel.ID)
			d.listeners = append(d.listeners, l)
		case <-d.quit:
			return
		}
	}
}

// AddListener is
func (d *ChannelDispatcher) AddListener(l ChannelListener) {
	d.appendQ <- l
}

// AppendMessage is
func (d *ChannelDispatcher) AppendMessage(msg datatype.Message) {
	d.messageQ <- msg
}

// Close is
func (d *ChannelDispatcher) Close() {
	close(d.quit)
}

func (d *ChannelDispatcher) isMember(username string) bool {
	for _, name := range d.channel.Member {
		if name == username {
			return true
		}
	}
	return false
}
