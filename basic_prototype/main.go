package main

import (
	"fmt"
	"hatchain/basic_prototype/blockchain"
	"time"
)

func main() {
	bc := blockchain.NewBlockChain()

	bc.AddBlock("send 9 BTC to Bob from Lily")
	bc.AddBlock("send 4.8 BTC to Lily from Tom")
	bc.AddBlock("send 7.9 BTC to Tom from Tony")

	for _, block := range bc.Blocks {
		fmt.Printf("%x\n", block.PrevHash)
		fmt.Printf("%s\n", string(block.Data))
		fmt.Printf("%x\n", block.Hash)
		fmt.Printf("%s\n", time.Unix(block.Timestamp,0).Format("2006-01-03 03:04:05 PM"))

		fmt.Println()
	}
}
