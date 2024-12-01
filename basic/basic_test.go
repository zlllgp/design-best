package basic

import (
	"design-best/money"
	"fmt"
	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"testing"
)

func newTestRedis() *redis.Client {
	mr, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})
	return client
}

var cli = newTestRedis()

func TestRedis(t *testing.T) {
	type args struct {
		cli   *redis.Client
		count int64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "set mock", args: args{cli: cli, count: 1000}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

		})
	}
}

func TestAssert(t *testing.T) {
	assert.Equal(t, 123, 123, "should be equal")
}

func TestMonkey(t *testing.T) {
	//monkey.Patch(time.Now, func() time.Time { return time.Now() })
}

func TestService_Charge(t *testing.T) {
	type args struct {
		c      *Card
		amount money.Money
	}
	tests := []struct {
		name    string
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{}
			tt.wantErr(t, s.Charge(tt.args.c, tt.args.amount), fmt.Sprintf("Charge(%v, %v)", tt.args.c, tt.args.amount))
		})
	}
}
