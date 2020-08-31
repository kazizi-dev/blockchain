/*
	References:
	https://stackoverflow.com/questions/5801008/go-and-operators
	https://yourbasic.org/golang/hash-md5-sha256-string-file/
	https://www.geeksforgeeks.org/how-to-use-strconv-itoa-function-in-golang/
*/

package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
)


type Block struct {
	PrevHash	[]byte
	Generation	uint64
	Difficulty 	uint8
	Data		string
	Proof		uint64
	Hash		[]byte
}


// Create new initial (generation 0) block.
func Initial(difficulty uint8) Block {
	blk := new(Block)

	blk.PrevHash = make([]byte, 32)
	blk.Generation = 0
	blk.Difficulty = difficulty
	blk.Data = ""
	blk.Proof = 0
	blk.Hash = nil

	return *blk
}

// Create new block to follow this block, with provided data.
func (prev_block Block) Next(data string) Block {
	newBlk := new(Block)

	newBlk.PrevHash = prev_block.Hash
	newBlk.Generation = prev_block.Generation + 1
	newBlk.Difficulty = prev_block.Difficulty
	newBlk.Data = data
	newBlk.Proof = 0
	newBlk.Hash = nil

	return *newBlk
}

// Calculate the block's hash.
func (blk Block) CalcHash() []byte {
	var separate = ":"
	str := hex.EncodeToString(blk.PrevHash) + separate +
		strconv.Itoa(int(blk.Generation)) + separate +
		strconv.Itoa(int(blk.Difficulty)) + separate +
		blk.Data + separate +
		strconv.Itoa(int(blk.Proof))

	hash := sha256.Sum256([]byte(str)[:])
	return hash[:]
}

// Check if the hash string has zero bits
func isZeroBits(str string, total int) bool {
	res := false

	// check to see how many zeros do we have
	size := 0
	for i := 0; i < total; i++ {
		if str[len(str)-1] == '0' {
			size++
			str = str[0 : len(str)-1]
		}
	}
	// check to see if the number of zeros match the total bits
	if size == total {
		res = true
	}

	return res
}

// Check if the bytes of the hash value are valid
func checkBytes(nBytes int, nBits uint, hash []byte) bool {
	result := false
	var lastByte uint8

	// if the total bytes is zero, check all bytes
	if nBytes != 0 {
		for i := 0; i < nBytes; i++ {
			if hash[len(hash)-i-1] != 0 {
				return false
			}
		}
	}

	// last is a byte slice (unsigned integer)
	lastByte = hash[len(hash) - nBytes-1]

	// if the unsigned integer is zer0
	if lastByte != 0 && lastByte % (1 << nBits) == 0 {
		// check if the last bits are zero
		result = isZeroBits(fmt.Sprintf("%b", lastByte), int(nBits))
	} else if lastByte == 0 && lastByte % (1 << nBits) == 0{
		result = true
	}
	return result
}

// check if the hash value is valid
func (blk Block) ValidHash() bool {
	if blk.Difficulty > 0 {
		var nBytes = int(blk.Difficulty / 8)
		var nBits = uint(blk.Difficulty % 8)
		var hash = blk.Hash

		return checkBytes(nBytes, nBits, hash)
	}
	return true
}

// Set the proof-of-work and calculate the block's "true" hash.
func (blk *Block) SetProof(proof uint64) {
	blk.Proof = proof
	blk.Hash = blk.CalcHash()
}

