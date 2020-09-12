package block

import (
	"crypto/sha256"
	"fmt"
	"hash"
	"time"
)

// Block is an individual block in a blockchain.
type Block struct {
	Index     int
	Timestamp time.Time
	Hash      hash.Hash
	PrevHash  hash.Hash
	Data      string
}

// Blockchain is a slice of Blocks.
type Blockchain []*Block

// AddBlock attempts to add a given block to the blockchain.
func (bc Blockchain) AddBlock(b *Block) (Blockchain, error) {
	pb := bc[len(bc)-1]
	if !isValidBlock(b, pb) {
		return nil, fmt.Errorf("Attempted to add invalid block to chain.")
	}

	bc = append(bc, b)
	return bc, nil
}

// isValidBlock returns true if the following conditions are met:
// - The previous block's Index is one below the new block's Index.
// - The previous block's Hash is equal to the new block's previous Hash.
// - The new block's Hash member variable must equal the result of actually calculating the block's Hash.
func isValidBlock(b *Block, pb *Block) bool {
	return b.Index-1 == pb.Index && b.PrevHash == pb.Hash && b.Hash == b.genHash()
}

// IsValid returns true if all blocks in the blockchain are valid.
func (bc Blockchain) IsValid() bool {
	for i := len(bc) - 1; i > 0; i-- {
		if !isValidBlock(bc[i], bc[i-1]) {
			return false
		}
	}
	return true
}

// NewBlock returns a new Block holding the information passed in, a Timestamp, and a generated Hash.
func (bc Blockchain) NewBlock(data string) *Block {
	b := &Block{Index: len(bc), Timestamp: time.Now(), Data: data, PrevHash: bc[len(bc)-1].Hash}
	b.Hash = b.genHash()

	return b
}

// NewChain returns an in memory blockchain, only containing the genesis block.
func NewChain() Blockchain {
	bc := make(Blockchain, 0)
	bc = append(bc, genesis())

	return bc
}

// genHash generates and saves a Hash for this block
func (b *Block) genHash() hash.Hash {
	// Init Hash
	h := sha256.New()

	// Create string payload
	payload := b.Timestamp.String() + b.Data + string(b.Index)

	// Write payload to Hash
	h.Write([]byte(payload))

	return h
}

// genesis returns the first Block in a blockchain.
func genesis() *Block {
	b := &Block{Index: 0, Timestamp: time.Now(), Data: "First block", PrevHash: nil}
	b.Hash = b.genHash()

	return b
}
