package main

import (
	"context"
	"embed"
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"io/fs"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"time"
)

//go:embed frontend/build/*
var content embed.FS

type LogHandler struct {
	handler http.Handler
}

func (h *LogHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.Infof("handle: " + req.URL.Path)
	h.handler.ServeHTTP(w, req)
}

func rootHandler() http.Handler {
	fsys := fs.FS(content)
	static, _ := fs.Sub(fsys, "frontend/build")

	return &LogHandler{http.FileServer(http.FS(static))}
}

func main() {
	var st = time.Now()
	runtime.GOMAXPROCS(runtime.NumCPU())
	sigs := make(chan os.Signal)
	signal.Notify(sigs, os.Interrupt)

	r := mux.NewRouter()
	r.HandleFunc("/api/task", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "success")
	})
	r.PathPrefix("/").Handler(rootHandler())

	log.Infof("Start app, %dms", time.Now().Sub(st).Milliseconds())

	srv := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {
		log.Printf("Server listening %s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				log.Fatal(err)
			}
		}
	}()

	<-sigs

	log.Printf("Server shutting down ...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	log.Printf("Server down ...")
}
