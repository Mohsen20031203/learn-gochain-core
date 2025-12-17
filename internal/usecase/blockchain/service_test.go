package blockchain

/*
type mockRepo struct {
	data map[string]*block.Block
}

func (m *mockRepo) Open() error {
	m.data = make(map[string]*block.Block)
	return nil
}
func (m *mockRepo) Save(key string, b *block.Block) error {
	m.data[key] = b
	return nil
}
func (m *mockRepo) Get(key string) (*block.Block, error) {
	if b, ok := m.data[key]; ok {
		return b, nil
	}
	return nil, nil
}
func (m *mockRepo) Close() error { return nil }

func TestAddBlock(t *testing.T) {

	folderPath := "mockpath"

	info, err := os.Stat(folderPath)
	if err == nil && info.IsDir() {
		err := os.RemoveAll(folderPath)
		if err != nil {
			fmt.Println("Error removing folder:", err)
			return
		}
		fmt.Println("Folder removed successfully")
	} else if os.IsNotExist(err) {
		fmt.Println("Folder does not exist, continue")
	} else if err != nil {
		fmt.Println("Error checking folder:", err)
		return
	}

	cng := &config.Config{
		FileStoragePath: folderPath,
		Difficulty:      2,
	}
	svc := NewService(*cng)

	blk, err := svc.AddBlock("Hello")
	if err != nil {
		t.Fatalf("AddBlock failed: %v", err)
	}

	if blk.Data != "Hello" {
		t.Error("Block data mismatch")
	}

	if blk.Index != 0 {
		t.Error("Block index should be 0 for first block")
	}
	if blk.PrevHash != "0" {
		t.Error("PrevHash should be 0 for first block")
	}

	blk2, err := svc.AddBlock("Second")
	if err != nil {
		t.Fatalf("AddBlock failed: %v", err)
	}

	if blk2.Index != 1 {
		t.Error("Block index should be 1 for second block")
	}
	if blk2.PrevHash != blk.Hash {
		t.Error("PrevHash should point to previous block hash")
	}
}

func TestAddBlockInvalid(t *testing.T) {

	folderPath := "mockpath"

	info, err := os.Stat(folderPath)
	if err == nil && info.IsDir() {
		err := os.RemoveAll(folderPath)
		if err != nil {
			fmt.Println("Error removing folder:", err)
			return
		}
		fmt.Println("Folder removed successfully")
	} else if os.IsNotExist(err) {
		fmt.Println("Folder does not exist, continue")
	} else if err != nil {
		fmt.Println("Error checking folder:", err)
		return
	}

	cng := &config.Config{
		FileStoragePath: folderPath,
		Difficulty:      2,
	}

	svc := NewService(*cng)

	_, err = svc.AddBlock("First")
	if err != nil {
		t.Fatalf("AddBlock failed: %v", err)
	}

	badBlock := block.Block{
		Timestamp: time.Now(),
		Data:      "Bad",
		Index:     1,
		PrevHash:  "wrong",
	}
	svc.repo.Save("last", &badBlock)

	_, err = svc.AddBlock("NextBlock")
	if err == nil {
		t.Error("Expected error for invalid block")
	}

}

*/
