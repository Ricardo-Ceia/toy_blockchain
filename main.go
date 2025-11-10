package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
)

type Block struct {
	index        uint32
	timestamp    uint32
	value        uint32
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
	valueBytes := make([]byte, 4)

	binary.LittleEndian.PutUint32(indexBytes, b.index)
	binary.LittleEndian.PutUint32(timestampBytes, b.timestamp)
	binary.LittleEndian.PutUint32(valueBytes, b.value)

	h.Write(indexBytes)
	h.Write(timestampBytes)
	h.Write(valueBytes)
	h.Write(b.previousHash)

	return h.Sum(nil)
}

func createGenesisBlock() Block {
	b := Block{
		index:        0,
		timestamp:    0,
		value:        0,
		previousHash: nil,
	}

	b.hash = computeHash(b)
	return b
}

func (l *BlockChain) addBlock(newBlock Block) {
	newBlockNode := &BlockNode{block: newBlock}

	if l.head == nil {
		l.head = newBlockNode
		return
	}

	current := l.head

	for current.next != nil {
		current = current.next
	}

	current.next = newBlockNode
}

func (l *BlockChain) checkValid() bool {
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
