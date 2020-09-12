package node

import (
	"fmt"

	"github.com/rafibarash/naivechain/block"
)

type Node struct {
	ID         string
	Blockchain block.Blockchain
}

func (n Node) String() string {
	return fmt.Sprintf("{\"id\": %q, {\"blockchain\": %v}", n.ID, n.Blockchain)
}

func New(id string) *Node {
	return &Node{ID: id, Blockchain: block.NewChain()}
}
