package handling

import (
	"fmt"
	"github.com/pkg/errors"
)

//方案一：简单错误，哨兵模式，不推荐

var (
	SUCCESS   = errors.New("success")
	FAIL      = errors.New("fail")
	EXCEPTION = errors.New("exception")
)

func CallBack() error {
	return FAIL
}

// usage
func usage() {
	err := CallBack()
	if errors.Is(err, FAIL) {
		fmt.Println("FAIL")
	}
}
