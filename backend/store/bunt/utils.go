package bunt

import (
	"encoding/json"

	"github.com/tidwall/buntdb"
)

func (s *Store) get(path string, v interface{}) error {
	return s.DB.View(func(tx *buntdb.Tx) error {
		val, err := tx.Get(path)
		if err != nil {
			return err
		}

		return json.Unmarshal([]byte(val), v)
	})
}
func (s *Store) first(name string, v interface{}) error {
	return s.DB.View(func(tx *buntdb.Tx) error {
		var e error
		if err := tx.Ascend(name, func(key, value string) bool {
			e = json.Unmarshal([]byte(value), v)
			return false
		}); err != nil {
			return err
		}
		return e
	})
}
func (s *Store) list(pattern string) ([]string, error) {
	result := []string{}
	err := s.DB.View(func(tx *buntdb.Tx) error {
		return tx.AscendKeys(pattern, func(key, value string) bool {
			result = append(result, value)
			return true
		})

	})
	return result, err
}

// if limit < 0 mean infinite
func (s *Store) find(pattern string, limit int) ([]string, error) {
	result := []string{}
	count := 0
	err := s.DB.View(func(tx *buntdb.Tx) error {

		return tx.DescendKeys(pattern, func(key, value string) bool {
			result = append(result, value)
			count++
			if count > limit && limit > 0 {
				return false
			}
			return true
		})

	})
	return result, err
}

func bind(v string, T interface{}) error {
	return json.Unmarshal([]byte(v), T)
}
func (s *Store) put(path string, v interface{}) error {
	return s.DB.Update(func(tx *buntdb.Tx) error {
		buf, err := json.Marshal(v)
		if err != nil {
			return err
		}
		_, _, err = tx.Set(path, string(buf), nil)
		if err != nil {
			return err
		}
		return nil
	})
}
func (s *Store) remove(path string) error {
	return s.DB.Update(func(tx *buntdb.Tx) error {
		_, err := tx.Delete(path)
		return err
	})
}
func isNotFound(err error) bool {
	return err == buntdb.ErrNotFound
}
