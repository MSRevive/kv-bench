package kv

import (
	"github.com/nutsdb/nutsdb"
)

type nutsdbStore struct {
	db *nutsdb.DB
}

func newNutsDB(path string) (Store, error) {
	options := nutsdb.DefaultOptions
	options.Dir = path
	options.EntryIdxMode = nutsdb.HintKeyAndRAMIdxMode
	options.SyncEnable = false

	db, err := nutsdb.Open(options)
	if err != nil {
		return nil, err
	}
	
	err = db.Update(func(tx *nutsdb.Tx) error {
		return tx.NewKVBucket("bucket")
	})
	if err != nil {
		return nil, err
	}

	return &nutsdbStore{db: db}, nil
}

func (n nutsdbStore) Put(key []byte, value []byte) error {
	return n.db.Update(func(tx *nutsdb.Tx) error {
		return tx.Put("bucket", key, value, nutsdb.Persistent)
	})
}

func (n nutsdbStore) Get(key []byte) ([]byte, error) {
	var (
		value []byte
	)
	err := n.db.View(func(tx *nutsdb.Tx) error {
		val, err2 := tx.Get("bucket", key)
		if err2 != nil {
			return err2
		}
		value = val
		return nil
	})
	return value, err
}

func (n nutsdbStore) Delete(key []byte) error {
	return n.db.Update(func(tx *nutsdb.Tx) error {
		return tx.Delete("bucket", key)
	})
}

func (n nutsdbStore) Close() error {
	return n.db.Close()
}
