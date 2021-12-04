package adder

import (
	"context"
	"fmt"
	"newAdder/pkg/api"
)

//GRPCServer ...
type GRPCServer struct{}

func (s GRPCServer) Add(ctx context.Context, req *api.AddRequest) (*api.AddResponse, error) {

	fmt.Println("Add")
	return &api.AddResponse{Result: req.GetX() + req.GetY()}, nil

}
