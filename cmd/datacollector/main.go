package main

import (
	"context"
	"github.com/ethcero/connected-pv/internal/datacollector/app"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	quit := CloseHandler()
	StartApp(quit)
	for {
		select {
		case <-ctx.Done():
			return
		case <-quit:
			return
		}
	}
}

func CloseHandler() chan os.Signal {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)

	return quit
}

func StartApp(quit chan os.Signal) {
	go func() {
		application := app.NewApp()
		application.Start()
		quit <- os.Interrupt
	}()

	<-quit
}
