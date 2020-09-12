package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rafibarash/naivechain/block"
	"github.com/rafibarash/naivechain/node"
)

var (
	nodes []*node.Node
)

func main() {
	blockchain := block.NewChain()
	log.Printf("%v", blockchain)
	initServer()
	return
}

func initServer() {
	r := gin.Default()
	r.GET("/blocks/:nodeID", func(c *gin.Context) {
		id := c.Param("nodeID")
		log.Printf("%s", id)

	})

	log.Panic(r.Run(":8000"))
}
