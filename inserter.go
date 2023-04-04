package dalgo2buntdb

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/dal-go/dalgo/dal"
	"github.com/tidwall/buntdb"
)

// ErrKeyAlreadyExists an error to be used in insert when generated key already exists
var ErrKeyAlreadyExists = errors.New("key already exists")

func (dtb database) Insert(ctx context.Context, record dal.Record, opts ...dal.InsertOption) error {
	return dtb.db.Update(func(tx *buntdb.Tx) error {
		t := transaction{tx: tx}
		return t.Insert(ctx, record, opts...)
	})
}

func (t transaction) Insert(ctx context.Context, record dal.Record, opts ...dal.InsertOption) (err error) {
	options := dal.NewInsertOptions(opts...)

	if generateID := options.IDGenerator(); generateID != nil {
		err = t.insertWithGenerator(ctx, generateID, record)
	} else {
		err = t.insert(record)
	}
	return err
}

func (t transaction) insertWithGenerator(ctx context.Context, generateID dal.IDGenerator, record dal.Record) (err error) {
	for i := 0; i < 10; i++ {
		if err = generateID(ctx, record); err != nil {
			break
		}
		if err = t.insert(record); err != nil {
			if err == ErrKeyAlreadyExists {
				continue
			}
			break
		}
	}
	return err
}

func (t transaction) insert(record dal.Record) error {
	key := record.Key()
	k := key.String()
	if _, err := t.tx.Get(k); err == nil {
		return ErrKeyAlreadyExists
	} else if err != buntdb.ErrNotFound {
		return err
	}
	s, err := json.Marshal(record.Data())
	if err != nil {
		return err
	}
	_, _, err = t.tx.Set(k, string(s), nil)
	return err
}
