package end2end

import (
	"errors"
	end2end "github.com/dal-go/dalgo-end2end-tests"
	"github.com/dal-go/dalgo2buntdb"
	"github.com/dal-go/dalgo2buntdb/testing4buntdb"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEndToEnd(t *testing.T) {
	db := dalgo2buntdb.NewDatabase(testing4buntdb.OpenInMemoryDB(t))
	assert.NotNil(t, db)
	end2end.TestDalgoDB(t, db, errors.New("problems with RLock during opening read transaction + reader is not implemented yet"), false)
}
