package main

import (
	"fmt"
	"unsafe"
)

type Country struct {
	Id         int16
	population int64
	IsAdvanced bool
}

func main() {
	var country = Country{}
	fmt.Println("对齐边界:", unsafe.Alignof(country))
	fmt.Println("占用空间:", unsafe.Sizeof(country))
}
