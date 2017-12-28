package etcd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/bluemir/zumo/datatype"
)

func (s *Store) GetToken(username, hashedKey string) (*datatype.Token, error) {

	resp, err := s.KV.Get(
		context.Background(),
		fmt.Sprintf("/tokens/%s:%s", username, hashedKey),
	)
	if err != nil {
		return nil, err
	}

	if len(resp.Kvs) != 1 {
		return nil, errors.New("not unique")
	}

	token := &datatype.Token{}
	err = json.Unmarshal(resp.Kvs[0].Value, token)
	if err != nil {
		return nil, err
	}

	return token, nil
}
func (s *Store) PutToken(token *datatype.Token) (*datatype.Token, error) {
	str, err := json.Marshal(token)
	if err != nil {
		return nil, err
	}

	_, err = s.KV.Put(
		context.Background(),
		fmt.Sprintf("/tokens/%s:%s", token.Username, token.HashedKey),
		string(str),
	)
	if err != nil {
		return nil, err
	}

	return token, nil
}
