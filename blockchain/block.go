package blockchain

type BlockChain struct {
	Blocks []*Block
}

func (b *BlockChain) AddBlock(data string) {
	prevBlock := b.Blocks[len(b.Blocks)-1]
	newBlock := createBlock(data, prevBlock.Hash)
	b.Blocks = append(b.Blocks, newBlock)
}

func Genesis() *Block {
	return createBlock("Genesis", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
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
