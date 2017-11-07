package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

var (
	maxNonce = math.MaxInt64
)

// 24 bits means 6 bytes. We want 6 0 on the beginning of the hash
const targetBits = 24

// Proof of Work (pow) structure
type ProofOfWork struct {
	block  *Block
	target *big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork {
	// Allocates and returns a new int64 set to 1.
	target := big.NewInt(1)

	// 1 is shifted left by 256 - targetBits
	target.Lsh(target, uint(256-targetBits))

	// Init pow structure
	pow := &ProofOfWork{b, target}

	return pow
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	// Concatenate Hash, Data, Timestamps, targetBits, nonce
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			IntToHex(pow.block.Timestamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		}, []byte{})

	return data
}

// Find the right nonce
func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int // signed integer
	var hash [32]byte
	nonce := 0

	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)

	for nonce < maxNonce {
		// Calculate Hash
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)

		// convert hash to bytes stored in int64
		hashInt.SetBytes(hash[:])

		// Find a nonce that verified bytes of sha256 < pow target
		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}

	fmt.Printf("\n\n")
	return nonce, hash[:]
}

func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1

	return isValid
}
