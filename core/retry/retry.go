package retry

import (
	"time"

	"github.com/subscan-explorer/subscan-common/core/backoff"
)

// Retriable defines contract for retriers to implement
type Retriable interface {
	NextInterval(retry int) time.Duration
	Do(fn RetryFunc, retries int) (err error)
}

type RetryFunc func() (err error)

// RetriableFunc is an adapter to allow the use of ordinary functions
// as a Retriable
type RetriableFunc func(retry int) time.Duration

func (f RetriableFunc) Do(fn RetryFunc, retries int) (err error) {
	return fn()
}

//NextInterval calls f(retry)
func (f RetriableFunc) NextInterval(retry int) time.Duration {
	return f(retry)
}

func (f RetryFunc) Do() (err error) {
	return f()
}

type retrier struct {
	backoff backoff.Backoff
}

// NewRetrier returns retrier with some backoff strategy
func NewRetrier(backoff backoff.Backoff) Retriable {
	return &retrier{
		backoff: backoff,
	}
}

// NewRetrierFunc returns a retrier with a retry function defined
func NewRetrierFunc(f RetriableFunc) Retriable {
	return f
}

// NextInterval returns next retriable time
func (r *retrier) NextInterval(retry int) time.Duration {
	return r.backoff.Next(retry)
}

type noRetrier struct {
}

// NewNoRetrier returns a null object for retriable
func NewNoRetrier() Retriable {
	return &noRetrier{}
}

// NextInterval returns next retriable time, always 0
func (r *noRetrier) NextInterval(retry int) time.Duration {
	return 0 * time.Millisecond
}

func (r *retrier) Do(fn RetryFunc, retries int) (err error) {
	for i := 0; i <= retries; i++ {
		if err = fn(); err != nil {
			time.Sleep(r.NextInterval(i))
			continue
		}
		break
	}
	return nil
}

func (r *noRetrier) Do(fn RetryFunc, retries int) (err error) {
	for i := 0; i <= retries; i++ {
		if err = fn(); err != nil {
			time.Sleep(r.NextInterval(i))
			continue
		}
		break
	}
	return nil
}
