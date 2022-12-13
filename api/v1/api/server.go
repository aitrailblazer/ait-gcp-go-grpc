package api

import (
	"fmt"
	"log"
	"net/http"
	"time"

	models "github.com/aitrailblazer/ait-gcp-go-grpc/api/v1/models"

	"github.com/labstack/echo/v4"
)

const VERSION = "1.02"

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
//		localhost:8080/v1/ping
func (h *Handler) AitrailblazerServiceSendPing(ctx echo.Context, params models.AitrailblazerServiceSendPingParams) error {

	var i int32 = 369
	msec := time.Now()
	log.Printf("AitrailblazerServiceSend VERSION %s ", VERSION) // <3>

	fmt.Println("AitrailblazerServiceSend: params ", params)
	// message := "pongTest"
	message := params.Message
	fmt.Println("AitrailblazerServiceSend: message ", message)
	var v string = VERSION
	pong := models.Pong{
		Index:      &i,
		Message:    message,
		ReceivedOn: &msec,
		Ver:        &v,
	}
	return ctx.JSON(http.StatusOK, pong)

}

// (GET /v1/shelves/{shelf})
func (h *Handler) AitrailblazerServiceGetShelf(ctx echo.Context, shelf int64) error {

	log.Printf("AitrailblazerServiceGetShelf VERSION %s ", VERSION) // <3>

	fmt.Println("AitrailblazerServiceGetShelf: shelf ", shelf)
	i := shelf

	theme := "My Bookshelf"
	bookshelf := models.Shelf{
		Id:    &i,
		Theme: &theme,
	}
	return ctx.JSON(http.StatusOK, bookshelf)
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
