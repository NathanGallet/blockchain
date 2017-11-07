package main

// Blockchain structure
type Blockchain struct {
	blocks []*Block
}

// Add a new block to the blockchain
func (bc *Blockchain) AddBlock(data string) {
	// Get the last block in the blockchain
	prevBlock := bc.blocks[len(bc.blocks)-1]

	// Generate a new block with the data and the previous block hash
	newBlock := NewBlock(data, prevBlock.Hash)

	// update the blockchain with this new block
	bc.blocks = append(bc.blocks, newBlock)
}

// First block
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

// Initialisation blockchain
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
