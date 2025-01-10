package db

import (
	"time"

	badger "github.com/dgraph-io/badger/v4"
)

func DeleteData(db *badger.DB, key string) (time.Duration, error) {
	start := time.Now()
	err := db.Update(func(txn *badger.Txn) error {
		return txn.Delete([]byte(key))
	})
	duration := time.Since(start)
	return duration, err
}
