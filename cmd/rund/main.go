package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/nullsaga/rund/internal/api"
	"github.com/nullsaga/rund/internal/cli"
	"github.com/nullsaga/rund/internal/conf"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	version = "0.0.2"
)

func main() {
	options := cli.NewWithDefaultOptions()
	options.Parse()

	if !options.Verbose {
		log.SetOutput(io.Discard)
	}

	projectsConf, err := conf.NewLoader().LoadConf(options.ConfPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	server := api.NewServer(fmt.Sprintf("%s:%d", options.Ip, options.Port))
	server.RegisterHandlers(projectsConf)

	go func() {
		if err = server.Start(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Println(err.Error())
		}
	}()

	shutdown, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-shutdown.Done()
	log.Println("shutdown signal received")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err = server.Stop(ctx); err != nil {
		log.Println(err.Error())
	}

	log.Println("server shutdown completed")
}
