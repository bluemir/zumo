package backend

import (
	"fmt"

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
	if d, ok := s.channels[channel.ID]; ok {

		defer func(c datatype.Channel) {
			d.channel = &c
		}(*channel)

		// get diff of member and emit join and leave
		add, remove := diff(d.channel.Member, channel.Member)
		if len(add) == 0 && len(remove) == 0 {

			s.events.UpdateChannel <- UpdateChannelEvent{
				ChannelID: channel.ID,
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
			//s.events.EmitJoin(channel.ID, a)
		}
		for _, r := range remove {
			s.events.Leave <- LeaveEvent{
				ChannelID: channel.ID,
				UserName:  r,
			}
			//s.events.EmitLeave(channel.ID, r)
		}

	} else {
		d, err := NewChannelDispatcher(*channel)
		if err != nil {
			s.events.Error <- fmt.Errorf("[error:StoreEventHandler:PutChannel] error on create channel dispatcher")
			return
		}

		s.channels[channel.ID] = d
		s.events.CreateChannel <- CreateChannelEvent{
			Channel: *channel,
		}
		//s.events.EmitCreateChannel(channel)
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
	s.channels[channelID].AppendMessage(*msg)
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
