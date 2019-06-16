package blockchain

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"hatchain/utils"
	"math"
	"math/big"
)

type ProofOfWork struct {
	block *Block
	target *big.Int
}

var maxNonce = math.MaxInt64

// 定义一个最大值等于256难度值，targetBits越大表示前导零越多也就意味这个计算量越大
const targetBits  = 18

// 工厂方法，返回一个pow对象，包含一个区块跟一个目标值
func NewProofOfWork(block *Block) *ProofOfWork{
	// 设置一个为1的目标数
	target := big.NewInt(1)
	target.Lsh(target, (256 - targetBits))

	return &ProofOfWork{block, target}
}

func (pow *ProofOfWork) Run() (int, []byte){
	var hashInt big.Int
	var hash [32]byte
	nonce := 0
	fmt.Printf("开始挖矿：%s\n", pow.block.Data)

	// 在0与最大随机数之间进行随机计算
	for nonce < maxNonce  {

		data := pow.perpareData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		// 将得到的hash（byte类型）值转化为big.Int类型，至于为什么使用big.Int类型是因为这个既可以存储大容量的int数据，
		// 又提供了方便的工具函数可以对数据进行处理，如SetBytes和Cmp函数
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			break
		}else {
			nonce++
		}

	}
	fmt.Println("\n\n")
	return nonce, hash[:]
}

func (pow ProofOfWork) perpareData(nonce int) []byte{
	data := bytes.Join([][]byte{
		pow.block.PrevHash,
		pow.block.Data,
		utils.IntToHex(pow.block.Timestamp),
		utils.IntToHex(int64(targetBits)),
		utils.IntToHex(int64(nonce)),
	}, []byte{},)

	return data
}