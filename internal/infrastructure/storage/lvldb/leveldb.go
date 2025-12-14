package lvldb

import (
	"encoding/json"

	"github.com/Mohsen20031203/learn-gochain-core/internal/domain/block"
	"github.com/syndtr/goleveldb/leveldb"
)

type leveldbStorage struct {
	path string
	db   *leveldb.DB
}

func New(path string) *leveldbStorage {
	return &leveldbStorage{path: path}
}

func (l *leveldbStorage) Open() error {
	db, err := leveldb.OpenFile(l.path, nil)
	if err != nil {
		return err
	}
	l.db = db
	return nil
}

func (l *leveldbStorage) Save(key string, block *block.Block) error {

	data, err := json.Marshal(block)
	if err != nil {
		return err
	}

	err = l.db.Put([]byte(key), data, nil)
	if err != nil {
		return err
	}
	return nil
}

func (l *leveldbStorage) Load() (*block.Block, error) {
	// Implementation for loading the blockchain from LevelDB goes here
	return nil, nil
}

func (l *leveldbStorage) Close() error {
	return l.db.Close()
}

func (l *leveldbStorage) Get(key string) (*block.Block, error) {

	block := block.Block{}
	value, err := l.db.Get([]byte(key), nil)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(value, &block)
	if err != nil {
		return nil, err
	}

	return &block, nil
}
