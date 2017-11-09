package bunt

import (
	"path"

	"github.com/bluemir/zumo/datatype"
)

func (s *Store) GetHook(hookID string) (*datatype.Hook, error) {
	hook := &datatype.Hook{}
	err := s.get(path.Join("hook", hookID), hook)
	if err != nil {
		return nil, err
	}
	return hook, nil
}
func (s *Store) PutHook(hook *datatype.Hook) (*datatype.Hook, error) {
	err := s.put(path.Join("hook", hook.ID), hook)
	if err != nil {
		return nil, err
	}
	return hook, nil
}
