package node

import (
	"github.com/rafibarash/naivechain/block"
)

type Node struct {
	ID         string
	Blockchain block.Blockchain
}

func New(id string) *Node {
	n := &Node{ID: id, Blockchain: block.NewChain()}
	// TODO: Add node to websocket
	// TODO: Sync node's blockchain with other nodes
	return n
}

func (n *Node) GenBlock(data string) error {
	b := n.Blockchain.NewBlock(data)
	bc, err := n.Blockchain.AddBlock(b)
	if err != nil {
		return err
	}
	n.Blockchain = bc
	return nil
	// TODO: Broadcast new blockchain to all nodes, update blockchain accordingly.
}
