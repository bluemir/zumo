package backend

import "github.com/bluemir/zumo/datatype"

// JoinEvent is
type JoinEvent struct {
	ChannelID string
	UserName  string
}

// LeaveEvent is
type LeaveEvent struct {
	ChannelID string
	UserName  string
}

// CreateChannelEvent is
type CreateChannelEvent struct {
	Channel datatype.Channel
}

// DeleteChannelEvent is
type DeleteChannelEvent struct {
	ChannelID string
}

// UpdateChannelEvent is
type UpdateChannelEvent struct {
	Channel datatype.Channel
}

// ReceiveMessageEvent is
type ReceiveMessageEvent struct {
	ChannelID string
	Message   datatype.Message
}

// SystemEvents is
type SystemEvents struct {
	Join  chan JoinEvent
	Leave chan LeaveEvent

	CreateChannel chan CreateChannelEvent
	DeleteChannel chan DeleteChannelEvent
	UpdateChannel chan UpdateChannelEvent

	ReceiveMessage chan ReceiveMessageEvent

	Error chan error
}

func NewSystemEvents() *SystemEvents {
	return &SystemEvents{
		Join:  make(chan JoinEvent, 8),
		Leave: make(chan LeaveEvent, 8),

		CreateChannel: make(chan CreateChannelEvent, 8),
		DeleteChannel: make(chan DeleteChannelEvent, 8),
		UpdateChannel: make(chan UpdateChannelEvent, 8),

		ReceiveMessage: make(chan ReceiveMessageEvent, 16),

		Error: make(chan error),
	}
}
