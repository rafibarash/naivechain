# Naivechain

**The notes contained here are 100% based off [Lauri Hartikka's implementation](https://medium.com/@lhartikk/a-blockchain-in-200-lines-of-code-963cc1cc0e54#.dttbm9afr5).** The golang code is directly based off these notes.

## Notes

### Block structure

A simple block will keep track of index, timestamp, data, hash, and the previous hash.

### Block hash

The block is hashed to keep integrity of the data. An SHA-256 is taken over the content of the block. This does not have to do with mining or proof of work.

### Generating a block

For this we need to know the hash of the previous block and create the rest of the required content (index, timestamp, data, hash). Block data is provided by the end-user.

### Storing the blocks

A database could be used to persis the blockchain, but here we will just use an in-memory array. The first block is always what's called a "genesis-block", and this is hard coded.

### Validating the integrity of blocks

The integrity of a block should always be validated when we receive new blocks from other nodes and must decide whether or not to accept them. This means the following must be true:

- The previous block's index is one below the new block's index.
- The previous block's hash is equal to the new block's "previousHash"
- The new block's hash member variable must equal the result of actually calculating the block's hash

### Choosing the longest chain

When deciding to replace a blockchain, check that it is a valid chain and replace if the new blockchain length is greater than the current one.

### Communicating with other nodes

One of the most important jobs of a node is to share and sync the blockchain with other nodes. The following rules keep a network in sync:

- When a node generates a new block, it broadcasts it to the network.
- When a node connects to a new peer it querys for the latest block.
- When a node encounters a block with an index larger than the current known block, it either adds the block to its current chain or querys for the full blockchain

### Controlling the node

The user can control the node through an HTTP server. A user should be able to interact with the node in the following ways:

- List all blocks.
- Create a new block with user content
- List or add peer nodes

### Architecture

A node should expose two web servers. One for the user to control the node (HTTP server) and one for peer-to-peer communication between nodes (Websocket HTTP server).
