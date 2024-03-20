package runtimex

import (
	"context"
)

func GoRoutine(method func()) {
	if method == nil {
		return
	}
	Go(context.Background(), func(ctx context.Context) {
		method()
	})

}

func Go(ctx context.Context, method func(context.Context), additionalHandlers ...func(context.Context, any)) {
	if method == nil {
		return
	}
	go func() {
		defer HandleCrash(ctx, additionalHandlers...)
		method(ctx)
	}()

}
