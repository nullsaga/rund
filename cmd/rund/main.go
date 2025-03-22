package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/nullsaga/rund/internal/api"
	"github.com/nullsaga/rund/internal/cli"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	version = "0.0.1"
)

func main() {
	options := cli.NewWithDefaultOptions()
	options.Parse()

	if options.Version {
		fmt.Println(version)
		os.Exit(0)
	}

	if options.Help {
		flag.Usage()
		os.Exit(0)
	}

	logLevel := slog.LevelError
	if options.Verbose {
		logLevel = slog.LevelInfo
	}

	handler := slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		Level: logLevel,
	})

	slog.SetDefault(slog.New(handler))

	server := api.NewServer(options.Addr)
	server.RegisterHandlers()

	go func() {
		if err := server.Start(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("failed to start api server", "error", err)
		}
	}()

	shutdown, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-shutdown.Done()
	slog.Info("shutdown signal received")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Stop(ctx); err != nil {
		slog.Error("api server shutdown failed", "error", err)
	}

	slog.Info("api server shutdown completed successfully")
}
