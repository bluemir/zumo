package bunt

import (
	"github.com/tidwall/buntdb"

	"github.com/bluemir/zumo/backend/store"
)

func init() {
	store.Register("bunt", New)
}

func New(path string, sync store.Sync, opt map[string]string) (store.Store, error) {
	db, err := buntdb.Open(path)
	if err != nil {
		return nil, err
	}

	return &Store{db, sync}, nil
}

type Store struct {
	*buntdb.DB
	sync store.Sync
}
