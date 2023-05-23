package service

import (
	"context"
	"log"

	pb "github.com/NikiTesla/link_shortener/api"
)

func (s *ShortenerServer) SaveOriginal(ctx context.Context, in *pb.SaveOriginalRequest) (*pb.SaveOriginalResponse, error) {
	log.Printf("Recieved: %v", in.GetOriginalLink())

	log.Print("Saved")

	return &pb.SaveOriginalResponse{
		ShortedLink: "hellohello",
	}, nil
}

func (s *ShortenerServer) GetOriginal(ctx context.Context, in *pb.GetOriginalRequest) (*pb.GetOriginalResponse, error) {
	log.Printf("Recieved: %v", in.GetShortedLink())

	log.Printf("Got")

	return &pb.GetOriginalResponse{
		OriginalLink: "http://127.0.0.1:8000/hello",
	}, nil
}
