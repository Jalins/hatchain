package main

import "hatchain/blockchain"

func main() {

	blockChain := blockchain.NewBlockChain()

	defer blockChain.DB.Close()

	cli := blockchain.CLI{blockChain}

	cli.Run()

}
