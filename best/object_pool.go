package best

import (
	"sync"
)

type Result struct {
	Code    int
	Msg     string
	Version string
	Data    interface{}
}

var pool = &sync.Pool{
	New: func() interface{} {
		return &Result{}
	},
}

func GetPoolObject(code int, msg string) *Result {
	result := pool.Get().(*Result)
	result.Code = code
	result.Msg = msg
	result.Version = "1.0"
	return result
}

func PutPoolObject(r *Result) {
	r.Msg = ""
	r.Code = 0
	pool.Put(r)
}
