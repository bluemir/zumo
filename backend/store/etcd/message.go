package etcd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/coreos/etcd/clientv3"

	"github.com/bluemir/zumo/datatype"
)

func (s *Store) FindMessages(channelID string, limit int) ([]datatype.Message, error) {
	resp, err := s.Get(context.Background(),
		fmt.Sprintf("/messages/%s/", channelID),
		clientv3.WithSort(clientv3.SortByKey, clientv3.SortDescend),
		//clientv3.WithLimit(limit),
	)
	if err != nil {
		return nil, err
	}

	result := []datatype.Message{}
	for _, pair := range resp.Kvs {
		msg := &datatype.Message{}
		err := json.Unmarshal(pair.Value, msg)
		if err != nil {
			return nil, err
		}
		result = append(result, *msg)
	}
	return result, nil
}
func (s *Store) PutMessage(channelID string, msg *datatype.Message) (*datatype.Message, error) {
	str, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}
	key := fmt.Sprintf("/messages/%s/%.16x", channelID, msg.Time.UnixNano())

	_, err = s.Put(context.Background(), key, string(str))
	if err != nil {
		return nil, err
	}
	return msg, nil
}
