package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/zfirdavs/go-blockchain/blockchain"
)

func main() {
	flag.Parse()

	blockChain := blockchain.InitBlockChain()
	blockChain.AddBlock("the first block after Genesis")
	blockChain.AddBlock("the second block after Genesis")
	blockChain.AddBlock("the third block after Genesis")

	fmt.Println(strings.Repeat("-", 100))
	for _, block := range blockChain.Blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println(strings.Repeat("-", 100))

		poW := blockchain.NewProof(block)
		fmt.Printf("PoW validation: %s\n", strconv.FormatBool(poW.Validate()))
		fmt.Println()
	}
}
