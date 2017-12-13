package etcd

import (
	"context"
	"encoding/json"
	"log"

	"github.com/bluemir/zumo/datatype"
	"github.com/coreos/etcd/client"
)

func (s *Store) FindChannels() ([]datatype.Channel, error) {
	return nil, nil
}
func (s *Store) GetChannel(channelID string) (*datatype.Channel, error) {
	return nil, nil
}
func (s *Store) PutChannel(channel *datatype.Channel) (*datatype.Channel, error) {
	str, err := json.Marshal(channel)
	if err != nil {
		return nil, err
	}

	resp, err := s.Set(context.Background(), "/channels/"+channel.ID, string(str), &client.SetOptions{
		PrevExist: client.PrevNoExist,
	})
	if err != nil {
		log.Fatal(err)
	} else {
		// print common key info
		log.Printf("Set is done. Metadata is %q\n", resp)
	}

	return channel, nil
}
func (s *Store) DeleteChannel(ID string) error {

	return nil
}
