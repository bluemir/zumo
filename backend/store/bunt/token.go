package bunt

import (
	"github.com/bluemir/zumo/datatype"
)

const (
	TokenPrefix = "/token/"
)

func (s *Store) GetToken(username, hashedKey string) (*datatype.Token, error) {
	token := &datatype.Token{}
	err := s.get(TokenPrefix+username+":"+hashedKey, token)
	if err != nil {
		return nil, err
	}
	return token, nil
}
func (s *Store) PutToken(token *datatype.Token) (*datatype.Token, error) {
	err := s.put(TokenPrefix+token.Username+":"+token.HashedKey, token)
	if err != nil {
		return token, err
	}
	return token, nil
}
