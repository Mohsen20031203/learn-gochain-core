package blockchain

/*

func TestValidateNewBlock(t *testing.T) {
	bc := Blockchain{Difficulty: 2}

	prev := block.Block{
		Hash: "0000abcd",
	}

	newBlock := block.Block{
		PrevHash: "0000abcd",
	}
	newBlock.Mine(bc.Difficulty)

	if !bc.ValidateNewBlock(prev, newBlock) {
		t.Error("Block should be valid")
	}

	newBlock.PrevHash = "wrong"
	if bc.ValidateNewBlock(prev, newBlock) {
		t.Error("Block should be invalid with wrong PrevHash")
	}
}

*/
