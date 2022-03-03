package main

import (
	"bytes"
	"crypto/sha256"
	"log"
)

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

type BlockChain struct {
	blocks []*Block
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}

	block.DeriveHash()
	return block
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)
}

func Genesis() *Block {
	return CreateBlock("Genesis - Init", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

func main() {
	chain := InitBlockChain()
	chain.AddBlock("First BLOCK after Genesis")
	chain.AddBlock("Second BLOCK after Genesis")
	chain.AddBlock("Third BLOCK after Genesis")
	chain.AddBlock("Fouth BLOCK after Genesis")

	for i, block := range chain.blocks {
		if i == 0 {
			log.Printf("Not contain Previous Hash")
		} else {
			log.Printf("Previous Block Hash: %x\n", block.PrevHash)
		}

		log.Printf("Data Block: %s\n", block.Data)
		log.Printf("Hash Block: %x\n\n", block.Hash)
	}
}
