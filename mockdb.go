package dalgo2buntdb

import (
	"github.com/dal-go/dalgo/dal"
	"github.com/dal-go/dalgo2buntdb/testing4buntdb"
	"github.com/tidwall/buntdb"
	"testing"
)

// NewInMemoryMockDB creates a new dal.Database as a wrapper of in-memory BuntDB instance
func NewInMemoryMockDB(t *testing.T) dal.DB {
	db, err := buntdb.Open(testing4buntdb.MemoryPath)
	if err != nil {
		t.Fatalf("Failed to create a new temporary in-memory mock BuntDB database: %v", err)
	}
	return NewDatabase(db)
}
