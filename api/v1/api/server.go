package api

import (
	"log"
	"net/http"
	"time"

	models "github.com/aitrailblazer/ait-gcp-go-grpc/api/v1/models"

	"github.com/labstack/echo/v4"
)

// type PongStore struct {
// 	Pong models.Pong
// 	// Pongs  map[int64]models.Pong
// 	// NextId int64
// 	// Lock   sync.Mutex
// }

// func NewPongStore() *PongStore {
// 	pong := Pong{
// 		Index:      1,
// 		Message:    "Pong!",
// 		ReceivedOn: time.Now(),
// 	}
// 	return pong
// }
// func NewPongStore() *PongStore {
// 	pong := Pong{
// 		Index:      1,
// 		Message:    "Pong!",
// 		ReceivedOn: time.Now(),
// 	}
// 	return pong
// }

type Handler struct{}

func NewHandler() Handler {
	// pong := Pong{
	// 	Index:      1,
	// 	Message:    "Pong!",
	// 	ReceivedOn: time.Now(),
	// }
	handler := Handler{}
	return handler
}

//	curl \
//	  -X GET \
//	  http://localhost:1323\?name\=Joe
func (h *Handler) AitrailblazerServiceSend(ctx echo.Context, params models.AitrailblazerServiceSendParams) error {

	var i int32 = 1
	msec := time.Now()
	log.Printf("ver %s ", "1.0") // <3>

	log.Println("params ", params)
	// message := "pongTest"
	message := params.Message
	log.Println("message ", message)

	pong := models.Pong{
		Index:      &i,
		Message:    message,
		ReceivedOn: &msec,
	}
	return ctx.JSON(http.StatusOK, pong)

}

// (GET /v1/ping)
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
