package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Dev-Siri/gdn/api"
	"github.com/Dev-Siri/gdn/db"
	"github.com/Dev-Siri/gdn/env"
	"github.com/Dev-Siri/gdn/logging"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func main() {
	port := env.GetPort()
	addr := ":" + port

	if err := db.SetupConfig(); err != nil {
		log.Fatalf("ERROR: %v", err)
	}

	if err := db.InitStorage(); err != nil {
		log.Fatal("ERROR: Failed to create cache folder")
	}

	if db.CDNConfig.Log {
		if err := logging.InitLogger(); err != nil {
			log.Fatal("ERROR: Failed to initialize logger")
		}

		logging.FlushLogger()
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-interrupt
		log.Print("INFO: Gracefully shutting down.")

		os.Exit(0)
	}()

	log.Print("GDN server started on port " + port)

	router := router.New()

	router.HandleOPTIONS = true
	router.HandleMethodNotAllowed = true

	api.RegisterRoutes(router)

	log.Fatal(fasthttp.ListenAndServe(addr, router.Handler))
}
