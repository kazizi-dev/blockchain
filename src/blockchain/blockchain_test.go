package blockchain

import (
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"testing"
)

// test calculate hash
// reference: prof from coursys
func TestForCalculateHash16(t *testing.T) {
	block1 := Initial(16)
	block1.Mine(1)
	assert.Equal(t, 
		hex.EncodeToString(block1.CalcHash()), 
		"6c71ff02a08a22309b7dbbcee45d291d4ce955caa32031c50d941e3e9dbd0000")

	block2 := block1.Next("message")
	block2.Mine(1)
	assert.Equal(t, 
		hex.EncodeToString(block2.CalcHash()), 
		"9b4417b36afa6d31c728eed7abc14dd84468fdb055d8f3cbe308b0179df40000")
}

// test valid hashes
// reference: prof from coursys
func TestForValidHash16(t *testing.T) {
	block1 := Initial(19)
	block1.SetProof(87745)
	block2 := block1.Next("hash example 1234")

	block2.SetProof(1407891)
	assert.Equal(t, block1.ValidHash(), true)

	block2.SetProof(346082)
	assert.Equal(t, block2.ValidHash(), false)
}

// test difficulty of 7
// reference: prof from coursys
func TestDifficulty7(t *testing.T) {
	block1 := Initial(7)
	block1.Mine(1)
	assert.Equal(t, 
		hex.EncodeToString(block1.Hash), 
		"379bf2fb1a558872f09442a45e300e72f00f03f2c6f4dd29971f67ea4f3d5300")
	assert.Equal(t, block1.Proof, uint64(385))
		

	block2 := block1.Next("this is an interesting message")
	block2.Mine(1)
	assert.Equal(t, 
		hex.EncodeToString(block2.Hash), 
		"4a1c722d8021346fa2f440d7f0bbaa585e632f68fd20fed812fc944613b92500")
	assert.Equal(t, block2.Proof, uint64(20))


	block3 := block2.Next("this is not interesting")
	block3.Mine(1)
	assert.Equal(t, 
		hex.EncodeToString(block3.Hash), 
		"ba2f9bf0f9ec629db726f1a5fe7312eb76270459e3f5bfdc4e213df9e47cd380")
	assert.Equal(t, block3.Proof, uint64(40))
}


// test difficulty of 20
// reference: prof from coursys
func TestDifficulty20(t *testing.T) {
	block1 := Initial(20)
	block1.Mine(1)
	assert.Equal(t, 
		hex.EncodeToString(block1.Hash), 
		"19e2d3b3f0e2ebda3891979d76f957a5d51e1ba0b43f4296d8fb37c470600000")
	assert.Equal(t, block1.Proof, uint64(1209938))
		

	block2 := block1.Next("this is an interesting message")
	block2.Mine(1)
	assert.Equal(t, 
		hex.EncodeToString(block2.Hash), 
		"a42b7e319ee2dee845f1eb842c31dac60a94c04432319638ec1b9f989d000000")
	assert.Equal(t, block2.Proof, uint64(989099))


	block3 := block2.Next("this is not interesting")
	block3.Mine(1)
	assert.Equal(t, 
		hex.EncodeToString(block3.Hash), 
		"6c589f7a3d2df217fdb39cd969006bc8651a0a3251ffb50470cbc9a0e4d00000")
	assert.Equal(t, block3.Proof, uint64(1017262))
}

