package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"animo.com/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	ph := handlers.NewProducts(l)
	//hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)
	sm := http.NewServeMux()
	sm.Handle("/", ph)
	sm.Handle("/goodbye", gh)
	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal("test %s\n", err)

		}
	}()
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)
	sig := <-sigChan
	l.Println("Rcd..Eceived terminate,graceful shutdown", sig)
	log.Println("Got signal:", sig)
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
	//http.ListenAndServe(":9090", sm)

	/*
			http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				log.Println("Hello World")
				d, err := ioutil.ReadAll(r.Body)
				if err != nil {
					http.Error(w, "error Occured", http.StatusBadRequest)

					return
				}go
				log.Printf("Data %s\n", d)

			})
		http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
			log.Println("Goodbye  World")

		})
		http.ListenAndServe(":9090", nil)*/
}
