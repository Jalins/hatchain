package blockchain

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

// 数据库的名字
const dbFile = "blockchain.db"

// 表命
const blocksBucket = "blocks "

// blockchain的数据结构，包括一个tip跟数据库
type BlockChain struct {
	Tip []byte   // 存储区块链中最后一个区块的hash值
	DB  *bolt.DB // 数据库对象
}

// 新增区块的方法
func (blockchain *BlockChain) AddBlock(data string) {
	// 创建一个新的区块
	newBlock := NewBlock(data, blockchain.Tip)

	// 更新数据库中表的数据
	err := blockchain.DB.Update(func(tx *bolt.Tx) error {
		// 通过Bucket获取表的数据
		b := tx.Bucket([]byte(blocksBucket))

		// 存储新的区块的数据
		err := b.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			log.Panic(err)
		}

		// 更新一下latest的值
		err = b.Put([]byte("latest"), newBlock.Hash)
		if err != nil {
			log.Panic(err)
		}

		// 将Tip更新为最新区块的hash值
		blockchain.Tip = newBlock.Hash
		return nil
	})
	if err != nil {
		log.Panic(err)
	}


}

// 创建一条带有创世区块的blockchain，使用数据库boltdb
func NewBlockChain() *BlockChain {
	// 定义一个Tip用于存储最后一个区块的hash
	var tip []byte

	// 打开数据库
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}

	// 更新数据
	// 1.首先是先判断表是否存在，如果不存在则创建，一张表代表一条链，所以每张表需要有一个创世区块
	err = db.Update(func(tx *bolt.Tx) error {
		// 判断数据库中是否存在这张表
		b := tx.Bucket([]byte(blocksBucket))

		// 如果b为nil，说明这张表不存在，则创建
		if b == nil {
			fmt.Println("区块链不存在，需重新创建！")
			// 2.创建创世区块
			genesis := GenesisBlock()
			// 创建表
			bucket, err := tx.CreateBucket([]byte(blocksBucket))
			if err != nil {
				log.Panic(err)
			}

			// 3.将创世区块序列化
			// 4.将创世区块的Hash作为key，区块的序列化数据作为value
			err = bucket.Put(genesis.Hash, genesis.Serialize())
			if err != nil {
				log.Panic(err)
			}

			// 5.设置一个key，nb，将hash作为value再次存储到数据里面
			err = bucket.Put([]byte("latest"), genesis.Hash)
			if err != nil {
				log.Panic(err)
			}
			tip = genesis.Hash
		} else {
			// nb是存储最后一个区块链hash的key
			tip = b.Get([]byte("latest"))
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	return &BlockChain{tip, db}
}
