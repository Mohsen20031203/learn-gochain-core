package block

/*
func TestCalculateHash(t *testing.T) {
	b := Block{
		Index:     1,
		Timestamp: time.Now(),
		Data:      "Hello",
		PrevHash:  "0000",
		Nonce:     0,
	}

	hash := b.CalculateHash()
	if hash == "" {
		t.Error("Hash should not be empty")
	}
}

func TestMineAndHasValidPoW(t *testing.T) {
	b := Block{
		Index:     1,
		Timestamp: time.Now(),
		Data:      "Test",
		PrevHash:  "0000",
	}

	difficulty := 2
	b.Mine(difficulty)

	if !b.HasValidPoW(difficulty) {
		t.Error("Block should have valid PoW")
	}
}

func TestIsValid(t *testing.T) {
	prev := Block{
		Hash: "abcd1234",
	}

	b := Block{
		PrevHash: "abcd1234",
	}

	b.Hash = b.CalculateHash()

	if !b.IsValid(prev) {
		t.Error("Block should be valid with correct PrevHash")
	}

	b.PrevHash = "wrong"
	if b.IsValid(prev) {
		t.Error("Block should be invalid with wrong PrevHash")
	}
}

*/
