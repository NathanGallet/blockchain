package main

import (
	"log"

	"github.com/boltdb/bolt"
)

const dbFile = "blockchain.db"
const blocksBucket = "blocks"

// Blockchain structure
type Blockchain struct {
	last_block_hash []byte
	database        *bolt.DB
}

// Add a new block to the blockchain
func (blockchain *Blockchain) AddBlock(data string) {
	var lastHash []byte

	// To start a read-only transaction
	err := blockchain.database.View(func(tx *bolt.Tx) error {
		// Get bucket
		b := tx.Bucket([]byte(blocksBucket))

		// l key is the last block hash
		lastHash = b.Get([]byte("l"))

		// commit changement
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	// create new block with previously hash and new data
	newBlock := NewBlock(data, lastHash)

	// Update our database
	err = blockchain.database.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		err := b.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			log.Fatal(err)
		}

		err = b.Put([]byte("l"), newBlock.Hash)
		if err != nil {
			log.Fatal(err)
		}

		blockchain.last_block_hash = newBlock.Hash

		return nil
	})
}

// First block
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

// Initialisation blockchain
func NewBlockchain() *Blockchain {
	var last_block_hash []byte

	// Open database
	database, err := bolt.Open(dbFile, 0600, nil)

	// If there is an error, log it
	if err != nil {
		log.Fatal(err)
	}

	// Inside the closure, you have a consistent view of the database. You commit the transaction by returning nil at the end.
	err = database.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))

		if b == nil {
			genesis := NewGenesisBlock()

			// Buckets are collections of key/value pairs within the database
			b, err := tx.CreateBucket([]byte(blocksBucket))
			if err != nil {
				return err
			}

			// Set genesis.Hash as key and genesis.Serialise as value
			err = b.Put(genesis.Hash, genesis.Serialize())
			if err != nil {
				return err
			}

			// 'l' -> 4-byte file number: the last block file number used
			err = b.Put([]byte("l"), genesis.Hash)
			if err != nil {
				return err
			}

			// keep the hash of our generic hash
			last_block_hash = genesis.Hash
		} else {
			// if we already have this bucket, we need the last block
			last_block_hash = b.Get([]byte("l"))
		}

		// commit our modification
		return nil
	})

	// Return an instance of our blockchain
	blockchain := Blockchain{last_block_hash, database}

	return &blockchain
}
