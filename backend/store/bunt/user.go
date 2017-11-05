package bunt

import (
	"github.com/bluemir/zumo/datatype"
	"github.com/sirupsen/logrus"
)

func (s *Store) FindUser() ([]datatype.User, error) {
	list, err := s.list("/users/*")
	if err != nil {
		return nil, err
	}
	result := []datatype.User{}
	for _, v := range list {
		user := &datatype.User{}

		if err := bind(v, user); err != nil {
			return nil, err
		}
		result = append(result, *user)
	}
	return result, nil
}
func (s *Store) GetUser(username string) (*datatype.User, error) {
	user := &datatype.User{}

	err := s.get("/users/"+username, user)
	if err != nil {
		if isNotFound(err) {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}
func (s *Store) PutUser(user *datatype.User) (*datatype.User, error) {
	logrus.Debugf("[store:PutUser] %+v", user)
	err := s.put("/users/"+user.Name, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
