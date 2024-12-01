package best

import (
	"fmt"
	"strconv"
	"testing"
)

func TestString(t *testing.T) {
	String()
}

func TestSlice(t *testing.T) {
	Slice()
}

func BenchmarkStrcovn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strconv.Itoa(i)
	}
}
func BenchmarkFmtSprint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprint(i)
	}
}

func BenchmarkPlus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := "a" + "b"
		_ = s
	}
}
func BenchmarkFmtSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := fmt.Sprintf("%s%s", "a", "b")
		_ = s
	}
}

func BenchmarkSliceWithCap(b *testing.B) {
	//Good
	list := make([]int, 0, 100)
	for i := 0; i < b.N; i++ {
		for n := 0; n < 10000; n++ {
			list = append(list, n)
		}
	}
}

func BenchmarkSliceNoCap(b *testing.B) {
	//Bad
	list := make([]int, 0)
	for i := 0; i < b.N; i++ {
		for n := 0; n < 10000; n++ {
			list = append(list, n)
		}
	}
}

func BenchmarkFor(b *testing.B) {
	//Good
	list := getSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var tmp int
		for k := 0; k < len(list); k++ {
			tmp = list[k]
		}
		_ = tmp
	}
}
func BenchmarkRange(b *testing.B) {
	//Good
	list := getSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var tmp int
		for _, num := range list {
			tmp = num
		}
		_ = tmp
	}
}

func getSlice() []int {
	list := make([]int, 0, 10000)
	for n := 0; n < 10000; n++ {
		list = append(list, n)
	}
	return list
}

func BenchmarkForStruct(b *testing.B) {
	//Good
	list := getStruct()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var tmp int
		for k := 0; k < len(list); k++ {
			tmp = list[k].Id
		}
		_ = tmp
	}
}
func BenchmarkRangeStruct(b *testing.B) {
	//Good
	list := getStruct()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var tmp int
		for _, num := range list {
			tmp = num.Id
		}
		_ = tmp
	}
}

func getStruct() []Item {
	list := make([]Item, 0, 10000)
	for n := 0; n < 10000; n++ {
		list = append(list, Item{Id: n})
	}
	return list
}

type Item struct {
	Id int
}

func TestChannel(t *testing.T) {
	Channel()
}
