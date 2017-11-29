package httpchecker

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/bluemir/zumo/bots"
	"github.com/bluemir/zumo/datatype"
)

func init() {
	bots.Register("http-checker", New)
}

func New(c bots.Connector, s bots.DataStore) (bots.Bot, error) {
	bot := &HTTPChecker{c, s, map[string]*RegisterContext{}}
	//bot.Load("schedules", data interface{})
	return bot, nil
}

type HTTPChecker struct {
	bots.Connector
	bots.DataStore

	contexts map[string]*RegisterContext
}

func (bot *HTTPChecker) OnMessage(channelId string, msg datatype.Message) {
	switch {
	case strings.HasPrefix(msg.Text, "ping"):
		bot.Say(channelId, "pong", nil)
	case strings.HasPrefix(msg.Text, "check "):
		// if url http.GET to service and return status code()
		ok, err := bot.checkURL(msg.Text[len("check "):])
		if err != nil {
			bot.Say(channelId, "Fail...: "+err.Error(), nil)
			return
		}
		if ok {
			bot.Say(channelId, "normal", nil)
		} else {
			bot.Say(channelId, "not normal", nil)
		}
	case strings.HasPrefix(msg.Text, bot.Name()+" register"):
		// TODO parse last part for fast register
		bot.contexts[channelId+":"+msg.Sender] = &RegisterContext{
			LastTime:     time.Now(),
			LastQuestion: "name",
		}
		bot.Say(channelId, "ok give me name", nil)
		// starting context
	case strings.HasPrefix(msg.Text, bot.Name()+" cancel"):
		// remove context
		delete(bot.contexts, channelId+":"+msg.Sender)
	case strings.HasPrefix(msg.Text, bot.Name()+" list"):
		// list schedules
		bot.listSchedule(channelId)
	default:
		// checkout context if exist, question to user for more information
		c, ok := bot.contexts[channelId+":"+msg.Sender]
		if !ok {
			return // just skip
		}
		// check timeout
		cur := time.Now()
		if cur.Sub(c.LastTime) > 30*time.Second {
			// Time Out!
			delete(bot.contexts, channelId+":"+msg.Sender)
			return
		}
		c.LastTime = time.Now()

		// fill data
		switch c.LastQuestion {
		case "url":
			c.URL = msg.Text
			c.LastQuestion = ""

		case "delay":
			var err error
			d, err := time.ParseDuration(msg.Text)
			if err != nil {
				bot.Say(channelId, "failed to parse duration. plz put more cafully", nil)
				return
			}
			c.Delay = d
			c.LastQuestion = ""
		case "name":
			c.Name = msg.Text
			bot.Say(channelId, "ok", nil)
		}

		// question to more information
		if c.URL == "" {
			c.LastQuestion = "url"
			bot.Say(channelId, "url? ", nil)
			return
		}
		if c.Delay == 0*time.Second {
			c.LastQuestion = "delay"
			bot.Say(channelId, "delay?", nil)
			return
		}

		bot.registerSchedule(channelId, c.Schedule)
		delete(bot.contexts, channelId+":"+msg.Sender)

		return
	}

}

type RegisterContext struct {
	Schedule

	LastTime     time.Time
	LastQuestion string
}
type Schedule struct {
	Name  string
	URL   string
	Delay time.Duration
}

func (bot *HTTPChecker) listSchedule(channelID string) {
	schdules := []Schedule{}
	if err := bot.Load(channelID, &schdules); err != nil {
		bot.Say(channelID, "opps! something wrong!", err.Error())
		return
	}
	str := ""
	for _, v := range schdules {
		str += fmt.Sprintf("[%s] %s / %s\n", v.Name, v.URL, v.Delay.String())
	}
	bot.Say(channelID, str, map[string]interface{}{
		"schdule": schdules,
	})
}
func (bot *HTTPChecker) registerSchedule(channelID string, schedule Schedule) {
	schdules := []Schedule{}
	bot.Load(channelID, &schdules)
	schdules = append(schdules, schedule)
	if err := bot.Save(channelID, schdules); err != nil {
		bot.Say(channelID, "opps! something wrong!", nil)
		return
	}
	bot.Say(channelID, "Ok, I will to that", schedule)
}

func (bot *HTTPChecker) checkURL(URL string) (bool, error) {
	c := http.Client{}
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return false, err
	}

	res, err := c.Do(req)
	if err != nil {
		return false, err
	}
	defer res.Body.Close()

	if res.StatusCode >= http.StatusOK && res.StatusCode < http.StatusMultipleChoices {
		return true, nil
	}

	return false, nil
}
