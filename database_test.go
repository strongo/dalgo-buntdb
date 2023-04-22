package dalgo2buntdb

import (
	"github.com/dal-go/dalgo2buntdb/testing4buntdb"
	"github.com/tidwall/buntdb"
	"testing"
)

func TestNewDatabase(t *testing.T) {
	db, err := buntdb.Open(testing4buntdb.MemoryPath)
	if err != nil {
		t.Fatalf("direct buntdb.Open(%s) returned error: %v", testing4buntdb.MemoryPath, err)
	}
	var dtb = NewDatabase(db)
	if dtb == nil {
		t.Error("NewDatabase returned nil")
	}
}
