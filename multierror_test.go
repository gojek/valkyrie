package valkyrie

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strconv"
	"sync"
)

func TestMultiErrorValidation(t *testing.T) {
	me := &MultiError{}
	me.Push("one")
	me.Push("two")

	res := me.HasError()

	assert.Error(t, res, "expected populated multi-error validation to be non-nil")
}

func TestMultiErrorRespresentation(t *testing.T) {
	me := &MultiError{}
	me.Push("one")
	me.Push("two")

	require.Error(t, me, "no multierror occured")

	res := me.Error()

	assert.Equal(t, "one, two", res)
}

func TestMultiErrorWithoutErrorsValidationIsNil(t *testing.T) {
	me := &MultiError{}

	err := me.HasError()

	assert.NoError(t, err, "expected empty multi-error validation to be nil")
}

func TestMultiErrorRespresentationIsEmpty(t *testing.T) {
	me := &MultiError{}

	res := me.Error()

	assert.Empty(t, res, "expected empty multi-error representation to be empty")
}

func TestMultiErrorConcurrentPushAndError(t *testing.T) {
	me := &MultiError{}

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			me.Push(strconv.Itoa(i))
			wg.Done()
		}(i)
	}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			_ = me.Error()
			wg.Done()
		}()
	}
	wg.Wait()
}
