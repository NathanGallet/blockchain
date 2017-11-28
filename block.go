package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
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
