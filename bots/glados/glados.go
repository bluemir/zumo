package glados

import (
	"os/exec"
	"strings"
	"time"

	"github.com/bluemir/zumo/bots"
	"github.com/bluemir/zumo/datatype"
)

func init() {
	bots.Register("glados", New)
}

func New(c bots.Connector, s bots.DataStore) (bots.Bot, error) {
	bot := &GLaDOS{c, s}
	return bot, nil
}

type GLaDOS struct {
	bots.Connector
	bots.DataStore
}

func (bot *GLaDOS) OnMessage(channelID string, msg datatype.Message) {
	switch {
	case strings.HasPrefix(msg.Text, "ping"):
		bot.Say(channelID, "pong", nil)
	case strings.HasPrefix(msg.Text, bot.Name()+" "):
		text := msg.Text[len(bot.Name())+1:]

		switch {
		case strings.Contains(text, "report"):
			bot.Say(channelID, "booting up...", nil)
			time.AfterFunc(1*time.Second, func() {
				bot.Say(channelID, "hello there...?", nil)
			})
		case strings.Contains(text, "create bot"):
			bot.Say(channelID, "I can't do that!", nil)
		case strings.Contains(text, "check status"):
			bot.Say(channelID, "I'm fine. really...", nil)
		case strings.HasPrefix(text, "shell "):
			//
			bot.Say(channelID, "It's unabled", nil)

			if msg.Text != " __unable" {
				break
			}
			command := text[len("shell "):]
			bot.Say(channelID, "I will do that!", nil)
			cmd := exec.Command("/bin/bash", "-c", command)

			if buf, err := cmd.CombinedOutput(); err != nil {
				bot.Say(channelID, "I'm Failed. "+err.Error(), map[string]string{
					"msg": err.Error(),
				})
			} else {
				bot.Say(channelID, "Here you are.<br/>"+string(buf), map[string]string{
					"output": string(buf),
				})
			}
		default:
			bot.Say(channelID, "I don't know that command :"+text, nil)
		}
	}
}
