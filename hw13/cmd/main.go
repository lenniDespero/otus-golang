package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/lenniDespero/otus-golang/hw13/internal/calendar"
	"github.com/lenniDespero/otus-golang/hw13/internal/pkg/config"
	"github.com/lenniDespero/otus-golang/hw13/internal/pkg/storage"

	"github.com/gorilla/mux"
	"github.com/lenniDespero/otus-golang/hw13/internal/pkg/logger"
	"github.com/spf13/pflag"
)

func main() {
	var configPath = flag.String("config", "../config/application.yml", "path to configuration flag")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	flag.Parse()
	conf := config.GetConfigFromFile(*configPath)
	logger.Init(conf.Log.LogLevel, conf.Log.LogFile)
	inMemoryStorage := storage.New()
	_ = calendar.New(inMemoryStorage)
	logger.Info("Calendar was created")
	InitServer(conf.HttpListen.Ip, conf.HttpListen.Port)
}

//Init http server
func InitServer(listenIp string, listenPort string) {
	router := mux.NewRouter()
	router.HandleFunc("/hello", hello).Methods("GET")

	srv := &http.Server{
		Addr:         listenIp + ":" + listenPort,
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			logger.Fatal("Error while starting HTTP server", "error", err)
		}
	}()
	logger.Info("HTTP server started on host: " + listenIp + ", port: " + listenPort)

	<-done
	logger.Info("HTTP server stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("Server shutdown failed", "error", err)
	}
	logger.Info("HTTP server exited properly")
}

func hello(w http.ResponseWriter, r *http.Request) {
	logger.Info("Incoming message",
		"host", r.Host,
		"url", r.URL.Path)
	message := []byte(`{"Message":"Hello world"}`)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", message)
}
