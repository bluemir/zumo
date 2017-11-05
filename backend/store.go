package backend

import "github.com/bluemir/zumo/backend/store"

// DataStore for pods or bots, it just for small data
// if want big data bot SHOULD make own store
type DataStore interface {
	Save(key string, data interface{}) error
	Load(key string, data interface{}) error
}
type dstore struct {
	store.Store
	namespace string
}

func (ds *dstore) Save(key string, data interface{}) error {
	return ds.PutData(ds.namespace, key, data)
}
func (ds *dstore) Load(key string, data interface{}) error {
	return ds.GetData(ds.namespace, key, data)
}

func (b *backend) RequestDataStore(namespace string) DataStore {
	return &dstore{b.store, namespace}
}
