package main

import (
	"encoding/gob"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/RAprogramm/smSystem/internal/config"
	"github.com/RAprogramm/smSystem/internal/handlers"
	"github.com/RAprogramm/smSystem/internal/models"
	"github.com/alexedwards/scs/v2"
	"github.com/pusher/pusher-http-go"
)

var (
	app           config.AppConfig
	repo          *handlers.DBRepo
	session       *scs.SessionManager
	preferenceMap map[string]string
	wsClient      pusher.Client
)

const (
	// SMSystemVersion ...
	SMSystemVersion   = "1.0.0"
	maxWorkerPoolSize = 5
	maxJobMaxWorkers  = 5
)

func init() {
	gob.Register(models.User{})
	_ = os.Setenv("TZ", "Korea/Seoul")
}

// main is the application entry point
func main() {
	// set up application
	insecurePort, err := setupApp()
	if err != nil {
		log.Fatal(err)
	}

	// close channels & db when application ends
	defer close(app.MailQueue)
	defer app.DB.SQL.Close()

	// print info
	log.Printf("******************************************")
	log.Printf(
		"** %sSMSystem%s v%s built in %s",
		"\033[31m",
		"\033[0m",
		SMSystemVersion,
		runtime.Version(),
	)
	log.Printf("**----------------------------------------")
	log.Printf("** Running with %d Processors", runtime.NumCPU())
	log.Printf("** Running on %s", runtime.GOOS)
	log.Printf("******************************************")

	// create http server
	srv := &http.Server{
		Addr:              *insecurePort,
		Handler:           routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	log.Printf("Starting HTTP server on port %s....", *insecurePort)

	// start the server
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
