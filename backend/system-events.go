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
	ChannelID string
}

// SystemEvents is
type SystemEvents struct {
	Join  chan JoinEvent
	Leave chan LeaveEvent

	CreateChannel chan CreateChannelEvent
	DeleteChannel chan DeleteChannelEvent
	UpdateChannel chan UpdateChannelEvent

	Error chan error
}
