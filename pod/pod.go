package pod

import (
	"github.com/sirupsen/logrus"

	"github.com/bluemir/zumo/backend"
	"github.com/bluemir/zumo/bots"
	"github.com/bluemir/zumo/datatype"
)

type Pod interface {
	CreateBot(name string, driver string) error
}

func New(b backend.Backend) (Pod, error) {

	p := &pod{b}
	return p, p.InitBots()
}

type pod struct {
	backend.Backend
}

func (p *pod) InitBots() error {
	users, err := p.ListUsers()
	if err != nil {
		return err
	}

	for _, user := range users {
		if driver, ok := user.Labels["zumo.bot.driver"]; ok {
			logrus.Debugf("[pod:InitBots] '%s' found. driver: '%s'", user.Name, driver)

			err := p.initBot(driver, user) // must use copy of datatype.user
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (p *pod) CreateBot(name string, driver string) error {
	user, err := p.CreateUser(name, map[string]string{
		"zumo.bot.driver": driver,
	})
	if err != nil {
		return err
	}
	return p.initBot(driver, *user)
}

func (p *pod) initBot(driver string, user datatype.User) error {
	store := p.Backend.RequestDataStore(driver + "/" + user.Name)

	bot, err := bots.New(driver, p.NewConnector(user), store)
	if err != nil {
		return err
	}

	p.RegisterUserAgent(user.Name, &proxy{bot, user.Name})
	return nil
}
