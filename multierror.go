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
	length := len(m.errs)
	m.mutex.Unlock()

	if length == 0 {
		return ""
	}

	formattedError := make([]string, length)
	for i := 0; i < length; i++ {
		formattedError[i] = m.errs[i].Error()
	}

	return strings.Join(formattedError, ", ")
}
