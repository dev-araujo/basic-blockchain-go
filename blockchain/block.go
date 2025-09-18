package blockchain

import (
	"bytes"
	"crypto/sha256"
	"strconv"
)

type Block struct {
	Index         uint64
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

func (b *Block) calculateHash() []byte {
	timestampBytes := strconv.FormatInt(b.Timestamp, 10)
	indexBytes := strconv.FormatUint(b.Index, 10)

	info := [][]byte{[]byte(timestampBytes), []byte(indexBytes), b.Data, b.PrevBlockHash}
	joinedInfo := bytes.Join(info, []byte{})

	hash := sha256.Sum256(joinedInfo)
	return hash[:]
}
