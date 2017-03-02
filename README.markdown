[![Build Status](https://travis-ci.org/gojek-engineering/go-multierror.svg?branch=master)](https://travis-ci.org/gojek-engineering/go-multierror)

# multierror
--
    import "go-multierror"


## Usage

#### type MultiError

```go
type MultiError struct {
}
```

MultiError implements error interface. An instance of MultiError has zero or
more errors.

#### func  NewMultiError

```go
func NewMultiError() *MultiError
```
NewMultiError: returns a thread safe instance of multierror

#### func (*MultiError) Error

```go
func (m *MultiError) Error() string
```
Error implements error interface.

#### func (*MultiError) HasError

```go
func (m *MultiError) HasError() error
```
HasError checks if MultiError has any error.

#### func (*MultiError) Push

```go
func (m *MultiError) Push(errString string)
```
Push adds an error to MultiError.
