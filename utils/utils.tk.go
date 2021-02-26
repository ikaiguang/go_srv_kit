package tkuitls

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
)

const (
	_stackTracerDepth = 7
)

// stackTracer errors.StackTrace
type stackTracer interface {
	StackTrace() errors.StackTrace
}

// InfoTracer .
func InfoTracer(err error, stackTracerDepths ...int) {
	if err == nil {
		return
	}

	trace, ok := err.(stackTracer)
	if !ok {
		fmt.Printf("%+v\n", err)
		return
	}
	fmt.Println(err)
	//fmt.Printf("%+v\n",err)

	st := trace.StackTrace()
	depth := _stackTracerDepth
	if len(stackTracerDepths) > 0 {
		depth = stackTracerDepths[0]
	}
	if len(st) > depth {
		fmt.Printf("%+v\n", st[0:depth]) // top n frames
		return
	}
	fmt.Printf("%+v\n", st) // top n frames
}

// FatalTracer .
func FatalTracer(err error, stackTracerDepths ...int) {
	InfoTracer(err, stackTracerDepths...)
	os.Exit(1)
}

// ExitTracer exit with code 0
func ExitTracer(err error, stackTracerDepths ...int) {
	InfoTracer(err, stackTracerDepths...)
	os.Exit(0)
}
