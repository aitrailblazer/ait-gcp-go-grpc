package api

import (
	"fmt"
	"sync"

	models "github.com/aitrailblazer/ait-gcp-go-grpc/api/v1/models"

	"github.com/labstack/echo/v4"
)

// type Pong struct {
// 	Pong   map[int64]models.Pong
// 	NextId int64
// 	Lock   sync.Mutex
// }
// Implements all the handlers in the ServerInterface

// (GET /v1/ping)
func (p *models.ServerInterfaceWrapper) AitrailblazerServiceSend(ctx echo.Context, params models.AitrailblazerServiceSendParams) error {
	var err error
	return fmt.Errorf("error marshaling '@type': %w", err)

}

// func (s *pingService) Send(ctx context.Context, req *pb.Request) (*pb.Response, error) {
// 	log.Print("sending ping response")
// 	return &pb.Response{
// 		Pong: &models.Pong{
// 			Index:      1,
// 			Message:    req.GetMessage(),
// 			ReceivedOn: ptypes.TimestampNow(),
// 		},
// 	}, nil
// }
