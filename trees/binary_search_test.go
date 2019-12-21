package tree

import (
	"encoding/binary"
	"testing"
)

func TestBinarySearchTreeAdd(t *testing.T) {
	b := NewBinarySearchTree()
	tmp := make([]byte, 4)
	for i:=0; i<100; i++ {
		binary.LittleEndian.PutUint32(tmp, uint32(i))
		b.Add(tmp, i)
	}

	for i:=0; i<100; i++ {
		binary.LittleEndian.PutUint32(tmp, uint32(i))
		value, err := b.Get(tmp)
		if err != nil {
			t.Fatalf("Expected nil. Got: %v", err)
		}
		if value.(int) != i {
			t.Fatalf("Expected %v. Got: %v", value, i)
		}
	}
}
