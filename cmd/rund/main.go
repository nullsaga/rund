package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/nullsaga/rund/internal/api"
	"github.com/nullsaga/rund/internal/cli"
	"github.com/nullsaga/rund/internal/conf"
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

	rootConf, err := conf.NewLoader().LoadConf(options.ConfPath)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	server := api.NewServer(options.Addr)
	server.RegisterHandlers(rootConf)

	go func() {
		if err = server.Start(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error(err.Error())
		}
	}()

	shutdown, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-shutdown.Done()
	slog.Info("shutdown signal received")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err = server.Stop(ctx); err != nil {
		slog.Error(err.Error())
	}

	slog.Info("api server shutdown completed successfully")
}
