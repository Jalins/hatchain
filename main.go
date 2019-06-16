package main

import (
	"fmt"
	"hatchain/blockchain"
	"math/big"
)

func main() {
	bc := blockchain.NewBlockChain()


	bc.AddBlock("send 9 BTC to Bob from Lily")
	bc.AddBlock("send 5 BTC to Bob from Lily")
	bc.AddBlock("send 7 BTC to Bob from Lily")

	for {
		var hashInt *big.Int

		blockChainIterator := bc.Iterator()

		fmt.Printf("%x\n", blockChainIterator.CurrentHash)

		blockChainIterator = blockChainIterator.Next()

		hashInt.SetBytes(blockChainIterator.CurrentHash)

		if hashInt.Cmp(big.NewInt(0)) == 0 {
			break
		}
	}


}
