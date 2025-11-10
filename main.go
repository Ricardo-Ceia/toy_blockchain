package main

import (
	"fmt"
	"strconv"
	"toy_blockchain/blockchain"
	"toy_blockchain/files_oper"
)

func main() {
	fileBytes := files_oper.ReadFromFile("./readme.txt")

	blockchain := blockchain.BlockChain{}

	blockchain.AddBlock(fileBytes)
	path := "./test"

	for i := range 3 {
		path = path + strconv.Itoa(i) + ".txt"
		fileBytes := files_oper.ReadFromFile(path)
		blockchain.AddBlock(fileBytes)
		path = "./test"
	}
	blockchain.TemperBlock(2, []byte("i have tried to hack this but failed"))
	fmt.Println(blockchain.CheckValid())
	blockchain.Print()
}
