package service

import (
	"context"
	"errors"
	"fmt"
	"math/rand"

	pb "github.com/NikiTesla/link_shortener/api"
	"github.com/NikiTesla/link_shortener/pkg/repository"
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
		return &pb.SaveOriginalResponse{}, fmt.Errorf("recieved empty link")
	}

	shortedLink, err := s.generateShortenedLink(originalLink)
	if err != nil {
		return &pb.SaveOriginalResponse{}, fmt.Errorf("cannot save link in database: %s", err)
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
			fmt.Errorf("recieved empty link")
	}

	shortedLink, err := s.env.DB.GetLink(shortedLink)
	if err != nil {
		return &pb.GetOriginalResponse{},
			fmt.Errorf("cannot get original link from database: %s", err)
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
