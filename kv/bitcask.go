package kv

import (
	"go.mills.io/bitcask/v2"
)

func newBitcask(path string) (Store, error) {
	db, err := bitcask.Open(path)
	return &bitcaskStore{db: db}, err
}

type bitcaskStore struct {
	db *bitcask.Bitcask
}

func (s *bitcaskStore) Put(key []byte, value []byte) error {
	return s.db.Put(key, value)
}

func (s *bitcaskStore) Get(key []byte) ([]byte, error) {
	return s.db.Get(key)
}

func (s *bitcaskStore) Delete(key []byte) error {
	return s.db.Delete(key)
}

func (s *bitcaskStore) Close() error {
	return s.db.Close()
}
