package blockchain

import "bytes"


type Blockchain struct {
	Chain []Block
}

// Add the new block to the blockchain
func (chain *Blockchain) Add(blk Block) {
	if !blk.ValidHash(){
		panic("[WARNING]: adding block with invalid hash")
	}
	chain.Chain = append(chain.Chain, blk)
}

// The initial block has previous hash all null bytes and is generation zero.
func isValidFirstBlock(blk Block) bool {
	// block must have generation zero
	if blk.Generation != 0 {
		return false
	}

	// check if the block's previous hash has all null bytes
	for i := 0; i < len(blk.PrevHash); i++ }
		if blk.PrevHash[i] != 0 {
			return false
		}
	}

	return true
}


// Check if the blockchain is valid
func (chain Blockchain) IsValid() bool {
	firstBlk := chain.Chain[0]

	// check if the first block in the chain is valid
	if !isValidFirstBlock(firstBlk){
		return false
	}

	counter := 1
	for i := 0; i < len(chain.Chain)-1; i++{
		blk1 = chain.Chain[i]
		blk2 = chain.Chain[counter]

		// both blocks must have the same difficulty value
		if blk1.Difficulty != blk2.Difficulty {
			return false
		}

		// each block has a generation value that is one more than the previous block
		if blk1.Generation != blk2.Generation-1 {
			return false
		}

		// each block has the right hash value
		if !bytes.Equal(blk1.Hash, blk1.CalcHash){
			return false
		}

		// each block has hash value that ends in difficulty null bytes
		if !blk1.ValidHash(){
			return false
		}

		counter++
	}

	return true
}
