package params

import (
	"sync"

	"github.com/dop251/goja"
)

var vmPool = sync.Pool{
	New: func() interface{} {
		return goja.New()
	},
}
