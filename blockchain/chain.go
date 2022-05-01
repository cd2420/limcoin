package blockchain

import (
	"fmt"
	"sync"
)

type blockchain struct {
	blocks []*Block
}

var b *blockchain
var once sync.Once

func (b *blockchain) getLastHash() string {
	totalBlocks := len(GetChain().blocks)
	if totalBlocks == 0 {
		return ""
	}

	return GetChain().blocks[totalBlocks-1].Hash
}

func (b *blockchain) ListBlocks() {
	for _, block := range b.blocks {
		fmt.Printf("Data : %s \n", block.Data)
		fmt.Printf("Hash : %s \n", block.Hash)
		fmt.Printf("Prev Hash : %s \n\n", block.PrevHash)
	}
}

func createBlock(data string) *Block {
	newBlock := &Block{data, "", b.getLastHash()}
	newBlock.calculateHash()
	return newBlock

}

func GetChain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{}
			b.blocks = append(b.blocks, createBlock("Genesis Block"))
		})
	}
	return b
}
