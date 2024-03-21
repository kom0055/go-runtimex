package runtimex

import (
	"context"
	"log"

	"github.com/pkg/errors"
)

var (
	logger = func(ctx context.Context, format string, v ...any) {
		log.Printf(format, v...)
	}
	// PanicHandlers is a list of functions which will be invoked when a panic happens.
	PanicHandlers = []func(context.Context, error){logPanic}
)

// HandleCrash simply catches a crash and logs an error. Meant to be called via
// defer.  Additional context-specific handlers can be provided, and will be
// called in case of panic.  HandleCrash actually crashes, after calling the
// handlers and logging the panic message.
//
// E.g., you can provide one or more additional handlers for something like shutting down go routines gracefully.
func HandleCrash(ctx context.Context, additionalHandlers ...func(context.Context, error)) {
	if r := recover(); r != nil {
		err, ok := r.(error)
		if ok {
			err = errors.WithStack(err)
		} else {
			err = errors.Errorf("%v", r)
		}
		for _, fn := range PanicHandlers {
			fn(ctx, err)
		}
		for _, fn := range additionalHandlers {
			fn(ctx, err)
		}

	}
}

// logPanic logs the caller tree when a panic occurs (except in the special case of http.ErrAbortHandler).
func logPanic(ctx context.Context, err error) {
	//if r == http.ErrAbortHandler {
	//	// honor the http.ErrAbortHandler sentinel panic value:
	//	//   ErrAbortHandler is a sentinel panic value to abort a handler.
	//	//   While any panic from ServeHTTP aborts the response to the client,
	//	//   panicking with ErrAbortHandler also suppresses logging of a stack trace to the server's error log.
	//	return
	//}

	if l := logger; l != nil {
		l(ctx, "panic happened, cause: %+v", err)
	}
}

func ReallyCrash(ctx context.Context, r any) {
	panic(r)
}

func SetLogger(l func(ctx context.Context, format string, v ...any)) {
	logger = l
}
