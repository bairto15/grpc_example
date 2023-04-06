package revers

import (
	"grpcExample/pkg/grpc"
	"log"
	"time"
)

type Flow struct{}

//Реализация интерфейса Flow
func (s *Flow) GetData(req *grpc.Number, srv grpc.Flow_GetDataServer) error {
	log.Printf("fetch response start: %d, end: %d", req.Start, req.End)

	for i := 0; i <= int(req.End); i++ {
		time.Sleep(1 * time.Second)

		resp := grpc.Response{Numb: int64(i)}
		err := srv.Send(&resp)
		if err != nil {
			log.Printf("send error %v", err)
		}
	}
	
	return nil
}
