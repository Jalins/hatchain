package main

import (
	"fmt"
	"hatchain/blockchain"
)

func main() {
	bc := blockchain.NewBlockChain()

	//bc.AddBlock("send 9 BTC to Bob from Lily")
	//bc.AddBlock("send 4.8 BTC to Lily from Tom")
	//bc.AddBlock("send 7.9 BTC to Tom from Tony")
	fmt.Printf("bc: %x\n", bc)

	fmt.Printf("Tip: %x\n", bc.Tip)

	bc.AddBlock("send 9 BTC to Bob from Lily")
	fmt.Printf("Tip: %x\n", bc.Tip)
}
