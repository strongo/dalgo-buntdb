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
			if errors.Is(err, ErrKeyAlreadyExists) {
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
	} else if !errors.Is(err, buntdb.ErrNotFound) {
		return err
	}
	record.SetError(nil)
	data := record.Data()
	s, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, _, err = t.tx.Set(k, string(s), nil)
	return err
}

// InsertMulti inserts multiple records in a single transaction.
func (t transaction) InsertMulti(ctx context.Context, records []dal.Record, opts ...dal.InsertOption) error {
	for _, record := range records {
		if err := t.Insert(ctx, record, opts...); err != nil {
			return err
		}
	}
	return nil
}
