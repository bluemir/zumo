package main

import (
	"github.com/sirupsen/logrus"

	"github.com/bluemir/zumo/backend"
	"github.com/bluemir/zumo/backend/lib"
	"github.com/bluemir/zumo/pod"
	"github.com/bluemir/zumo/server"
)

var VERSION string

func main() {
	logrus.Infof("version: %s", VERSION)

	logrus.SetLevel(logrus.DebugLevel)

	bConf := &backend.Config{}
	bConf.Store.Driver = "bunt"
	bConf.Store.Endpoint = "temp.db"

	b, err := backend.New(bConf)
	if err != nil {
		logrus.Error(err)
		return
	}

	key := lib.RandomString(32)
	if _, err := b.CreateToken("root", key); err != nil {
		logrus.Error(err)
		return
	} else {
		logrus.Infof("root token: '%s'", key)
	}

	// init bots
	p, err := pod.New(b)
	if err != nil {
		logrus.Error(err)
	}

	// http connector
	if err := server.Run(b, p); err != nil {
		logrus.Error(err)
		return
	}
}
