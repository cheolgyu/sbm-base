package logging_test

import (
	"testing"

	"github.com/cheolgyu/base/logging"
)

func Test(t *testing.T) {
	defer logging.ElapsedTime()()
	logging.Log.Debug("aaaaaaa")
}
