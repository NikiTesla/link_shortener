package service

import (
	"fmt"
	"log"
	"net"

	pb "github.com/NikiTesla/link_shortener/api"
	"google.golang.org/grpc"
)

// ShortenerServer is an implementation of rpc service
type ShortenerServer struct {
	pb.UnimplementedShortenerServer
}

// RunShortenerServer starts listening on port and registers methods of ShortenerServer as procedures
func RunShortenerServer(port int, s *grpc.Server) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		return fmt.Errorf("listening failed: %v", err)
	}

	pb.RegisterShortenerServer(s, &ShortenerServer{})

	log.Printf("server listening at %v", lis.Addr())
	return fmt.Errorf("serving failed: %v", s.Serve(lis))
}
