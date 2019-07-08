package errors

import (
	"fmt"
)

type LogrError interface {
	error
	WithKVs(...interface{}) error
	KVs() []interface{}
}

type WrappedLogrError interface {
	LogrError
	Cause() error
}

type logrError struct {
	msg string
	kvs []interface{}
}

type wrappedLogrError struct {
	*logrError
	cause error
}

func New(msg string) *logrError {
	return &logrError{msg: msg}
}

func Newf(format string, args ...interface{}) *logrError {
	return &logrError{msg: fmt.Sprintf(format, args...)}
}

func (this logrError) Error() string {
	return this.msg
}

func (this *logrError) WithKVs(kvs ...interface{}) error {
	this.kvs = kvs
	return this
}

func (this logrError) KVs() []interface{} {
	return this.kvs
}

func Wrap(err error, msg string) *wrappedLogrError {
	return &wrappedLogrError{
		logrError: &logrError{msg: msg},
		cause:     err,
	}
}

func Wrapf(err error, format string, args ...interface{}) *wrappedLogrError {
	return &wrappedLogrError{
		logrError: &logrError{msg: fmt.Sprintf(format, args...)},
		cause:     err,
	}
}

func (this wrappedLogrError) Error() string {
	return this.msg + ": " + this.cause.Error()
}

func (this *wrappedLogrError) WithKVs(kvs ...interface{}) error {
	this.kvs = kvs
	return this
}

func (this wrappedLogrError) Cause() error {
	return this.cause
}
