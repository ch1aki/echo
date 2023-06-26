package main

import (
	"context"
	"echo/gen/api"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
)

func main() {
        target := "localhost:50051"
        conn, err := grpc.Dial(target, grpc.WithInsecure())
        if err != nil {
                log.Fatalf("did not connect: %s", err)
        }
        defer conn.Close()
        client := api.NewEchoServiceClient(conn)
        msg := os.Args[1]
        ctx, cancel := context.WithTimeout(context.Background(), time.Second)
        defer cancel()
        r, err := client.Echo(ctx, &api.EchoRequest{Message: msg})
        if err != nil {
                log.Println(err)
        }
        log.Println(r.GetMessage())
}
