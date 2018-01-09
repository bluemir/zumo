package backend

import (
	"github.com/bluemir/zumo/datatype"
)

// UserAgent mapping to client.
type UserAgent interface {
	OnMessage(channelID string, msg datatype.Message) error
	OnJoinChannel(channelID string)
	OnLeaveChannel(channelID string)

	// ReadEvent() <-chan interface{}
}

// SystemAgent is
type SystemAgent interface {
	OnCreateChannel(channel datatype.Channel) error
	OnDeleteChannel(channelID string)
}
