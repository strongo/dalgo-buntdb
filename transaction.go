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
	return dtb.db.Update(func(tx *buntdb.Tx) (err error) {
		return f(ctx, transaction{tx: tx, options: dal.NewTransactionOptions(options...)})
		// NOTE: managed tx rollback not allowed by buntdb
		//if err != nil {
		//	if rollbackErr := tx.Rollback(); rollbackErr != nil {
		//		return fmt.Errorf("failed to rollbacktransaction: %v: original error: %w", rollbackErr, err)
		//	}
		//}
		//return tx.Commit()
	})
}

var _ dal.ReadwriteTransaction = (*transaction)(nil)

type transaction struct {
	tx      *buntdb.Tx
	options dal.TransactionOptions
}

func (t transaction) QueryReader(c context.Context, query dal.Query) (dal.Reader, error) {
	return getReader(t.tx, query)
}

func (t transaction) QueryAllRecords(ctx context.Context, query dal.Query) (records []dal.Record, err error) {
	var reader buntdbReader
	if reader, err = getReader(t.tx, query); err != nil {
		return
	}
	limit := query.Limit()
	return dal.SelectAllRecords(reader, limit)
}

func (t transaction) ID() string {
	return ""
}

func (t transaction) Options() dal.TransactionOptions {
	return t.options
}

func (t transaction) Upsert(ctx context.Context, record dal.Record) error {
	return t.Set(ctx, record)
}
