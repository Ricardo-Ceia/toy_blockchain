package main

import (
	"toy_blockchain/blockchain"
	"toy_blockchain/files_oper"
)

func main() {
	fileBytes := files_oper.ReadFromFile("./readme.txt")

	blockchain := blockchain.BlockChain{}

	blockchain.AddBlock(fileBytes)

	blockchain.Print()
}
