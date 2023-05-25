package runtimex

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// ServeWithSignalContext is a helper function that runs a server until a signal is received.
func ServeWithSignalContext(srv func(ctx context.Context)) {
	// shutdownSignals is the list of signals that can trigger a graceful shutdown.
	shutdownSignals := []os.Signal{os.Interrupt, syscall.SIGTERM, os.Kill}
	shutdownChannel := make(chan os.Signal, 2)

	signal.Notify(shutdownChannel, shutdownSignals...)
	ctx, cancel := context.WithCancel(context.Background())
	GoRoutine(func() {
		defer cancel()
		srv(ctx)
	})
	sig := <-shutdownChannel
	log.Printf("receive first signal. sig: %v\r\n", sig)
	cancel()
	sig = <-shutdownChannel
	log.Fatalf("receive second signal. sig: %v\r\n", sig)
}
