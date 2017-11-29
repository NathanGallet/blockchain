package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
)

// Constant bug should take in count speed of minage
const amount_of_reward = 25

type Transaction struct {
	ID   []byte
	Vin  []TXInput
	Vout []TXOutput
}

func NewTransaction(to, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Reward to '%s'", to)
	}

	txin := TXInput{[]byte{}, -1, data}
	txout := TXOutput{amount_of_reward, to}
	tx := Transaction{nil, []TXInput{txin}, []TXOutput{txout}}
	tx.SetID()

	return &tx
}

// Set an ID
func (transaction Transaction) SetID() {
	var encoded bytes.Buffer
	var hash [32]byte

	encoder := gob.NewEncoder(&encoded)
	// Encode the transaction
	err := encoder.Encode(transaction)
	if err != nil {
		log.Fatal(err)
	}

	// Sum256 in order to have an ID
	hash = sha256.Sum256(encoded.Bytes())
	transaction.ID = hash[:]
}
