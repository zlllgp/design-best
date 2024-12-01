package handling

import "fmt"

//方案二：对象封装

// 全局错误类型
type DreamErrorCode int

// 全局错误码
const (
	ErrorBookNotFoundCode DreamErrorCode = iota + 1
	ErrorBookHasBeenBorrwedCode
)

var errCodeMap = map[DreamErrorCode]string{
	ErrorBookNotFoundCode:       "Book not found",
	ErrorBookHasBeenBorrwedCode: "Book has been borrowed",
}

var (
	ErrCodeBookNotFound     = NewDreamError(ErrorBookNotFoundCode)
	ErrorBookHasBeenBorrwed = NewDreamError(ErrorBookHasBeenBorrwedCode)
)

type DreamError struct {
	Code    DreamErrorCode
	Message string
}

func NewDreamError(code DreamErrorCode) *DreamError {
	return &DreamError{
		Code:    code,
		Message: errCodeMap[code],
	}
}

func (e *DreamError) Error() string {
	fmt.Println("error happen , ", e.Message)
	return e.Message
}
