package etcd

import (
	"context"
	"fmt"
	"time"

	"github.com/bluemir/zumo/backend/store"
	"github.com/coreos/etcd/clientv3"
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
	for {
		select {
		case res := <-channelCh:
			for _, ev := range res.Events {
				fmt.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			}

		case res := <-msgCh:
			for _, ev := range res.Events {
				fmt.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			}
		}
	}
}
