package db

import (
	"time"

	"github.com/dgraph-io/badger/v4"
)

func ReadData(db *badger.DB, key string) (string, time.Duration, error) {
	var value string
	start := time.Now()
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
	duration := time.Since(start)
	return value, duration, err
}
