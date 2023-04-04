package dalgo2buntdb

import (
	"context"
	"github.com/dal-go/dalgo/dal"
)

func (database) Select(ctx context.Context, query dal.Select) (dal.Reader, error) {
	return nil, nil
}
