package tree

import (
	"errors"
	"github.com/cespare/xxhash"
)

var ErrKeyNotFound = errors.New("key not found")

type Node struct {
	Left  *Node
	Right *Node
	HKey  uint64
	Key   []byte
	Value interface{}
}

type BinarySearchTree struct {
	root *Node
}

func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{
		root: &Node{},
	}
}

func (b *BinarySearchTree) Add(key []byte, value interface{}) {
	hkey := xxhash.Sum64(key)
	b.add(hkey, key, value, b.root)
}

func (b *BinarySearchTree) add(hkey uint64, key []byte, value interface{}, node *Node) {
	if node.HKey == 0 {
		// Root node
		node.Key = key
		node.Value = value
		node.HKey = hkey
		return
	}

	if node.HKey > hkey {
		if node.Left != nil {
			b.add(hkey, key, value, node.Left)
			return
		}
		node.Left = &Node{
			HKey:  hkey,
			Key:   key,
			Value: value,
		}
	} else {
		if node.Right != nil {
			b.add(hkey, key, value, node.Right)
			return
		}
		node.Right = &Node{
			HKey:  hkey,
			Key:   key,
			Value: value,
		}
	}
}

func (b *BinarySearchTree) Get(key []byte) (interface{}, error) {
	hkey := xxhash.Sum64(key)
	return b.get(hkey, b.root)
}

func (b *BinarySearchTree) get(hkey uint64, node *Node) (interface{}, error) {
	if node.HKey == hkey {
		return node.Value, nil
	}

	if node.HKey > hkey {
		if node.Left != nil {
			return b.get(hkey, node.Left)
		}
		return nil, ErrKeyNotFound
	}

	if node.Right != nil {
		return b.get(hkey, node.Right)
	}
	return nil, ErrKeyNotFound
}
