/*
	References:
	https://medium.com/@mycoralhealth/code-your-own-blockchain-mining-algorithm-in-go-82c6a71aba1f
*/

package blockchain

import "work_queue"

type miningWorker struct {
	blk 	Block
	start 	uint64
	end 	uint64
}

// code for work_queue.Worker
func (mWorker miningWorker) Run() interface{} {
	miningResult := new(MiningResult)
	miningResult.Found = false

	// for each worker, create the block's hash
	for i := mWorker.start; i <= mWorker.end; i++ {
		mWorker.blk.Proof = i
		mWorker.blk.Hash = mWorker.blk.CalcHash()

		// if the worker's block has a valid hash
		if mWorker.blk.ValidHash(){
			miningResult.Proof = i
			miningResult.Found = true
			return *miningResult
		}
	}

	return *miningResult
}

type MiningResult struct {
	Proof	uint64
	Found 	bool
}


/*
	Mine the range of proof values, by breaking up into chunks and checking
	workers" chunks concurrently in a work queue. Should return shortly after
	a result is found.
*/
func (blk Block) MineRange(start uint64, end uint64, workers uint64, chunks uint64) MiningResult {
	chunkSize := (end - start) / chunks
	if chunkSize == 0 {
		chunkSize = end
	}

	workQueue := work_queue.Create(uint(workers), uint(chunks))
    for i := start; i < end; i += chunkSize {
        mWorker := new(miningWorker)
		mWorker.start = i

        if end < i + (chunkSize-1){
            mWorker.end = end
        } else{
            mWorker.end = i + (chunkSize-1)
		}

        mWorker.blk = blk
        workQueue.Enqueue(mWorker)
    }

	mResult := new(MiningResult)
	mResult.Found = false
	var status = true
	for status {
		r := <- workQueue.Results
		if r.(MiningResult).Found == true {
			mResult.Found = true
			mResult.Proof = r.(MiningResult).Proof
			workQueue.Shutdown()
			status = false
		}
	}

	return *mResult
}


/*
	Call .MineRange with some reasonable values that will probably find a result.
	Good enough for testing at least. Updates the block's .Proof and .Hash if successful.
*/
func (blk *Block) Mine(workers uint64) bool {
	reasonableRangeEnd := uint64(4 * 1 << blk.Difficulty) // 4 * 2^(bits that must be zero)
	mr := blk.MineRange(0, reasonableRangeEnd, workers, 4321)
	if mr.Found {
		blk.SetProof(mr.Proof)
	}
	return mr.Found
}
