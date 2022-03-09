package blockchain

func Genesis() *Block {
	return createBlock("Genesis", []byte{})
}

type Block struct {
	Hash, Data, PrevHash []byte
	Nonce                int
}

func createBlock(data string, prevHash []byte) *Block {
	block := &Block{
		Hash:     []byte{},
		Data:     []byte(data),
		PrevHash: prevHash,
	}

	poW := NewProof(block)
	nonce, hash := poW.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}
