package etcd

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/bluemir/zumo/backend/store"
	"github.com/bluemir/zumo/datatype"
	"github.com/coreos/etcd/clientv3"
	"github.com/sirupsen/logrus"
)

func init() {
	store.Register("etcd", New)
}

// New is
func New(path string, sync store.Sync, opt map[string]string) (store.Store, error) {
	d, err := time.ParseDuration(opt["time-out"])
	if err != nil {
		d = 5 * time.Second
	}
	// TODO opt for time out
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{path},
		DialTimeout: d,
	})
	if err != nil {
		// handle error!
	}

	go watchInit(cli, sync)

	//defer cli.Close()

	return &Store{sync, cli}, nil
}

// Store is
type Store struct {
	sync store.Sync
	clientv3.KV
}

func watchInit(cli clientv3.Watcher, sync store.Sync) {
	channelCh := cli.Watch(
		context.Background(),
		"/channels",
		clientv3.WithPrefix(),
	)
	msgCh := cli.Watch(
		context.Background(),
		"/messages",
		clientv3.WithPrefix(),
	)
	logrus.Debug("[store:etcd:watchStart]")
	for {
		select {
		case res := <-channelCh:
			for _, ev := range res.Events {
				logrus.Debugf("[store:etcd:receive-event] channel changed %s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
				if ev.Type == clientv3.EventTypePut {
					channel := &datatype.Channel{}
					err := json.Unmarshal(ev.Kv.Value, channel)
					if err != nil {
						logrus.Warnf("[store:etcd:receive-event] Unmarshal error: %s", err.Error())
						continue
					}
					sync.PutChannel(channel)
				}
			}

		case res := <-msgCh:
			for _, ev := range res.Events {
				logrus.Debugf("[store:etcd:receive-event] channel changed %s %q", ev.Type, ev.Kv.Key)
				if ev.Type == clientv3.EventTypePut {
					msg := &datatype.Message{}
					err := json.Unmarshal(ev.Kv.Value, msg)
					if err != nil {
						logrus.Warnf("[store:etcd:receive-event] Unmarshal error: %s", err.Error())
						continue
					}
					arr := strings.SplitN(string(ev.Kv.Key), "/", 4)
					if len(arr) < 4 {
						logrus.Warnf("[store:etcd:receive-event] parse error: connot find channelID from key")
						continue
					}
					channelID := arr[2]
					sync.PutMessage(channelID, msg)
				}
			}
		}
	}
}
