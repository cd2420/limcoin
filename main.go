package main

import "github.com/cd2420/limcoin/blockchain"

func main() {
	chain := blockchain.GetChain()
	chain.AddBlock("My First Block!!!!")
	chain.AddBlock("My Second Block!!!!")
	chain.AddBlock("My Third Block!!!!")
	chain.ListBlocks()
}
