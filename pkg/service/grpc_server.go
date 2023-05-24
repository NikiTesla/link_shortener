package service

import (
	"fmt"
	"log"
	"net"

	pb "github.com/NikiTesla/link_shortener/api"
	"github.com/NikiTesla/link_shortener/pkg/environment"
	"google.golang.org/grpc"
)

// ShortenerServer is an implementation of rpc service
type ShortenerServer struct {
	pb.UnimplementedShortenerServer
	env *environment.Environment
}

// NewShortenerServer creates ShortenerServer with environment
func NewShortenerServer(env *environment.Environment) *ShortenerServer {
	return &ShortenerServer{
		env: env,
	}
}

// RunShortenerServer starts listening on port and registers methods of ShortenerServer as procedures
func RunShortenerServer(shortener *ShortenerServer, s *grpc.Server) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", shortener.env.Port))
	if err != nil {
		return fmt.Errorf("listening failed: %s", err)
	}

	pb.RegisterShortenerServer(s, shortener)

	log.Printf("server listening at %v", lis.Addr())
	return fmt.Errorf("serving failed: %v", s.Serve(lis))
}
