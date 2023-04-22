package testing4buntdb

import (
	"github.com/tidwall/buntdb"
	"testing"
)

const MemoryPath = ":memory:"

func OpenInMemoryDB(t *testing.T) *buntdb.DB {
	db, err := buntdb.Open(MemoryPath)
	if err != nil {
		t.Fatal(err)
	}
	return db
}
