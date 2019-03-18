package blockchain

import (
	"fmt"
	"math/big"
)

type ProofOfWork struct {
	block *Block
	target *big.Int
}

// 定义一个最大值等于256难度值，难度值越小，则计算量越大
const targetBits  = 24

// 工厂方法，返回一个pow对象，包含一个区块跟一个难度值
func NewProofOfWork(block *Block) *ProofOfWork{
	// 第一位数为1
	target := big.NewInt(1)
	target.Lsh(target, (256 - targetBits))
	fmt.Println(target)
	return &ProofOfWork{block, target}
}

func (pow *ProofOfWork) Run(){

}