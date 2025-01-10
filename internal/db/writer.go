package db

import (
	badger "github.com/dgraph-io/badger/v4"
)

func WriteData(db *badger.DB, key, value string) error {
	return db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(key), []byte(value))
	})
}
