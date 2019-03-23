package blockchain

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// 区块的数据结构
type Block struct {

	// 时间戳
	Timestamp int64
	// 前一个hash
	PrevHash []byte
	// 交易数据
	Data []byte
	// 当前hash
	Hash []byte
	// 随机数
	Nonce int
}

// 在区块中设置当前的hash值
func (block *Block) setHash(){
	// 将int64的时间戳转化为字节数组，这里的第二个参数的意思是转化为几进制，可选范围从2到36，FormatInt返回的是一个字符串
	timestamp := []byte(strconv.FormatInt(block.Timestamp, 10))

	// 使用bytes.Join方法对各个字节数组进行拼接，第二个参数的意思是切割符，例如[]byte(",")，以逗号来切割
	header := bytes.Join([][]byte{timestamp, block.PrevHash, block.Data},[]byte{})

	// 对得到的字节数据进行hash运算
	hash := sha256.Sum256(header)
	block.Hash = hash[:]

}

// 创建新的区块，工厂方法
func NewBlock(data string, prevHash []byte) *Block {

	block := &Block{time.Now().Unix(), prevHash, []byte(data), []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce
	//block.setHash()
	return block
}


func  GenesisBlock() *Block{
	block := NewBlock("Genesis block!",[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0})

	return block
}
