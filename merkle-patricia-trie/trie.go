package merkle_patricia_trie

import (
	"go/ast"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethdb"
)

type Trie struct {
	db *ethdb.Database
	root ast.Node
	originalRoot common.Hash
}

func New(root common.Hash, db *ethdb.Database) (*Trie, error) {
	if db == nil {
		panic("trie.New called without a database")
	}
	trie := &Trie{
		db:           db,
		originalRoot: root,
	}
	return trie, nil
}