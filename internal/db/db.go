package db

import (
	badger "github.com/dgraph-io/badger/v4"
)

func InitDB(path string) (*badger.DB, error) {
	options := badger.DefaultOptions(path).WithLoggingLevel(badger.INFO)
	db, err := badger.Open(options)
	if err != nil {
		return nil, err
	}
	return db, nil
}
