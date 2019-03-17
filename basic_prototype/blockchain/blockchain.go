package blockchain

// blockchain的数据结构
type BlockChain struct {
	// 一个可以存储多个block的数组
	Blocks  []*Block
}

// 创建的这条链可以添加新的区块
func (blockchain *BlockChain) AddBlock(data string) {
	// 首先获取新的区块
	prevBlockHash := blockchain.Blocks[len(blockchain.Blocks) - 1].Hash
	newBlock := NewBlock(data, prevBlockHash)

	// 将区块添加到区块数组中去
	blockchain.Blocks = append(blockchain.Blocks, newBlock)

}

// 创建一条带有创世区块的blockchain
func NewBlockChain() *BlockChain {
	return &BlockChain{[]*Block{GenesisBlock()}}
}