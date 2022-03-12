package blockchain

import (
	"fmt"
	"log"

	"github.com/dgraph-io/badger/v3"
)

const (
	dbPath = "./blocks"
)

type BlockChain struct {
	LastHash []byte
	Database *badger.DB
}

type BlockChainIterator struct {
	CurrentHash []byte
	Database    *badger.DB
}

func InitBlockChain() *BlockChain {
	var lastHash []byte

	opts := badger.DefaultOptions(dbPath)

	db, err := badger.Open(opts)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Update(func(txn *badger.Txn) error {
		if _, err := txn.Get([]byte("lh")); err == badger.ErrKeyNotFound {
			fmt.Println("No existing blockchain found")
			genesis := Genesis()
			fmt.Println("Genesis proved")

			serialized := genesis.Serialize()
			err = txn.Set(genesis.Hash, serialized)
			if err != nil {
				return fmt.Errorf("failed txn set %w", err)
			}

			// set lash hash
			err = txn.Set([]byte("lh"), genesis.Hash)
			lastHash = genesis.Hash
			return err
		} else {
			item, err := txn.Get([]byte("lh"))
			if err != nil {
				return err
			}

			err = item.Value(func(val []byte) error {
				lastHash = val
				return nil
			})
			return err
		}
	})
	if err != nil {
		log.Fatal("failed to init blockhain ", err)
	}

	return &BlockChain{lastHash, db}
}

func (b *BlockChain) AddBlock(data string) {
	var lastHash []byte

	err := b.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("lh"))
		if err != nil {
			return err
		}

		err = item.Value(func(val []byte) error {
			lastHash = val
			return nil
		})
		return err
	})

	if err != nil {
		log.Fatal(err)
	}

	newBlock := createBlock(data, lastHash)

	err = b.Database.Update(func(txn *badger.Txn) error {
		err := txn.Set(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			return err
		}

		err = txn.Set([]byte("lh"), newBlock.Hash)
		if err != nil {
			return err
		}

		b.LastHash = newBlock.Hash
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}

func (b *BlockChain) Iterator() *BlockChainIterator {
	return &BlockChainIterator{b.LastHash, b.Database}
}

func (iter *BlockChainIterator) Next() *Block {
	var block *Block

	err := iter.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get(iter.CurrentHash)
		if err != nil {
			return err
		}

		err = item.Value(func(val []byte) error {
			block = Deserialize(val)
			return nil
		})
		return err
	})

	if err != nil {
		log.Fatal(err)
	}

	iter.CurrentHash = block.PrevHash
	return block
}
