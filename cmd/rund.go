package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"rund/internal/api"
	"syscall"
)

const (
	currentVersion = "0.0.1"
)

const usage = `Usage: alfred [options...]:
  -h, --ip <ip>        the IP address or hostname on which the API server will listen
  -p, --port <port>    the port number on which the API server will listen
  -c, --config <file>  the projects configuration file to load
  -v, --verbose        increase output verbosity
  -V, --version        display version information and exit
  -h, --help           display this help text and exit
`

type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

func main() {
	var ip, configFilePath string
	var port int
	var verbose, help, version bool

	flag.StringVar(&ip, "ip", "0.0.0.0", "")
	flag.StringVar(&ip, "i", "0.0.0.0", "")
	flag.StringVar(&configFilePath, "config", "", "")
	flag.StringVar(&configFilePath, "c", "", "")
	flag.IntVar(&port, "port", 8080, "")
	flag.IntVar(&port, "p", 8080, "")
	flag.BoolVar(&help, "help", false, "")
	flag.BoolVar(&version, "version", false, "")
	flag.BoolVar(&version, "V", false, "")
	flag.BoolVar(&verbose, "verbose", false, "")
	flag.BoolVar(&verbose, "v", false, "")
	flag.Usage = func() {
		fmt.Print(usage)
	}
	flag.Parse()

	addr := fmt.Sprintf("%s:%d", ip, port)

	if help {
		flag.Usage()
		os.Exit(0)
	}

	if version {
		fmt.Println("alfred", currentVersion)
		os.Exit(0)
	}

	if !verbose {
		log.SetOutput(io.Discard)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/api/webhook/{provider}/{project}", makeHandler(api.NewWebhookHandler().Handle))

	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	log.Printf("API server started! Listening on %s", addr)

	shutdown, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	<-shutdown.Done()

	if err := server.Shutdown(context.Background()); err != nil {
		log.Printf("API server shutdown with error %v", err)
	}
}

func makeHandler(h HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			slog.Error(r.URL.Path, "message", err)
		}
	}
}
