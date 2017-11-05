package bots

import (
	"fmt"

	"github.com/bluemir/zumo/backend"
	"github.com/bluemir/zumo/datatype"
)

type Bot interface {
	OnMessage(channelId string, msg datatype.Message)
	// OnJoin(channelId)
	// OnLeave(channelId)

	// http.HandlerFunc
}
type Connector interface {
	Name() string

	// send message
	Say(channelId string, text string, detail interface{})
}
type DataStore = backend.DataStore

var drivers = map[string]InitFunc{}

type InitFunc func(c Connector, s DataStore) (Bot, error)

func Register(name string, f InitFunc) {
	drivers[name] = f
}

func New(name string, c Connector, s DataStore) (Bot, error) {
	if d, ok := drivers[name]; !ok {
		return nil, fmt.Errorf("'%s' is not found in bot drivers", name)
	} else {
		b, err := d(c, s)
		return b, err
	}
}

func List() []string {
	result := []string{}

	for driver, _ := range drivers {
		result = append(result, driver)
	}
	return result
}
