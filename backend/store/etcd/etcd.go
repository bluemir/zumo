package etcd

import (
	"time"

	"github.com/bluemir/zumo/backend/store"
	"github.com/bluemir/zumo/datatype"
	"github.com/coreos/etcd/client"
)

func init() {
	store.Register("etcd", New)
}

// New is
func New(path string, sync store.Sync, opt map[string]string) (store.Store, error) {
	cfg := client.Config{
		Endpoints: []string{"http://127.0.0.1:2379"},
		Transport: client.DefaultTransport,
		// set timeout per request to fail fast when the target endpoint is unavailable
		HeaderTimeoutPerRequest: time.Second,
	}
	c, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	api := client.NewKeysAPI(c)

	return &Store{sync, api}, nil
}

// Store is
type Store struct {
	sync store.Sync
	client.KeysAPI
}

func (s *Store) FindMessages(channelID string, limit int) ([]datatype.Message, error) {
	return nil, nil
}
func (s *Store) PutMessage(channelID string, msg *datatype.Message) (*datatype.Message, error) {
	return nil, nil
}

func (s *Store) FindUser() ([]datatype.User, error) {
	return nil, nil
}
func (s *Store) GetUser(username string) (*datatype.User, error) {
	return nil, nil
}
func (s *Store) PutUser(user *datatype.User) (*datatype.User, error) {
	return nil, nil
}

func (s *Store) GetToken(username, hashedKey string) (*datatype.Token, error) {
	return nil, nil
}
func (s *Store) PutToken(token *datatype.Token) (*datatype.Token, error) {
	return nil, nil
}

func (s *Store) GetHook(hookID string) (*datatype.Hook, error) {
	return nil, nil
}
func (s *Store) PutHook(*datatype.Hook) (*datatype.Hook, error) {
	return nil, nil
}

// genaral data, for pod or bots.
func (s *Store) GetData(namespace, key string, data interface{}) error {
	return nil
}
func (s *Store) PutData(namespace, key string, data interface{}) error {
	return nil
}
