package main

import (
	"context"
	"log"
	"time"

	pb "github.com/NikiTesla/link_shortener/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Cannot connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewShortenerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SaveOriginal(ctx, &pb.SaveOriginalRequest{OriginalLink: "http://127.0.0.1:8000/hello"})
	if err != nil {
		log.Printf("Cannot save original link: %v", err)
	}

	log.Printf("Shorted link: %v", r.ShortedLink)
}
