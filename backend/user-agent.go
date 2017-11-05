package backend

import (
	"github.com/sirupsen/logrus"

	"github.com/bluemir/zumo/datatype"
)

// UserAgent mapping to client.
type UserAgent interface {
	OnMessage(channelID string, msg datatype.Message) error
	OnJoinChannel(channelID string)
	OnLeaveChannel(channelID string)
}

func (b *backend) RegisterUserAgent(username string, ua UserAgent) error {

	// QUESTION consider seprate interface and metdhod
	a := &agent{ua, username, b}
	b.agents = append(b.agents, a) // later we must AddListener when channel event

	for _, d := range b.channels {
		if d.isMember(username) {
			d.AddListener(ua.OnMessage)
		}
	}

	b.events.AddListener(a) // TODO remove userAgent from agent list
	return nil
}

// mapping to one client. it transelate server event to client event
type agent struct {
	UserAgent
	username string
	backend  *backend
}

func (a *agent) OnCreateChannel(channel *datatype.Channel) {

}
func (a *agent) OnDeleteChannel(channelID string) {

}
func (a *agent) OnUpdateChannel(channel *datatype.Channel) {

}

func (a *agent) OnJoin(channelID string, username string) {
	logrus.Debugf("[Agent:OnJoin] %s, %s", channelID, username)
	if username == a.username {
		a.backend.channels[channelID].AddListener(a.UserAgent.OnMessage)
		a.UserAgent.OnJoinChannel(channelID)
	}
}
func (a *agent) OnLeave(channelID string, username string) {
	if username == a.username {
		a.UserAgent.OnLeaveChannel(channelID)
	}
}
