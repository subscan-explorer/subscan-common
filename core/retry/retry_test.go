package retry

import (
	"testing"
	"time"

	"github.com/pkg/errors"
	"github.com/subscan-explorer/subscan-common/core/backoff"
	"github.com/subscan-explorer/subscan-common/core/util/xtime"
)

func TestRetrier_Do(t *testing.T) {
	bo := backoff.NewConstantBackoff(xtime.Duration(100 * time.Millisecond))
	err := NewRetrier(bo).Do(HelloDo, 5)
	t.Log(err)
}

func HelloDo() (err error) {
	err = errors.New("retry testing")
	return
}
