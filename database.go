package dalgo2buntdb

import (
	"context"
	"github.com/dal-go/dalgo/dal"
	"github.com/tidwall/buntdb"
)

type database struct {
	db *buntdb.DB
	dal.QueryExecutor
}

func (dtb database) ID() string {
	//TODO implement me
	panic("implement me")
}

func (dtb database) Adapter() dal.Adapter {
	return dal.NewAdapter("buntdb", "v1.2.10")
}

var _ dal.DB = (*database)(nil)

// NewDatabase creates a new instance of DALgo adapter for BungDB
func NewDatabase(db *buntdb.DB) dal.DB {
	if db == nil {
		panic("db is a required parameter, got nil")
	}
	var database database
	var getReader = func(ctx context.Context, query dal.Query) (dal.Reader, error) {
		tx, err := db.Begin(false)
		if err != nil {
			return nil, err
		}
		return getReader(tx, query)
	}
	database.db = db
	database.QueryExecutor = dal.NewQueryExecutor(getReader)
	return database
}

func (dtb database) Upsert(ctx context.Context, record dal.Record) error {
	return dal.ErrNotImplementedYet
}

func getReader(buntdbTx *buntdb.Tx, query dal.Query) (buntdbReader, error) {
	return buntdbReader{query: query, tx: buntdbTx}, nil
}
