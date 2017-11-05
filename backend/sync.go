package backend

import (
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/bluemir/zumo/datatype"
)

type sync struct {
	*backend
}

func (s *sync) PutChannel(channel *datatype.Channel) {
	logrus.Debug("[sync:PutChannel]")
	// find channel
	if d, ok := s.channels[channel.ID]; ok {
		// get diff of member and emit join and leave
		defer func() {
			c := *channel // copy
			d.channel = &c
		}()
		add, remove := diff(d.channel.Member, channel.Member)
		if len(add) == 0 && len(remove) == 0 {
			s.events.EmitUpdateChannel(channel)
			return
		}
		for _, a := range add {
			s.events.EmitJoin(channel.ID, a)
		}
		for _, r := range remove {
			s.events.EmitLeave(channel.ID, r)
		}
	} else {
		d, err := NewChannelDispatcher(*channel)
		if err != nil {
			s.events.EmitError(fmt.Errorf("[error:sync:PutChannel] error on create channel dispatcher"))
			return
		}

		s.channels[channel.ID] = d
		s.events.EmitCreateChannel(channel)
	}
}

// DeleteChannel is
func (s *sync) DeleteChannel(channelID string) {
	logrus.Debug("[sync:DeleteChannel]")
	s.channels[channelID].Close()
	delete(s.channels, channelID)
	s.events.EmitDeleteChannel(channelID)
}
func (s *sync) PutMessage(channelID string, msg *datatype.Message) {
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
