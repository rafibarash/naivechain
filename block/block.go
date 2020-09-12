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
	timestamp time.Time
	hash      hash.Hash
	prevHash  hash.Hash
	Data      string
}

// Blockchain is a slice of Blocks.
type Blockchain []*Block

func (b Block) String() string {
	return fmt.Sprintf("{\"Index\": %d, {\"Data\": %q}", b.Index, b.Data)
}

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
// - The previous block's hash is equal to the new block's previous hash.
// - The new block's hash member variable must equal the result of actually calculating the block's hash.
func isValidBlock(b *Block, pb *Block) bool {
	return b.Index-1 == pb.Index && b.prevHash == pb.hash && b.hash == b.genhash()
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
	b := &Block{Index: i, timestamp: time.Now(), Data: d, prevHash: ph}
	b.genhash()

	return b
}

// NewChain returns an in memory blockchain, only containing the genesis block.
func NewChain() Blockchain {
	bc := make(Blockchain, 0)
	bc = append(bc, genesis())

	return bc
}

// genhash generates and saves a hash for this block
func (b *Block) genhash() hash.Hash {
	// Init hash
	h := sha256.New()

	// Create string payload
	payload := b.timestamp.String() + b.Data + string(b.Index)

	// Write payload to hash
	h.Write([]byte(payload))

	return h
}

// genesis returns the first Block in a blockchain.
func genesis() *Block {
	b := &Block{Index: 0, timestamp: time.Now(), Data: "First block", prevHash: nil}
	b.hash = b.genhash()

	return b
}
