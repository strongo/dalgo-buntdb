package dalgo2buntdb

import (
	"github.com/dal-go/dalgo/dal"
	"github.com/tidwall/buntdb"
)

var _ dal.Reader = (*buntdbReader)(nil)

type buntdbReader struct {
	//i     int // iteration
	query dal.Query
	tx    *buntdb.Tx
	//iterator *datastore.Iterator
}

func (reader buntdbReader) Close() error {
	return nil
}

func (reader buntdbReader) Next() (record dal.Record, err error) {
	return nil, dal.ErrNotImplementedYet
}

func (reader buntdbReader) Cursor() (string, error) {
	return "", dal.ErrNotSupported
}
