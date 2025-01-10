package db

import (
	"time"

	badger "github.com/dgraph-io/badger/v4"
)

func WriteData(db *badger.DB, key, value string) (time.Duration, error) {
	start := time.Now()
	err := db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(key), []byte(value))
	})
	duration := time.Since(start)
	return duration, err
}
