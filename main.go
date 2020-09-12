package main

import (
	"log"

	"github.com/rafibarash/naivechain/block"
)

func main() {
	blockchain := block.NewChain()
	log.Printf("%v", blockchain)
	return
}
