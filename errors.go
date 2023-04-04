package dalgo2buntdb

import (
	"fmt"
	"github.com/strongo/dalgo/dal"
)

var errNotSupportedYet = fmt.Errorf("%w yet", dal.ErrNotSupported)
