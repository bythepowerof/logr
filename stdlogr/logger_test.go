package stdlogr

import (
	"errors"
	"time"
)

func ExampleInfoLogger() {
	mock, _ := time.Parse("2006-01-02", "2015-12-15")

	logger := New("foo", true, 0)
	logger.(*StdLogr).clock = &clock{mock: mock}
	logger.Info("test log", "hello", "world")
	// Output: level=info ts=1450137600 name=foo msg="test log" hello=world
}

func ExampleErrLogger() {
	mock, _ := time.Parse("2006-01-02", "2015-12-15")
	err := errors.New("BOOM SUCKA!")

	logger := New("bar", true, 0)
	logger.(*StdLogr).clock = &clock{mock: mock}
	logger.Error(err, "test error log", "hello", "world")
	// Output: level=error ts=1450137600 name=bar msg="test error log" hello=world error="BOOM SUCKA!"
}

func ExampleNonVerboseLogger() {
	mock, _ := time.Parse("2006-01-02", "2015-12-15")

	logger := New("sucka", true, 0)
	logger.(*StdLogr).clock = &clock{mock: mock}
	logger.V(1).Info("test verbose log", "hello", "crazy world")
	// Output:
}

func ExampleVerboseLogger() {
	mock, _ := time.Parse("2006-01-02", "2015-12-15")

	logger := New("sucka", true, 1)
	vLogger := logger.V(1)
	vLogger.(*StdLogr).clock = &clock{mock: mock}
	vLogger.Info("test verbose log", "hello", "crazy world")
	// Output: level=info ts=1450137600 name=sucka msg="test verbose log" v=1 hello="crazy world"
}

func ExampleNamedLogger() {
	mock, _ := time.Parse("2006-01-02", "2015-12-15")

	logger := New("foo", true, 0)
	namedLogger := logger.WithName("bar")
	namedLogger.(*StdLogr).clock = &clock{mock: mock}
	namedLogger.Info("test log", "hello", "world")
	// Output: level=info ts=1450137600 name=foo.bar msg="test log" hello=world
}

func ExampleValuesLogger() {
	mock, _ := time.Parse("2006-01-02", "2015-12-15")

	logger := New("foo", true, 0)
	valuesLogger := logger.WithValues("goodbye", "crazy world")
	valuesLogger.(*StdLogr).clock = &clock{mock: mock}
	valuesLogger.Info("test log", "hello", "world")
	// Output: level=info ts=1450137600 name=foo msg="test log" goodbye="crazy world" hello=world
}
