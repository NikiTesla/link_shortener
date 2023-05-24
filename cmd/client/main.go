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
	address = "localhost:8080"
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

	// Saving link
	save_response, err := c.SaveOriginal(ctx, &pb.SaveOriginalRequest{OriginalLink: "OK.ru"})
	if err != nil {
		log.Printf("Cannot save original link: %v", err)
	} else {
		log.Printf("Shorted link: %s", save_response.GetShortedLink())
	}

	// Getting link
	get_response, err := c.GetOriginal(ctx, &pb.GetOriginalRequest{ShortedLink: "OyZz51wdUO"})
	if err != nil {
		log.Printf("Cannot get original link: %v", err)
	} else {
		log.Printf("Original link: %s", get_response.GetOriginalLink())
	}
}
