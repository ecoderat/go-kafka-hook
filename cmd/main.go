package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"

	logg "github.com/ecoderat/go-kafka-hook/pkg/logwrapper"
)

type application struct {
	logger *logg.StandardLogger
	//hook   *kafkalogrus.KafkaLogrusHook
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")

	app := &application{
		logger: logg.NewLogger(),
	}

	srv := &http.Server{
		Addr:    *addr,
		Handler: app.routes(),
	}

	// Create a new hook
	hook, err := logg.NewHook()
	if err != nil {
		panic(err)
	}

	// Add hook to logger
	app.logger.Hooks.Add(hook)

	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	mw := io.MultiWriter(os.Stdout, file)
	app.logger.SetOutput(mw)
	// app.logger.SetOutput(file)

	ctx := context.Background()
	go consume(ctx)

	fmt.Printf("Starting server on %s\n", *addr)
	srv.ListenAndServe()
}
