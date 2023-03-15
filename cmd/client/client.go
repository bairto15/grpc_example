package main

import (
	"context"
	"flag"
	"log"

	proto "grpcExample/grpc"

	"google.golang.org/grpc"
)

func main() {
	flag.Parse()

	if flag.NArg() < 1 {
		log.Fatal("Не указаны аргументы")
	}

	str := flag.Arg(0)

	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	req := proto.Request{ Message: str }

	c := proto.NewReverseClient(conn)
	res, err := c.Do(context.Background(), &req)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(res.Message)
}