package block

import (
	"crypto/sha256"
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

// New returns a new Block holding the information passed in, a timestamp, and a generated hash.
func New(i int, d string, ph hash.Hash) *Block {
	b := &Block{index: i, timestamp: time.Now(), data: d, prevHash: ph}
	b.genHash()

	return b
}

// NewChain returns an in memory blockchain, only containing the genesis block.
func NewChain() []*Block {
	bc := make([]*Block, 5)
	bc = append(bc, genesis())

	return bc
}

// genHash generates and saves a hash for this block
func (b *Block) genHash() {
	// Init hash
	h := sha256.New()

	// Create string payload
	payload := b.timestamp.String() + b.data + string(b.index)

	// Write payload to hash
	h.Write([]byte(payload))

	// Save hash
	b.hash = h
}

// genesis returns the first Block in a blockchain.
func genesis() *Block {
	b := &Block{index: 0, timestamp: time.Now(), data: "First block", prevHash: nil}
	b.genHash()

	return b
}
