package revers

import (
	"context"
	"grpcExample/grpc"
)

type GRPCServer struct{}

//Реализация интерфейса Do revers
func (s *GRPCServer) Do(ctx context.Context, req *grpc.Request) (*grpc.Response, error) {
	strRev := reverse(req.Message)

	res := grpc.Response{Message: strRev }
	
	return &res, nil
}

func reverse(s string) string {
    rns := []rune(s) // convert to rune
    for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
  
        rns[i], rns[j] = rns[j], rns[i]
    }
  
    return string(rns)
}