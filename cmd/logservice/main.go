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

	ctx, err := service.Start(context.Background(), "Log Service",
		host,
		port,
		log.RegisterHandlers)
	if err != nil {
		stlog.Fatalln(err)
		return
	}
	<-ctx.Done()

	fmt.Println("Shutting down log service")
}
