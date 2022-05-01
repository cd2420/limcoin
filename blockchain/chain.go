package blockchain

import (
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

func (b *blockchain) AllBlocks() []*Block {
	return b.blocks
}

func (b *blockchain) AddBlock(data string) {
	b.blocks = append(b.blocks, createBlock(data))
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
			b.AddBlock("Genesis Block")
		})
	}
	return b
}
