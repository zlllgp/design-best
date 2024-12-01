package best

import (
	"fmt"
	"unsafe"
)

func Str2Bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func Bytes2Str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func String() {
	clusterName := "CN"
	groupName := "sh"

	//Good
	hostName1 := clusterName + "@" + groupName
	//Bad
	hostName2 := fmt.Sprintf("%s@%s", clusterName, groupName)

	fmt.Println(hostName1)
	fmt.Println(hostName2)
}
