package stdlogr

import (
	"fmt"
	"time"

	"github.com/activeshadow/logr/util"
	"github.com/go-logr/logr"
)

type clock struct {
	mock time.Time // set in tests
}

func (this *clock) now() time.Time {
	if this == nil { // normal operations
		return time.Now()
	}

	return this.mock // testing only
}

type StdInfoLogr struct {
	enabled bool
	name    string
	kvs     map[string]interface{}

	// Normally nil, set by test code only.
	clock *clock
}

func (this StdInfoLogr) Info(msg string, kvs ...interface{}) {
	if !this.enabled {
		return
	}

	fmt.Println("level=info " + this.fmtMsg(msg, kvs...))
}

func (this StdInfoLogr) Enabled() bool {
	return this.enabled
}

func (this StdInfoLogr) fmtMsg(msg string, kvs ...interface{}) string {
	now := this.clock.now().UTC()

	msg = utils.QuoteSpaces(msg)
	msg = fmt.Sprintf("ts=%d name=%s msg=%s", now.Unix(), this.name, msg)

	for k, v := range this.kvs {
		str := utils.QuoteSpaces(v)
		msg = fmt.Sprintf("%s %s=%s", msg, k, str)
	}

	for i := 0; i < len(kvs); i += 2 {
		str := utils.QuoteSpaces(kvs[i+1])
		msg = fmt.Sprintf("%s %s=%s", msg, kvs[i], str)
	}

	return msg
}

type StdLogr struct {
	StdInfoLogr
	verbosity int
}

func (this StdLogr) Error(err error, msg string, kvs ...interface{}) {
	kvs = append(kvs, "error", err)
	fmt.Println("level=error " + this.fmtMsg(msg, kvs...))
}

func (this StdLogr) V(level int) logr.InfoLogger {
	if level > this.verbosity {
		return StdInfoLogr{enabled: false}
	}

	return this.WithValues("v", level)
}

func (this StdLogr) WithValues(kvs ...interface{}) logr.Logger {
	newKVs := make(map[string]interface{})

	for k, v := range this.kvs {
		newKVs[k] = v
	}

	for i := 0; i < len(kvs); i += 2 {
		newKVs[kvs[i].(string)] = kvs[i+1]
	}

	return &StdLogr{
		StdInfoLogr: StdInfoLogr{
			enabled: this.enabled,
			name:    this.name,
			kvs:     newKVs,
		},
		verbosity: this.verbosity,
	}
}

func (this StdLogr) WithName(name string) logr.Logger {
	name = fmt.Sprintf("%s.%s", this.name, name)

	return &StdLogr{
		StdInfoLogr: StdInfoLogr{
			enabled: this.enabled,
			name:    name,
			kvs:     this.kvs,
		},
		verbosity: this.verbosity,
	}
}

func New(name string, enabled bool, verbosity int) logr.Logger {
	return &StdLogr{
		StdInfoLogr: StdInfoLogr{
			enabled: enabled,
			name:    name,
			kvs:     make(map[string]interface{}),
		},
		verbosity: verbosity,
	}
}
