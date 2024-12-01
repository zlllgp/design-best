package handling

import (
	"fmt"
	"github.com/cloudwego/kitex/tool/internal_pkg/log"
	"github.com/pkg/errors"
)

var (
	ErrBookNotFound = errors.New("book not found")
)

// 对象包含，不推荐
type Book struct {
	Name  string
	Price int
	Store int
	error error //每个struct增加一个error
}

type BookBuy interface {
	CalcDiscount(count int) (*Book, error)
}

func (b *Book) CalcDiscount(count int) (*Book, error) {
	if b.error != nil {
		log.Info("Book calc discount handling, %d ", count)
		return b, nil
	}

	//业务逻辑...
	b.error = someError()
	return b, nil
}

func (b *Book) CalcDiscount2(count int) (*Book, error) {
	if b.error != nil {
		return b, fmt.Errorf("CalcDiscount2 - > %w", b.error)
	}

	//业务逻辑...
	b.error = someError()
	return b, nil
}

func (b *Book) CalcDiscount3(count int) (*Book, error) {
	if b.error != nil {
		return b, fmt.Errorf("CalcDiscount3 - > %w , %w", b.error, ErrBookNotFound)
	}

	//业务逻辑...
	b.error = someError()
	return b, nil
}

func (b *Book) CalcDiscount4(count int, m MyFun) (*Book, error) {
	if b.error != nil {
		log.Infof("Book calc discount handling, %d ", count)
		return b, errors.Wrapf(b.error, "CalcDiscount handling , name is %s", b.Name)
	}

	//业务逻辑...
	b.error = someError()
	return b, nil
}

func Controller() error {
	b := &Book{}
	_, err := b.CalcDiscount3(2)
	if err != nil {
		err = errors.WithMessagef(err, "Controller calc discount , count : %d", 2)
		log.Infof("Controller calc discount handling, %d ", 2)
		return GrpcError(err)
	}
	return nil
}

func GrpcError(err error) error {
	return err
}

func someError() error {
	return errors.New("some handling")
}

type MyFun func(b Book) error

func NewEnterFun(fun MyFun) MyFun {
	return func(b Book) error {
		return fun(b)
	}
}

func main() {
	b := &Book{}
	_, err := b.CalcDiscount3(2)
	if err != nil {
		fmt.Errorf("calc %v", err)
	}

	if err != nil {
		if errors.Is(err, ErrBookNotFound) {
			//todo
		}
	}

}
