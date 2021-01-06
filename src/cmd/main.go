package main

import (
	"github.com/frantacer/go-backend-template/src/infrastructure/http"

	"context"
	"fmt"
	"log"
)

func main() {
	ctx := context.Background()

	httpHandler := http.NewHandler()
	server := http.NewServer(ctx, httpHandler)

	port := 3000

	readyCh := make(chan struct{})
	go func() {
		<-readyCh
		log.Println(fmt.Sprintf("system ready to serve on port %d", port))
	}()
	if err := server.ListenAndServe(port, readyCh); err != nil {
		log.Fatal(err.Error())
	}
}
