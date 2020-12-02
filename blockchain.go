package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type Block struct {
	hash     []byte
	data     []byte
	prevHash []byte
}

type BlockChain struct {
	blockchain []*Block
}

func (a *Block) getHash() {
	partA := [][]byte{a.data, a.prevHash}
	partB := []byte{}
	incHash := bytes.Join(partA, partB)
	compHash := sha256.Sum256(incHash)
	a.hash = compHash[:]
}

func makeBlock(Data string, PrevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(Data), PrevHash}
	block.getHash()
	return block
}

func Genesis() *Block {
	return makeBlock("First Block", []byte{})
}

func (b *BlockChain) addBlock(Data string) {
	lastBlock := b.blockchain[len(b.blockchain)-1]
	newBlock := makeBlock(Data, lastBlock.hash)
	b.blockchain = append(b.blockchain, newBlock)
}

func makeBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

func main() {
	bChain := makeBlockChain()
	bChain.addBlock("Second Block")
	bChain.addBlock("Third Block")

	for _, block := range bChain.blockchain {
		fmt.Printf("Prev Hash : %x", block.prevHash)
		fmt.Printf("\nData : %s", block.data)
		fmt.Printf("\nHash : %x\n", block.hash)
	}
}
