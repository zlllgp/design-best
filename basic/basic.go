package basic

import (
	"design-best/money"
	"fmt"
	"strings"
	"time"
)

type Writer interface {
	Write(p []byte) (n int, err error)
}
type Reader interface {
	Read(p []byte) (n int, err error)
}
type ReadWriter interface {
	Reader
	Writer
}

type Card struct {
	// omitted
}

type Service struct {
	// omitted
}
type Stub struct{}

func (s *Service) Charge(c *Card, amount money.Money) error { /* omitted */
	return nil
}

func (Stub) Charge(*Card, money.Money) error { return nil }

func Abs(i int) int {
	if i < 0 {
		i *= -1
	}
	return i
}

type Person struct {
	Name string
	Age  int
}

// 值接收者
func (p Person) SayHello() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

// 指针接收者
func (p *Person) GrowOlder(years int) {
	p.Age += years
}

const (
	n1 = iota
	n2 = 100
	n3 = iota
	n4
)
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

// 全局变量m
var m = 100

func main1() {
	s := "this is a "
	strings.Contains(s, "a")
	n := 10
	m := 200 // 此处声明局部变量m
	fmt.Println(m, n)

	// 创建一个 Person 实例
	p := &Person{Name: "Alice", Age: 30}

	// 调用值接收者方法
	p.SayHello() // 输出 Hello, my name is Alice and I am 30 years old.

	// 调用指针接收者方法
	p.GrowOlder(5)

	// 再次调用值接收者方法
	p.SayHello() // 输出 Hello, my name is Alice and I am 35 years old.

	var x int = 10
	var a *int = &x
	*a = 20
	fmt.Println(*a)

	// 创建一个双向通道
	ch := make(chan bool)

	// 启动一个 Goroutine 来发送数据
	go Sender(ch)

	// 启动一个 Goroutine 来接收数据
	go Receiver(ch)

	time.Sleep(1 * time.Second)

	traversalString()

	changeString()
}

// 遍历字符串
func traversalString() {
	s := "pprof.cn博客"
	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()
	for _, r := range s { //rune
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()
}

func changeString() {
	s1 := "hello"
	// 强制类型转换
	byteS1 := []byte(s1)
	byteS1[0] = 'H'
	fmt.Println(string(byteS1))

	s2 := "博客"
	runeS2 := []rune(s2)
	runeS2[0] = '狗'
	fmt.Println(string(runeS2))

	var a [5]string
	fmt.Println(a)

	var arr0 [5]int = [5]int{1, 2, 3}
	for _, v := range arr0 {
		fmt.Println(v)
	}
}

// Sender 函数接收一个只能发送 bool 值的通道
func Sender(ch chan<- bool) {
	ch <- true // 发送一个 bool 值
}

// Receiver 函数接收一个只能接收 bool 值的通道
func Receiver(ch <-chan bool) {
	val := <-ch // 接收一个 bool 值
	fmt.Println("Received:", val)
}

func sum(values <-chan int) int {
	// ...
	return 0
}
