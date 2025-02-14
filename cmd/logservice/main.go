package main

import (
	"LR/go-simple-service/log"
	"LR/go-simple-service/service"
	"context"
	"fmt"
	stlog "log"
)

func main() {
	log.Run("./distributed.log")
	host, port := "localhost", "4000"

	ctx, err := service.Start(
		context.Background(),
		"log Service",
		host,
		port,
		log.RegisterHandlers,
	)
	if err != nil {
		stlog.Fatalln(err)
	}
	<-ctx.Done()

	fmt.Println("Shutting down log service")
}
