package main

import (
	"bjss.com/ashley.winter/to_do/part2_todo_app/cmd/cliapp"
	"bjss.com/ashley.winter/to_do/part2_todo_app/repo"
	"bjss.com/ashley.winter/to_do/part2_todo_app/web/api"
	"bjss.com/ashley.winter/to_do/part2_todo_app/web/ssr"
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	ctx := context.WithValue(context.Background(), "logger", logger)

	sharedStore := repo.InitRepo(ctx)

	go cliapp.RunCli(ctx, sharedStore)
	go ssr.ListenAndServe(ctx, sharedStore)
	go api.ListenAndServe(ctx, sharedStore)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	doneChan := make(chan bool, 1)

	go func() {
		<-signalChan
		slog.InfoContext(ctx, "\nReceived an interrupt, performing cleanup...")
		// Perform any cleanup here
		doneChan <- true
	}()

	fmt.Println("Press Ctrl+C to exit")
	<-doneChan
	slog.InfoContext(ctx, "Cleanup complete, exiting.")
}
