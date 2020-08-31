# blockchain-go
A blockchain program that calculates and processes hash codes.

## Features:

**work_queue.go**
- Manages work queue to organize tasks

**block.go**
- Creates a new block 
- Calculates the block's hash
- Verify if hash is valid

**blockchain.go**
- Adds blocks to exisiting chain
- Checks if a block is valid
- Displays error message to the user if there is an error with adding the block or checking validity of block  

**blockchain_test.go**
- Contains test cases to test the functionality of the blocks in the blockchain

**mining.go**
- Mines the blocks in the blockchain


## Setup:
- Install Golang
- Clone this repository
- Enter blockchain-go directory
		cd blockchain-go
- Set GOPATH
		 export GOPATH=`pwd`
- To test Work Queue
         go test work_queue -v
- To test Blockchain
         go test blockchain -v 
