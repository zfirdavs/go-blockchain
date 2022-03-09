package blockchain

type BlockChain struct {
	Blocks []*Block
}

func (b *BlockChain) AddBlock(data string) {
	prevBlock := b.Blocks[len(b.Blocks)-1]
	newBlock := createBlock(data, prevBlock.Hash)
	b.Blocks = append(b.Blocks, newBlock)
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
