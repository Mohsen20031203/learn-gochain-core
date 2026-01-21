package blockchain

/*
import (
	"errors"
	"testing"

	"github.com/Mohsen20031203/learn-gochain-core/internal/domain/block"
)

type FakeRepo struct {
	data map[string]*block.Block
	err  error
}

func NewFakeRepo() *FakeRepo {
	return &FakeRepo{
		data: make(map[string]*block.Block),
	}
}

func (f *FakeRepo) Get(key string) (*block.Block, error) {
	if f.err != nil {
		return nil, f.err
	}
	v, ok := f.data[key]
	if !ok {
		return nil, errors.New("not found")
	}
	return v, nil
}

func (f *FakeRepo) Save(key string, value *block.Block) error {
	if f.err != nil {
		return f.err
	}
	f.data[key] = value
	return nil
}

type FakeNode struct {
	lastHash string
	updated  bool
}

func (f *FakeNode) GetChainLastBlockHash() string {
	return f.lastHash
}

func (f *FakeNode) UpdateChain(b block.Block) {
	f.lastHash = b.Hash
	f.updated = true
}

func TestSaveBlock_HappyPath(t *testing.T) {
	repo := NewFakeRepo()
	node := &FakeNode{}

	service := &NodeService{
		repo: repo,
		node: node,
	}

	b := &block.Block{
		Hash: "hash1",
	}

	err := service.saveBlock(b)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// بلاک ذخیره شده؟
	if repo.data[b.Hash] == nil {
		t.Error("block was not saved by hash")
	}

	// last block ذخیره شده؟
	if repo.data[LastBlockKey] == nil {
		t.Error("last block was not saved")
	}

	// node آپدیت شده؟
	if !node.updated {
		t.Error("node chain was not updated")
	}
}

func TestSaveBlock_RepoError(t *testing.T) {
	repo := NewFakeRepo()
	repo.err = errors.New("repo failed")

	node := &FakeNode{}

	service := &NodeService{
		repo: repo,
		node: node,
	}

	b := &block.Block{Hash: "hash1"}

	err := service.saveBlock(b)

	if err == nil {
		t.Error("expected error, got nil")
	}
}

func TestGetBlockByHash_HappyPath(t *testing.T) {
	repo := NewFakeRepo()
	node := &FakeNode{}

	b := &block.Block{Hash: "hash1"}
	repo.data["hash1"] = b

	service := &NodeService{
		repo: repo,
		node: node,
	}

	result, err := service.GetBlockByHash("hash1")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.Hash != "hash1" {
		t.Error("wrong block returned")
	}
}

func TestGetBlockByHash_NotFound(t *testing.T) {
	repo := NewFakeRepo()
	node := &FakeNode{}

	service := &NodeService{
		repo: repo,
		node: node,
	}

	_, err := service.GetBlockByHash("unknown")

	if err == nil {
		t.Error("expected error, got nil")
	}
}
*/
