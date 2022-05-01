package main

import "github.com/cd2420/limcoin/blockchain"

func main() {
	chain := blockchain.GetChain()
	chain.ListBlocks()
}
