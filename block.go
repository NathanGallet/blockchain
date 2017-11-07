package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// Block structure (Bitcoin like)
type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

// Set Hash to the block
func (b *Block) SetHash() {
	// Convert timestamp int64 to string and put in a slice
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))

	// Concatenate PrevBlockHash, Data, timestamp ([]byte{} for no separator)
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})

	// Calculate sha256 for this slice of byte
	hash := sha256.Sum256(headers)

	// Save this hash in the block hash
	b.Hash = hash[:]
}

// Generate a new block
func NewBlock(data string, prevBlockHash []byte) *Block {
	// Init new block structure
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}

	// Init proof of work
	pow := NewProofOfWork(block)

	// Launch proof of work
	nonce, hash := pow.Run()

	// Set Hash and nonce
	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}
