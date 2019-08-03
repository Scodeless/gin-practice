package server

import (
	"context"
	"fmt"
	pb "gin-practice/pkg/rpc/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":50051"
)

// server is used to implement auth.AuthServer
type Server struct{}

// Login implements auth.AuthServer
func (s *Server) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginReply, error) {
	if in.UserId == 1 {
		return &pb.LoginReply{UserId: 1, UserName: "songwei"}, nil
	} else {
		return &pb.LoginReply{}, fmt.Errorf("登陆失败，用户错误")
	}
}

func StartRpc()  {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAuthServer(s, &Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	fmt.Println("rpc start success")
}