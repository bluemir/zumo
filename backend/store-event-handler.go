package backend

import (
	"github.com/sirupsen/logrus"

	"github.com/bluemir/zumo/datatype"
)

// StoreEventHandler is
type StoreEventHandler struct {
	// need for transelate
	*backend
	events *SystemEvents
}

// PutChannel is callback from store
func (s *StoreEventHandler) PutChannel(channel *datatype.Channel) {
	logrus.Debug("[StoreEventHandler:PutChannel]")
	// find channel
	if channel, ok := s.channels[channel.ID]; ok {

		defer func(c datatype.Channel) {
			s.channelsLock.Lock()
			defer s.channelsLock.Unlock()
			s.channels[c.ID] = c
		}(channel)

		// get diff of member and emit join and leave
		add, remove := diff(channel.Member, channel.Member)
		if len(add) == 0 && len(remove) == 0 {
			s.events.UpdateChannel <- UpdateChannelEvent{
				Channel: channel,
			}
			// update channel info
			return
		}
		logrus.Debugf("added user: %+v, deleted user: %+v", add, remove)

		for _, a := range add {
			s.events.Join <- JoinEvent{
				ChannelID: channel.ID,
				UserName:  a,
			}
		}
		for _, r := range remove {
			s.events.Leave <- LeaveEvent{
				ChannelID: channel.ID,
				UserName:  r,
			}
		}
	} else {
		s.events.CreateChannel <- CreateChannelEvent{
			Channel: channel,
		}
	}
}

// DeleteChannel is
func (s *StoreEventHandler) DeleteChannel(channelID string) {
	logrus.Debug("[StoreEventHandler:DeleteChannel]")
	delete(s.channels, channelID)

	s.events.DeleteChannel <- DeleteChannelEvent{
		ChannelID: channelID,
	}
}

// PutMessage is
func (s *StoreEventHandler) PutMessage(channelID string, msg *datatype.Message) {
	s.events.ReceiveMessage <- ReceiveMessageEvent{
		ChannelID: channelID,
		Message:   *msg,
	}
}

func diff(old, new []string) (add, remove []string) {
	for _, o := range old {
		found := false
		for _, n := range new {
			if o == n {
				// found
				found = true
				break
			}
		}
		if !found {
			remove = append(remove, o)
		}
	}

	for _, n := range new {
		found := false
		for _, o := range old {
			if o == n {
				// found
				found = true
				break
			}
		}
		if !found {
			add = append(add, n)
		}
	}
	return
}
