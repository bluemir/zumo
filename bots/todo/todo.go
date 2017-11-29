package todo

import (
	"strings"

	"github.com/bluemir/zumo/bots"
	"github.com/bluemir/zumo/datatype"
)

func init() {
	bots.Register("todo", New)
}

func New(c bots.Connector, s bots.DataStore) (bots.Bot, error) {
	bot := &TodoBot{c, s}
	return bot, nil
}

type TodoBot struct {
	bots.Connector
	bots.DataStore
}

func (bot *TodoBot) OnMessage(channelId string, msg datatype.Message) {
	// TODO we will use machin learning or others
	switch {
	case strings.HasPrefix(msg.Text, "ping"):
		bot.Say(channelId, "pong", nil)
	case strings.HasPrefix(msg.Text, "TODO "):
		bot.Add(channelId, msg.Text[5:])
	case strings.HasPrefix(msg.Text, "DONE "):
		bot.Done(channelId, msg.Text[5:])
	case strings.HasPrefix(msg.Text, bot.Name()+" remove "):
		l := len(bot.Name() + " remove ")
		bot.Remove(channelId, msg.Text[l:]) // i know bad parsing
	case (strings.Contains(msg.Text, bot.Name()) && strings.Contains(msg.Text, "list")):
		bot.List(channelId)
	case (strings.Contains(msg.Text, bot.Name())) && strings.Contains(msg.Text, "clean"):
		bot.Cleanup(channelId)
	case msg.Text == "todo list":
		bot.List(channelId)
	}
}

func (bot *TodoBot) Add(channelId string, text string) {
	bot.Say(channelId, "ok, I will add to TODO list.", nil)

	data := &Data{}
	bot.Load(channelId, data) // ignore error

	data.Jobs = append(data.Jobs, Job{text, false})

	bot.MustSave(channelId, data)
}

func (bot *TodoBot) Done(channelId string, text string) {
	bot.do(channelId, text, func(data *Data) {
		data.check(text)
		bot.Say(channelId, "Did you do that? great!", nil)
	})
}
func (bot *TodoBot) Remove(channelId string, text string) {
	bot.do(channelId, text, func(data *Data) {
		data.remove(text)
	})
}
func (bot *TodoBot) Cleanup(channelID string) {
	data := &Data{}
	bot.MustLoad(channelID, data)

	jobs := []Job{}

	for _, job := range data.Jobs {
		if !job.IsDone {
			jobs = append(jobs, job)
		}
	}

	data.Jobs = jobs
	bot.MustSave(channelID, data)
	bot.Say(channelID, "ok jobs complete", nil)
}

func (bot *TodoBot) List(channelId string) {
	data := &Data{}
	if err := bot.Load(channelId, data); err != nil {
		bot.Say(channelId, "Sorry! I have problem! :(", map[string]string{"msg": err.Error()})
	}

	if len(data.Jobs) == 0 {
		bot.Say(channelId, "There is nothing to do!", nil)
		return
	}

	str := `TODO List`
	html := `<ul>`

	for _, job := range data.Jobs {
		if job.IsDone {
			str += "\n[x] " + job.Text
			html += `<li><input type="checkbox" checked disabled/>` + job.Text + `</li>`
		} else {
			str += "\n[ ] " + job.Text
			html += `<li><input type="checkbox" disabled/>` + job.Text + `</li>`
		}
	}
	html += "</ul>"

	bot.Say(channelId, str, map[string]string{
		"zumo.message.html": html,
	})
}

func (bot *TodoBot) do(channelId, text string, cb func(*Data)) {
	data := &Data{}
	bot.MustLoad(channelId, data)

	if count := data.find(text); count > 1 {
		bot.Say(channelId, "There is more than two matching jobs. plz say more detail", nil)
		return
	} else if count == 0 {
		bot.Say(channelId, "There is no matching jobs. :(", nil)
		return
	}

	cb(data)

	bot.MustSave(channelId, data)
}

func (bot *TodoBot) MustSave(channelId string, data interface{}) {
	if err := bot.Save(channelId, data); err != nil {
		bot.Say(channelId, "Sorry! I have problem on saving data! :(", map[string]string{"msg": err.Error()})
	}
}

func (bot *TodoBot) MustLoad(channelId string, data interface{}) {
	if err := bot.Load(channelId, data); err != nil {
		bot.Say(channelId, "Sorry! I have problem on loading data! :(", map[string]string{"msg": err.Error()})
	}
}

type Data struct {
	Jobs []Job
}

func (d *Data) find(str string) int {
	count := 0
	for _, job := range d.Jobs {
		if strings.Contains(job.Text, str) {
			count++
		}
	}
	return count
}
func (d *Data) check(str string) {
	for i, job := range d.Jobs {
		if strings.Contains(job.Text, str) {
			d.Jobs[i].IsDone = true
		}
	}
}
func (d *Data) remove(str string) {
	for i, job := range d.Jobs {
		if strings.Contains(job.Text, str) {
			d.Jobs = append(d.Jobs[:i], d.Jobs[i+1:]...)
		}
	}
}

type Job struct {
	Text   string
	IsDone bool
}
