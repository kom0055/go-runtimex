package runtimex

import "context"

func WithCancel(parent context.Context, ctrlCtxs ...context.Context) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(parent)
	for i := range ctrlCtxs {
		ctrlCtx := ctrlCtxs[i]
		GoRoutine(func() {
			defer cancel()
			select {
			case <-ctx.Done():
				return
			case <-ctrlCtx.Done():
				return
			}
		})

	}

	return ctx, cancel
}
