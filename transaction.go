package dalgo2buntdb

import (
	"context"
	"github.com/dal-go/dalgo/dal"
	"github.com/tidwall/buntdb"
)

func (dtb database) RunReadonlyTransaction(ctx context.Context, f dal.ROTxWorker, options ...dal.TransactionOption) error {
	return dtb.db.View(func(tx *buntdb.Tx) error {
		return f(ctx, transaction{tx: tx, options: dal.NewTransactionOptions(options...)})
	})
}

func (dtb database) RunReadwriteTransaction(ctx context.Context, f dal.RWTxWorker, options ...dal.TransactionOption) error {
	return dtb.db.Update(func(tx *buntdb.Tx) error {
		return f(ctx, transaction{tx: tx, options: dal.NewTransactionOptions(options...)})
	})
}

var _ dal.ReadwriteTransaction = (*transaction)(nil)

type transaction struct {
	tx      *buntdb.Tx
	options dal.TransactionOptions
}

func (t transaction) Options() dal.TransactionOptions {
	return t.options
}

func (t transaction) Upsert(ctx context.Context, record dal.Record) error {
	return t.Set(ctx, record)
}

func (t transaction) Select(context.Context, dal.Select) (dal.Reader, error) {
	return nil, errNotSupportedYet // TODO(help-wanted): needs implementation
}
