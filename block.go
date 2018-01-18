package main

import (
	"bytes"
	"encoding/gob"
	"log"
	"time"
)

// Block structure (Bitcoin like)
type Block struct {
	Timestamp     int64
	Transaction   []*Transaction
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

func (b *Block) Serialize() []byte {
	// A Buffer is a variable-sized buffer of bytes with Read and Write methods.
	var result bytes.Buffer

	// Link encoder to result
	encoder := gob.NewEncoder(&result)

	// Encode Block
	err := encoder.Encode(b)

	if err != nil {
		log.Fatal(err)
	}

	return result.Bytes()
}

func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))

	// Decode Block
	err := decoder.Decode(&block)

	if err != nil {
		log.Fatal(err)
	}

	return &block
}

// Generate a new block
func NewBlock(transaction []*Transaction, prevBlockHash []byte) *Block {
	// Init new block structure
	block := &Block{time.Now().Unix(), transaction, prevBlockHash, []byte{}, 0}

	// Init proof of work
	pow := NewProofOfWork(block)

	// Launch proof of work
	nonce, hash := pow.Run()

	// Set Hash and nonce
	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

func NewGenesisBlock(coinbase *Transaction) *Block {
	return NewBlock([]*Transaction{coinbase}, []byte{})
}
