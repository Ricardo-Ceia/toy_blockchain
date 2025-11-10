package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"time"
)

type Block struct {
	index        uint32
	timestamp    uint32
	value        []byte
	hash         []byte
	previousHash []byte
}

type BlockNode struct {
	block Block
	next  *BlockNode
}

type BlockChain struct {
	head *BlockNode
}

func computeHash(b Block) []byte {
	h := sha256.New()

	indexBytes := make([]byte, 4)
	timestampBytes := make([]byte, 4)

	binary.LittleEndian.PutUint32(indexBytes, b.index)
	binary.LittleEndian.PutUint32(timestampBytes, b.timestamp)

	h.Write(indexBytes)
	h.Write(timestampBytes)
	h.Write(b.value)
	h.Write(b.previousHash)

	return h.Sum(nil)
}

func CreateGenesisBlock(value []byte) Block {
	b := Block{
		index:        0,
		timestamp:    uint32(time.Now().Unix()),
		value:        value,
		previousHash: nil,
	}

	b.hash = computeHash(b)
	return b
}

func (l *BlockChain) AddBlock(value []byte) {
	var newBlock Block

	if l.head == nil {
		newBlock = CreateGenesisBlock(value)
		l.head = &BlockNode{block: newBlock}
		return
	}

	current := l.head
	for current.next != nil {
		current = current.next
	}
	newBlock = Block{
		index:        current.block.index + 1,
		timestamp:    uint32(time.Now().Unix()),
		value:        value,
		previousHash: current.block.hash,
	}
	newBlock.hash = computeHash(newBlock)

	current.next = &BlockNode{block: newBlock}
}

func (l *BlockChain) CheckValid() bool {
	prev := l.head
	current := l.head.next
	for current.next != nil {
		if !bytes.Equal(current.block.hash, computeHash(current.block)) {
			return false
		}
		if !bytes.Equal(current.block.previousHash, prev.block.hash) {
			return false
		}
		prev = current
		current = current.next
	}
	return true
}

func (l *BlockChain) Print() {
	current := l.head
	for current != nil {
		b := current.block
		fmt.Printf("Index: %d\n", b.index)
		fmt.Printf("Timestamp: %d\n", b.timestamp)
		fmt.Printf("Value: %s\n", string(b.value)) // decode bytes to string
		fmt.Printf("Hash: %x\n", b.hash)           // print hash as hex
		if b.previousHash != nil {
			fmt.Printf("Previous Hash: %x\n", b.previousHash)
		} else {
			fmt.Printf("Previous Hash: nil\n")
		}
		fmt.Println("---------------------")
		current = current.next
	}
}
