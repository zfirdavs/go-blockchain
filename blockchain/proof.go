package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"math"
	"math/big"
)

var difficulty = flag.Int("difficulty", 12, "set the algorithm hash difficulty")

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

func NewProof(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-(*difficulty)))
	return &ProofOfWork{b, target}
}

func (pow *ProofOfWork) InitData(nonce int) []byte {
	info := bytes.Join([][]byte{
		pow.Block.PrevHash,
		pow.Block.Data,
		ToByteHex(int64(nonce)),
		ToByteHex(int64(*difficulty)),
	},
		[]byte{},
	)
	return info
}

func (pow *ProofOfWork) Run() (int, []byte) {
	var (
		intHash big.Int
		hash    [32]byte
		nonce   int
	)

	for nonce < math.MaxInt64 {
		data := pow.InitData(nonce)
		hash = sha256.Sum256(data)

		fmt.Printf("\r%x", hash)
		intHash.SetBytes(hash[:])

		if intHash.Cmp(pow.Target) == -1 {
			break
		} else {
			nonce++
		}
	}

	fmt.Println()
	return nonce, hash[:]
}

func (pow *ProofOfWork) Validate() bool {
	var intHash big.Int

	data := pow.InitData(pow.Block.Nonce)
	hash := sha256.Sum256(data)

	intHash.SetBytes(hash[:])
	return intHash.Cmp(pow.Target) == -1
}

func ToByteHex(num int64) []byte {
	buff := new(bytes.Buffer)
	if err := binary.Write(buff, binary.BigEndian, num); err != nil {
		log.Fatal(err)
	}

	return buff.Bytes()
}
