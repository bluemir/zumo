package bunt

import "path"

// this one used by pod, bots, other plugins. it is not offer listing...
// no indexing, no listing

// GetData is get genaral Data from store
func (s *Store) GetData(namespace, key string, data interface{}) error {
	return s.get(path.Join("data", namespace, key), data)
}

// PutData is get genaral Data from store
func (s *Store) PutData(namespace, key string, data interface{}) error {
	return s.put(path.Join("data", namespace, key), data)
}
