package pod

import (
	"encoding/json"

	"github.com/sirupsen/logrus"

	"github.com/bluemir/zumo/bots"
	"github.com/bluemir/zumo/datatype"
)

type connector struct {
	user *datatype.User
	*pod
}

func (p *pod) NewConnector(user datatype.User) bots.Connector {
	return &connector{&user, p}
}

func (c *connector) Name() string {
	return c.user.Name
}

// Say is
func (c *connector) Say(channelID string, text string, detail interface{}) {
	if detail == nil {
		detail = map[string]string{}
	}
	buf, err := json.Marshal(detail)
	if err != nil {
		logrus.Warnf("[bot:%s:say] %s", c.user.Name, err.Error())
	}
	c.AppendMessage(c.user.Name, channelID, text, json.RawMessage(buf))
}
