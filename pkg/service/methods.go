package service

import (
	"context"
	"errors"
	"log"
	"math/rand"

	pb "github.com/NikiTesla/link_shortener/api"
	"github.com/NikiTesla/link_shortener/pkg/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	chars      = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
	linkLength = 10
)

// SaveOriginal gets context and SaveOriginalRequest with original link inside.
// Calls generateShortened to create and save unique link in database.
// returns SaveOriginalResponse with shortened link inside
func (s *ShortenerServer) SaveOriginal(ctx context.Context, in *pb.SaveOriginalRequest) (*pb.SaveOriginalResponse, error) {
	originalLink := in.GetOriginalLink()
	// log.Printf("Recieved: %v\n", originalLink)
	if originalLink == "" {
		return &pb.SaveOriginalResponse{}, status.Error(codes.InvalidArgument, "Empty link is given")
	}

	shortedLink, err := s.generateShortenedLink(originalLink)
	if err != nil {
		log.Printf("cannot save link in database: %s\n", err)
		return &pb.SaveOriginalResponse{}, status.Error(codes.Internal, "Cannot save link in database")
	}

	return &pb.SaveOriginalResponse{
		ShortedLink: shortedLink,
	}, nil
}

// GetOriginal gets context and GetOriginalRequest with shortened link inside.
// Request repository database for original one.
// returns GetOriginalResponse with original link inside and error
func (s *ShortenerServer) GetOriginal(ctx context.Context, in *pb.GetOriginalRequest) (*pb.GetOriginalResponse, error) {
	shortedLink := in.GetShortedLink()
	// log.Printf("Recieved: %v\n", shortedLink)

	if shortedLink == "" {
		return &pb.GetOriginalResponse{},
			status.Error(codes.InvalidArgument, "Empty link is given")
	}

	shortedLink, err := s.env.DB.GetLink(shortedLink)
	if errors.Is(err, repository.ErrLinkNotFound) {
		return &pb.GetOriginalResponse{},
			status.Error(codes.NotFound, "No such shortened link in database")
	}
	if err != nil {
		log.Printf("cannot get original link from database: %s", err)
		return &pb.GetOriginalResponse{},
			status.Error(codes.Internal, "Cannot get original link from database")
	}

	return &pb.GetOriginalResponse{
		OriginalLink: shortedLink,
	}, nil
}

// generateShortenedLink is an internal method for short link creatiion and saving.
// Randomize each symbol of const linkLength from const chars list
func (s *ShortenerServer) generateShortenedLink(originalLink string) (string, error) {
	linkBytes := make([]byte, linkLength)
	for i := 0; i < linkLength; i++ {
		linkBytes[i] = chars[rand.Intn(len(chars))]
	}

	shortedLink := string(linkBytes)
	err := s.env.DB.SaveLink(originalLink, shortedLink)
	if errors.Is(err, repository.ErrLinkAlreadyExists) {
		return s.generateShortenedLink(originalLink)
	}
	if err != nil {
		return "", err
	}

	return shortedLink, nil
}
