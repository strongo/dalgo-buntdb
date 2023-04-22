package testing4buntdb

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOpenInMemoryDB(t *testing.T) {
	db := OpenInMemoryDB(t)
	assert.NotNil(t, db)
}
