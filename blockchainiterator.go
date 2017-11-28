package main

import (
	"log"

	"github.com/boltdb/bolt"
)

type BlockchainIterator struct {
	currentHash []byte
	database    *bolt.DB
}

// Create a new iterator
func (blockchain *Blockchain) NewIterator() *BlockchainIterator {
	blockchain_iterator := &BlockchainIterator{blockchain.last_block_hash, blockchain.database}

	return blockchain_iterator
}

// Return next block starting from last_block_hash
func (iterator *BlockchainIterator) Next() *Block {
	var block *Block

	// connection to the database
	err := iterator.database.View(func(tx *bolt.Tx) error {
		// read the bucket
		b := tx.Bucket([]byte(blocksBucket))

		// get block with the current hash (if it's the first time, last_block_hash)
		encodedBlock := b.Get(iterator.currentHash)

		// decode the block
		block = DeserializeBlock(encodedBlock)

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	// Previous block hash became current hash
	iterator.currentHash = block.PrevBlockHash

	return block
}
