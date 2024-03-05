package merkledag

import (
	"hash"
)

func Add(store KVStore, node Node, h hash.Hash) []byte {
	if node.Type == "file" {
		store.Save([]byte("file"), node.Data)
		return hash(node.Data, h)
	} else if node.Type == "dir" {
		for _, child := range getDirContents(node) {
			Add(store, child, h)
		}
		return hashDirContents(node, h)
	}
	return nil 
}

func hash(data []byte, h hash.Hash) []byte {
	h.Reset()
	h.Write(data)
	return h.Sum(nil)
}

func getDirContents(dir Node) []Node {
	return []Node{
		{Type: "file", Data: []byte("file1")},
		{Type: "file", Data: []byte("file2")},
		{Type: "dir", Data: nil},
	}
}

func hashDirContents(dir Node, h hash.Hash) []byte {
	for _, child := range getDirContents(dir) {
		childHash := Add(store, child, h)
		h.Write(childHash)
	}
	return h.Sum(nil)
}
