package backend

import (
	"github.com/sirupsen/logrus"

	"github.com/bluemir/zumo/datatype"
)

// EventListener is
type EventListener interface {
	OnCreateChannel(channel *datatype.Channel)
	OnUpdateChannel(channel *datatype.Channel)
	OnDeleteChannel(channelID string)
	OnJoin(channelID, username string)
	OnLeave(channelID, username string)
}

func (b *backend) AddEventListener(el EventListener) {
	b.events.AddListener(el)
}

type emmiter struct {
	listeners []EventListener
}

func newEmmiter() (*emmiter, error) {
	return &emmiter{}, nil
}

func (e *emmiter) AddListener(el EventListener) {
	e.listeners = append(e.listeners, el)
}

func (e *emmiter) each(do func(l EventListener)) {
	for _, l := range e.listeners {
		do(l)
	}
}
func (e *emmiter) EmitError(err error) {
	logrus.Warn(err)
}
func (e *emmiter) EmitCreateChannel(channel *datatype.Channel) {
	logrus.Debugf("[emmiter:EmitCreateChannel] %+v", channel)
	go e.each(func(l EventListener) { l.OnCreateChannel(channel) })
}
func (e *emmiter) EmitUpdateChannel(channel *datatype.Channel) {
	logrus.Debugf("[emmiter:EmitUpdateChannel] %+v", channel)
	go e.each(func(l EventListener) { l.OnUpdateChannel(channel) })
}
func (e *emmiter) EmitDeleteChannel(channelID string) {
	logrus.Debugf("[emmiter:EmitDeleteChannel] Id: %s", channelID)
	go e.each(func(l EventListener) { l.OnDeleteChannel(channelID) })
}
func (e *emmiter) EmitJoin(channelID, username string) {
	logrus.Debugf("[emmiter:EmitJoin] channelID: %s, username: %s", channelID, username)
	go e.each(func(l EventListener) { l.OnJoin(channelID, username) })
}
func (e *emmiter) EmitLeave(channelID, username string) {
	logrus.Debugf("[emmiter:EmitLeave] channelID: %s, username: %s", channelID, username)
	go e.each(func(l EventListener) { l.OnLeave(channelID, username) })
}
