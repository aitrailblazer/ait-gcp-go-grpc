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

// This function wraps sending of an error in the Error format, and
// handling the failure to marshal that.
// func sendServiceError(ctx echo.Context, code int, message string) error {
// 	petErr := models.Error{
// 		Code:    int32(code),
// 		Message: message,
// 	}
// 	err := ctx.JSON(code, petErr)
// 	return err
// }

// curl localhost:8080/v1/ping
// curl localhost:8080/v1/ping?message=test
func (h *Handler) AitrailblazerServiceSendPing(ctx echo.Context, params models.AitrailblazerServiceSendPingParams) error {

	var i int32 = 369
	msec := time.Now()
	log.Printf("AitrailblazerServiceSendPing VERSION %s ", VERSION) // <3>

	message := params.Message
	fmt.Println("AitrailblazerServiceSendPing: message ", *message)
	var p string = "path parameter=" + *message + " Ping from AitrailblazerServiceSendPing"

	var v string = VERSION
	pong := models.Pong{
		Index:      &i,
		Message:    &p,
		ReceivedOn: &msec,
		Ver:        &v,
	}
	return ctx.JSON(http.StatusOK, pong)
}

// (GET /echo/:message)
// curl localhost:8080/echo/test
// curl localhost:8080/echo/test\?value\=test1234
func (h *Handler) AitrailblazerServiceEcho(ctx echo.Context, message string, params models.AitrailblazerServiceEchoParams) error {

	// ------------- Path parameter "message" -------------
	fmt.Println("AitrailblazerServiceEcho: Path parameter ", message)
	// paramsValue := params.Value

	if params.Value != nil {
		fmt.Println("AitrailblazerServiceEcho: params.Value ", *params.Value)
		message = "path parameter=" + message + " query=" + *params.Value + " Echo from AitrailblazerServiceEcho"
	} else {
		message = "path parameter=" + message + " Echo from AitrailblazerServiceEcho"
	}

	echoMessage := models.EchoMessage{
		Value: &message,
	}
	return ctx.JSON(http.StatusOK, echoMessage)
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
