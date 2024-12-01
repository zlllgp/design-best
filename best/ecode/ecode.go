package ecode

//
//import (
//	"fmt"
//)
//
//// Codes see https://github.com/zlllgp/ecode
//type Codes interface {
//	Code() int
//	Error() string
//	Message() string
//	Details() []interface{}
//}
//
//type Code int
//
//// Error 覆盖ecode message
//func Error(code Code, message string) *api.Status {
//	return &api.Status{}
//}
//
//// Errorf 覆盖 ecode message，使用自定义format
//func Errorf(code Code, format string, args ...interface{}) *api.Status {
//	return Error(code, fmt.Sprintf(format, args...))
//}
//
//func (c *Code) Error() string {
//	return "" //todo
//}
