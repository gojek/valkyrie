package valkyrie

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMultiErrorValidation(t *testing.T) {
	me := &MultiError{mutex: new(sync.Mutex)}
	me.Push("one")
	me.Push("two")

	res := me.HasError()

	assert.Error(t, res, "expected populated multi-error validation to be non-nil")
}

func TestMultiErrorRespresentation(t *testing.T) {
	me := &MultiError{mutex: new(sync.Mutex)}
	me.Push("one")
	me.Push("two")

	require.Error(t, me, "no multierror occured")

	res := me.Error()

	assert.Equal(t, "one, two", res)
}

func TestMultiErrorWithoutErrorsValidationIsNil(t *testing.T) {
	me := &MultiError{mutex: new(sync.Mutex)}

	err := me.HasError()

	assert.NoError(t, err, "expected empty multi-error validation to be nil")
}

func TestMultiErrorRespresentationIsEmpty(t *testing.T) {
	me := &MultiError{mutex: new(sync.Mutex)}

	res := me.Error()

	assert.Empty(t, res, "expected empty multi-error representation to be empty")
}
