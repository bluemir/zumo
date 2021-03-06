package main

import (
	"flag"
	"io/ioutil"
	"math/rand"
	"time"

	"github.com/ghodss/yaml"
	"github.com/sirupsen/logrus"

	"github.com/bluemir/zumo/backend"
	"github.com/bluemir/zumo/backend/lib"
	"github.com/bluemir/zumo/pod"
	"github.com/bluemir/zumo/server"
)

const defaultConf = `
server:
  bind: localhost:4000
backend:
  store:
    driver: bunt
    endpoint: temp.db
`

// VERSION string for build number
var VERSION string

func main() {
	logrus.Infof("version: %s", VERSION)
	logrus.SetLevel(logrus.DebugLevel)
	conf, err := config()
	if err != nil {
		logrus.Error(err)
		return
	}

	b, err := backend.New(&conf.Backend)
	if err != nil {
		logrus.Error(err)
		return
	}

	key := lib.RandomString(32)
	if _, err := b.CreateToken("root", key); err != nil {
		logrus.Error(err)
	} else {
		logrus.Infof("root token: '%s'", key)
	}

	// init bots
	p, err := pod.New(b)
	if err != nil {
		logrus.Error(err)
	}

	// http connector
	if err := server.Run(b, p, &conf.Server); err != nil {
		logrus.Error(err)
		return
	}
}

// Config contain application wide configs
type Config struct {
	Backend backend.Config
	Server  server.Config
}

func config() (*Config, error) {
	conf := &Config{}
	// parse default value
	err := yaml.Unmarshal([]byte(defaultConf), conf)
	if err != nil {

		return nil, err
	}

	var confFile = flag.String("config", "", "config file path")
	flag.Parse()
	if *confFile == "" {
		return conf, nil
	}

	logrus.Debugf("config file: %s", *confFile)

	content, err := ioutil.ReadFile(*confFile)
	if err != nil {

		return nil, err
	}
	err = yaml.Unmarshal(content, conf)
	if err != nil {

		return nil, err
	}

	return conf, nil
}

func test() {
	time.Sleep(1 * time.Second)
	c := make(chan int, 4)

	for i := 1; i <= 5; i++ {
		go func(i int, ci <-chan int) {

			for {
				j := <-ci
				logrus.Infof("Start worker: %d, job: %d", i, j)
				time.Sleep(time.Duration(rand.Intn(500)+500) * time.Millisecond)
				logrus.Warnf("End   worker: %d, job: %d", i, j)
			}
		}(i, c)
	}

	go func(co chan<- int) {
		for i := 0; i < 1000; i++ {
			co <- i
		}
	}(c)

}
