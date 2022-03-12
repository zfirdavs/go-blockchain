package blockchain

import (
	"bytes"
	"encoding/gob"
	"log"
)

func Genesis() *Block {
	return createBlock("Genesis", []byte{})
}

type Block struct {
	Hash, Data, PrevHash []byte
	Nonce                int
}

func (b *Block) Serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)

	if err := encoder.Encode(b); err != nil {
		log.Fatal("failed to encode block ", err)
	}

	return res.Bytes()
}

func Deserialize(data []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(data))

	if err := decoder.Decode(&block); err != nil {
		log.Fatal("failed to decode block ", err)
	}
	return &block
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
