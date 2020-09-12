package block

import (
	"crypto/sha256"
	"fmt"
	"hash"
	"time"
)

type Block struct {
	index     int
	timestamp time.Time
	hash      hash.Hash
	prevHash  hash.Hash
	data      string
}

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
// - The previous block's index is one below the new block's index.
// - The previous block's hash is equal to the new block's previous hash.
// - The new block's hash member variable must equal the result of actually calculating the block's hash.
func isValidBlock(b *Block, pb *Block) bool {
	return b.index-1 == pb.index && b.prevHash == pb.hash && b.hash == b.genHash()
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

// New returns a new Block holding the information passed in, a timestamp, and a generated hash.
func New(i int, d string, ph hash.Hash) *Block {
	b := &Block{index: i, timestamp: time.Now(), data: d, prevHash: ph}
	b.genHash()

	return b
}

// NewChain returns an in memory blockchain, only containing the genesis block.
func NewChain() Blockchain {
	bc := make(Blockchain, 0)
	bc = append(bc, genesis())

	return bc
}

// genHash generates and saves a hash for this block
func (b *Block) genHash() hash.Hash {
	// Init hash
	h := sha256.New()

	// Create string payload
	payload := b.timestamp.String() + b.data + string(b.index)

	// Write payload to hash
	h.Write([]byte(payload))

	return h
}

// genesis returns the first Block in a blockchain.
func genesis() *Block {
	b := &Block{index: 0, timestamp: time.Now(), data: "First block", prevHash: nil}
	b.hash = b.genHash()

	return b
}
