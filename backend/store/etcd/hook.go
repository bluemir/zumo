package etcd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/bluemir/zumo/datatype"
)

func (s *Store) GetHook(hookID string) (*datatype.Hook, error) {
	resp, err := s.KV.Get(
		context.Background(),
		fmt.Sprintf("/hooks/%s", hookID),
	)
	if err != nil {
		return nil, err
	}

	if len(resp.Kvs) != 1 {
		return nil, errors.New("not unique")
	}

	hook := &datatype.Hook{}
	err = json.Unmarshal(resp.Kvs[0].Value, hook)
	if err != nil {
		return nil, err
	}

	return hook, nil
}
func (s *Store) PutHook(hook *datatype.Hook) (*datatype.Hook, error) {
	str, err := json.Marshal(hook)
	if err != nil {
		return nil, err
	}

	_, err = s.KV.Put(
		context.Background(),
		fmt.Sprintf("/hooks/%s", hook.ID),
		string(str),
	)

	if err != nil {
		return nil, err
	}

	return hook, nil
}
