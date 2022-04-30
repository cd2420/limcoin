package blockchain

import (
	"crypto/sha256"
	"fmt"
)

type blockchain struct {
	blocks []Block
}

var b *blockchain

func (b *blockchain) getLastHash() string {

	if len(b.blocks) > 0 {
		return b.blocks[len(b.blocks)-1].Hash
	}
	return ""

}

func (b *blockchain) AddBlock(data string) {
	newBlock := Block{data, "", b.getLastHash()}
	hash := sha256.Sum256([]byte(newBlock.Data + newBlock.PrevHash))
	newBlock.Hash = fmt.Sprintf("%x", hash)
	b.blocks = append(b.blocks, newBlock)
}

func (b *blockchain) ListBlocks() {
	for _, block := range b.blocks {
		fmt.Printf("Data : %s \n", block.Data)
		fmt.Printf("Hash : %s \n", block.Hash)
		fmt.Printf("Prev Hash : %s \n\n", block.PrevHash)
	}
}

func GetChain() *blockchain {
	b = &blockchain{}
	return b
}
