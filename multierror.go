package valkyrie

import (
	"errors"
	"strings"
	"sync"
)

// MultiError implements error interface.
// An instance of MultiError has zero or more errors.
type MultiError struct {
	mutex sync.Mutex
	errs  []error
}

// Push adds an error to MultiError.
func (m *MultiError) Push(errString string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.errs = append(m.errs, errors.New(errString))
}

// HasError checks if MultiError has any error.
func (m *MultiError) HasError() error {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	if len(m.errs) == 0 {
		return nil
	}

	return m
}

// Error implements error interface.
func (m *MultiError) Error() string {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	formattedError := make([]string, len(m.errs))
	for i, e := range m.errs {
		formattedError[i] = e.Error()
	}

	return strings.Join(formattedError, ", ")
}

// Cause returns the first original error.
func (m *MultiError) Cause() error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if len(m.errs) == 0 {
		return nil
	}

	return m.errs[0]
}

// Unwrap provides compatibility for go 1.13 error chains.
func (m *MultiError) Unwrap() error {
	return m.Cause()
}
