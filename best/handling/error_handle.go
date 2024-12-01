package handling

import (
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
)

func normal() error {
	_, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return err
	}
	return nil
}

func warp() error {
	_, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return errors.Wrap(err, "read from stdin err")
	}
	return nil
}
