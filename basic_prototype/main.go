package main

import (
	"fmt"
	"hatchain/basic_prototype/blockchain"
)

func main() {
	block := blockchain.NewBlock("Genesis block!",[]byte{'t'})

	fmt.Println(block.Timestamp)
	fmt.Printf("%x\n",block.Hash)
	fmt.Printf("%x\n",block.PrevHash)
	fmt.Println(string(block.Data))

}
