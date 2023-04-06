package revers

import (
	"grpcExample/pkg/grpc"
	"log"
	"sync"
	"time"
)

type Flow struct{}

//Реализация интерфейса Flow
func (s *Flow) GetData(req *grpc.Number, srv grpc.Flow_GetDataServer) error {
	log.Printf("fetch response start: %d, end: %d", req.Start, req.End)

	var wg sync.WaitGroup
	for i := 0; i <= int(req.End); i++ {
		wg.Add(1)
		go func(count int64) {
			defer wg.Done()
      
			time.Sleep(time.Duration(count) * time.Second)
			
			resp := grpc.Response{Numb: count}
			if err := srv.Send(&resp); err != nil {
				log.Printf("send error %v", err)
			}			
		}(int64(i))
	}

	wg.Wait()
	return nil
}
