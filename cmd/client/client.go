package main

import (
	"context"
	"flag"
	"io"
	"log"
	"strconv"

	proto "grpcExample/pkg/grpc"

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

	end, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	req := proto.Number{Start: 0, End: end}

	c := proto.NewFlowClient(conn)
	stream, err := c.GetData(context.Background(), &req)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(stream)

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			log.Printf("finished")
			return
		}
		if err != nil {
			log.Fatalf("cannot receive %v", err)
		}
		log.Printf("Resp received: %d", resp.Numb)
	}
}
