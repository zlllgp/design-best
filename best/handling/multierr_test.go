package handling

import (
	"github.com/pkg/errors"
	"go.uber.org/multierr"
	"testing"
)

func TestMultierr(t *testing.T) {
	var err error
	err = multierr.Append(err, errors.New("call 1 failed"))
	err = multierr.Append(err, errors.New("call 2 failed"))

	for _, e := range multierr.Errors(err) {
		t.Log(e.Error())
	}
}
