package main

import (
	"context"
	"echo/gen/api"
)

type echoService struct {
	api.UnimplementedEchoServiceServer
}

func (s *echoService) Echo(ctx context.Context, req *api.EchoRequest) (*api.EchoResponse, error) {
	return &api.EchoResponse{Message: req.GetMessage()}, nil
}
