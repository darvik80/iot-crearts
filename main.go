package main

import (
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "<html><title>iot.crearts.xyz</title><body>Stub page</body></html>")
}

func main() {
	var st = time.Now()
	runtime.GOMAXPROCS(runtime.NumCPU())
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("Hello World")

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)

	log.Infof("[app] start app, %dms", time.Now().Sub(st).Milliseconds())

	srv := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
