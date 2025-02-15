package service

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func Start(ctx context.Context, serviceName, host, port string, registerHandlers func()) (context.Context, error) {
	registerHandlers()
	ctx = startService(ctx, serviceName, host, port)
	return ctx, nil
}

func startService(ctx context.Context, serviceName, host, port string) context.Context {
	ctx, cancel := context.WithCancel(ctx)

	var srv http.Server
	srv.Addr = host + ":" + port

	go func() {
		log.Println(srv.ListenAndServe())
		cancel()
	}()

	go func() {
		log.Printf("%v started. Press any key to stop.\n", serviceName)
		var s string
		fmt.Scanf("%s", &s)
		srv.Shutdown(ctx)
		cancel()
	}()

	return ctx
}
