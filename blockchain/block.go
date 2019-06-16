package blockchain

import (
	"bytes"
	"encoding/gob"
	"log"
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
	block := NewBlock("Genesis block!",[]byte{})

	return block
}

//区块序列化
func (block *Block) Serialize() []byte  {
	// 开辟内存，存放字节集合
	var result bytes.Buffer
	// 创建编码对象
	encoder := gob.NewEncoder(&result)
	// 进行编码操作
	err := encoder.Encode(block)
	if err != nil{

		log.Panic(err)
	}

	return result.Bytes()
}

//区块反序列化
func DeSerializeBlock(blockBytes []byte) *Block  {
	// 用于存储字节转换的对象
	var block *Block
	// 创建解码对象，传入一个io读器
	dencoder := gob.NewDecoder(bytes.NewReader(blockBytes))
	// 进行解码操作
	err := dencoder.Decode(&block)
	if err != nil{

		log.Panic(err)
	}

	return block
}
