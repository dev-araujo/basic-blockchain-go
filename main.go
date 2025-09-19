package main

import (
	"fmt"

	"github.com/dev-araujo/basic-blockchain-go/blockchain"
)

func main() {
	bc := blockchain.NewBlockchain()
	bc.AddBlock([]byte("First Block after Genesis"))
	bc.AddBlock([]byte("Second Block after Genesis"))

	for _, block := range bc.Blocks {
		fmt.Printf("Index: %d\n", block.Index)
		fmt.Printf("Timestamp: %d\n", block.Timestamp)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println("==============================")
	}

}
