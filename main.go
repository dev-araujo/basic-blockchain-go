package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/dev-araujo/basic-blockchain-go/blockchain"
)

func main() {
	bc := blockchain.NewBlockchain()
	reader := bufio.NewReader(os.Stdin)

	for {
		printInline(bc)
		fmt.Println("enter the data for a new block: ")
		data, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("\n exit")
				break
			}
			fmt.Printf("error %v\n", err)
			break
		}

		trimmedData := strings.TrimSpace(data)
		if len(trimmedData) > 0 && trimmedData != "/exit" {
			bc.AddBlock([]byte(trimmedData))
			fmt.Println("====NEW===DATA==ADDED=========")
			printInline(bc)
		}
		if trimmedData == "/exit" {
			break
		}

	}

}

func printInline(bc *blockchain.Blockchain) {
	for _, block := range bc.Blocks {
		fmt.Println("==============================")
		fmt.Printf("Index: %d\n", block.Index)
		fmt.Printf("Timestamp: %d\n", block.Timestamp)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println("==============================")
	}
}
