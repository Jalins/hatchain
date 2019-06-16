package blockchain

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type CLI struct {
 	BlockChain *BlockChain
 }

func (cli *CLI) Run()  {

	addBlockCmd := flag.NewFlagSet("addBlock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("addPrint", flag.ExitOnError)

	addBlockData := addBlockCmd.String("data", "", "Blcok data")

	if len(os.Args) < 1 {
		switch os.Args[1] {
		case "addBlock":
			err := addBlockCmd.Parse(os.Args[2:])
			if err != nil {
				log.Panic(err)
			}
		case "printChain":
			err := printChainCmd.Parse(os.Args[2:])
			if err != nil {
				log.Panic(err)
			}
		default:
			flag.Usage()
			os.Exit(1)

		}
	}else {
		flag.Usage()
	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			os.Exit(1)
		}

		cli.BlockChain.AddBlock(*addBlockData)
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}

}

func (cli *CLI) addBlock(data string) {
	cli.BlockChain.AddBlock(data)
	fmt.Println("Success!")
}

func (cli *CLI) printChain() {
	bci := cli.BlockChain.Iterator()

	for {
		block := bci.Next()

		fmt.Printf("Prev. hash: %x\n", block.PrevHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
		if len(block.PrevHash) == 0 {
			break
		}
	}
}
