package db

import "github.com/dgraph-io/badger/v4"

func ReadData(db *badger.DB, key string) (string, error) {
	var value string
	err := db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}
		return item.Value(func(val []byte) error {
			value = string(val)
			return nil
		})
	})
	return value, err
}
