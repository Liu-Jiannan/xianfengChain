package main

import (
	"./chain"
	"fmt"
)

func main() {
	fmt.Println("hello world")

	blockchain := chain.CreateChainWithGensis([]byte("Hello world"))
	blockchain.CreateNewBlock([]byte("hello"))

	fmt.Println("区块的个数,",len(blockchain.Blocks))

	fmt.Println("区块0的哈希值:",blockchain.Blocks[0])
	fmt.Println("区块1的哈希值:",blockchain.Blocks[1])
}