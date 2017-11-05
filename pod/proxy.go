package pod

import (
	"github.com/bluemir/zumo/bots"
	"github.com/bluemir/zumo/datatype"
)

type proxy struct {
	bots.Bot
	username string
}

func (p *proxy) OnMessage(channelID string, msg datatype.Message) error {
	if msg.Sender == p.username {
		return nil // skip self looping
	}
	p.Bot.OnMessage(channelID, msg)
	return nil
}
func (p *proxy) OnJoinChannel(channelID string) {
	// just skip
}
func (p *proxy) OnLeaveChannel(channelID string) {
	// just skip
}
