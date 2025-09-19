package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/dev-araujo/basic-blockchain-go/blockchain"
)

func main() {
	bc := blockchain.NewBlockchain()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("enter the data for a new block: ")
		data, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("\nA sair.")
				break
			}
			fmt.Printf("Erro ao ler a entrada: %v\n", err)
			break
		}

		if len(data) > 0 {
			bc.AddBlock([]byte(data))
			fmt.Println("====NEW===DATA==ADDED=========")
			fmt.Println("==============================")
			printInline(bc)
		}

	}

}

func printInline(bc *blockchain.Blockchain) {
	for _, block := range bc.Blocks {
		fmt.Printf("Index: %d\n", block.Index)
		fmt.Printf("Timestamp: %d\n", block.Timestamp)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println("==============================")
	}
}
