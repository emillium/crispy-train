package chain

func NewBlockHash(block *Block) string {
	blockInfo := block.Timestamp + block.PrevHash + block.Hash + block.ValidatorAddr
	return NewHash(blockInfo)
}
