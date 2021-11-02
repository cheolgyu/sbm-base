package logging_test

import (
	"testing"

	"github.com/cheolgyu/stock-write-common/logging"
)

func Test(t *testing.T) {
	defer logging.ElapsedTime()()
	logging.Log.Debug("aaaaaaa")
}
