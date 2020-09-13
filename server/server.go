package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rafibarash/naivechain/node"
)

var (
	nodes = make(map[string]*node.Node, 0)
)

// RunHTTPServer runs the HTTP server to interact with nodes.
func RunHTTPServer() {
	r := gin.Default()

	r.GET("/nodes", getNodes)
	r.POST("/nodes", createNode)
	r.GET("/nodes/:nodeID/blocks", getNodeBlocks)
	r.POST("/nodes/:nodeID/blocks", createNodeBlock)

	log.Panic(r.Run(":8000"))
}

// getNodes gets information about all the nodes.
func getNodes(c *gin.Context) {
	b, err := json.Marshal(nodes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintln("Error occured while marshalling blockchain.")})
		return
	}
	c.JSON(http.StatusOK, gin.H{"payload": string(b)})
	return
}

// createNode creates a new node and adds to the webserver.
func createNode(c *gin.Context) {
	id := strconv.Itoa(len(nodes))
	n := node.New(id)
	nodes[id] = n
	b, err := json.Marshal(nodes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintln("Error occured while marshalling nodes.")})
		return
	}
	c.JSON(http.StatusOK, gin.H{"payload": string(b)})
	return
}

// getNodeBlocks gets a node's blockchain.
func getNodeBlocks(c *gin.Context) {
	id := c.Param("nodeID")
	n, ok := nodes[id]
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintln("Requesting blockchain from node that does not exist.")})
		return
	}
	b, err := json.Marshal(n.Blockchain)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintln("Error occured while marshalling blockchain.")})
		return
	}
	c.JSON(http.StatusOK, gin.H{"payload": string(b)})
	return
}

// createNodeBlock creates a new block on a node's blockchain
func createNodeBlock(c *gin.Context) {
	id := c.Param("nodeID")
	n, ok := nodes[id]
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintln("Requesting blockchain from node that does not exist.")})
		return
	}
	req := struct {
		Data string `form:"data" json:"data" xml:"data"  binding:"required"`
	}{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintln("Must include 'data' field when generating a new block.")})
		return
	}
	if err := n.GenBlock(req.Data); err != nil {
		log.Printf("Error occured while generating block: %q", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintln("Internal error while generating a new block.")})
		return
	}
	b, err := json.Marshal(nodes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintln("Error occured while marshalling nodes.")})
		return
	}
	c.JSON(http.StatusOK, gin.H{"payload": string(b)})
	return
}