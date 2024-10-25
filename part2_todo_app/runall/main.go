package main

import (
	"bjss.com/ashley.winter/to_do/part2_todo_app/cmd/cliapp"
	"bjss.com/ashley.winter/to_do/part2_todo_app/repo"
	"bjss.com/ashley.winter/to_do/part2_todo_app/web/api"
	"bjss.com/ashley.winter/to_do/part2_todo_app/web/ssr"
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

type CustomHandler struct {
	slog.Handler
}

func (h *CustomHandler) Handle(ctx context.Context, r slog.Record) error {
	if value := ctx.Value("requestId"); value != nil {
		r.AddAttrs(slog.String("requestId", ctx.Value("requestId").(string)))
	}
	if value := ctx.Value("server"); value != nil {
		r.AddAttrs(slog.String("server", value.(string)))
	}
	return h.Handler.Handle(ctx, r)
}

func main() {
	baseLogger := slog.NewJSONHandler(os.Stdout, nil)

	customLoggerHandler := &CustomHandler{Handler: baseLogger}

	newLogger := slog.New(customLoggerHandler)

	ctx := context.WithValue(context.Background(), "logger", newLogger)

	sharedStore := repo.InitRepo(ctx)

	if ctx.Value("logger") == nil {
		log.Fatal("No logger in context")
	}

	go cliapp.RunCli(ctx, sharedStore)
	go ssr.ListenAndServe(ctx, sharedStore)
	go api.ListenAndServe(ctx, sharedStore)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	doneChan := make(chan bool, 1)

	go func() {
		<-signalChan
		newLogger.InfoContext(ctx, "\nReceived an interrupt")
		// Perform any cleanup here
		doneChan <- true
	}()

	fmt.Println("Press Ctrl+C to exit")
	<-doneChan
	newLogger.InfoContext(ctx, "Exiting.")
}
