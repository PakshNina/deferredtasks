package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"deferredtasks/internal/cron"
	"deferredtasks/internal/ticker"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	ts := ticker.NewTickerService()
	go ts.Run(ctx)

	cs := cron.NewCronService()
	go cs.Run(ctx)

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)
	<-shutdown
	cancel()
}
