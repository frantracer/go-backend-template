package main

import (
	"github.com/frantacer/go-backend-template/src/application/handlers"
	"github.com/frantacer/go-backend-template/src/infrastructure/http"
	"github.com/frantacer/go-backend-template/src/infrastructure/inmemory"

	"context"
	"fmt"
	"log"
)

func main() {
	ctx := context.Background()

	uowc := inmemory.NewUnitOfWorkCreator()

	appHandlers := http.ApplicationHandlers{
		FindTasksHandler:  handlers.NewFindTasksCommandHandler(uowc),
		InsertTaskHandler: handlers.NewInsertTaskCommandHandler(uowc),
	}

	httpHandler := http.NewHandler(appHandlers)
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
