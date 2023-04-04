package dalgo2buntdb

import (
	"fmt"
	"github.com/dal-go/dalgo/dal"
)

var errNotSupportedYet = fmt.Errorf("%w yet", dal.ErrNotSupported)
