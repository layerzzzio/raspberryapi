package transport

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/raspibuddy/rpi/pkg/api/metrics/host"
	// "golang.org/x/net/websocket"
)

// HTTP is a struct implementing a core application service.
type HTTP struct {
	svc host.Service
}

// NewHTTP creates new user http service
func NewHTTP(svc host.Service, r *echo.Group) {
	h := HTTP{svc}
	cr := r.Group("/hosts")
	cr.GET("", h.list)
	cr.GET("-ws", h.listws)
}

func (h *HTTP) list(ctx echo.Context) error {
	result, err := h.svc.List()
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, result)
}

var upgrader = websocket.Upgrader{}

func (h *HTTP) listws(ctx echo.Context) error {
	ws, err := upgrader.Upgrade(ctx.Response(), ctx.Request(), nil)

	if err != nil {
		return err
	}

	defer ws.Close()

	for {
		// Write data to client
		msg, errR := h.svc.List()
		if errR != nil {
			ctx.Logger().Error(err)
			// break
		}

		err := ws.WriteJSON(msg)

		if err != nil {
			ctx.Logger().Error(err)
			// break
		}

		time.Sleep(15 * time.Second)
		fmt.Println(fmt.Sprint(time.Now()) + " : host sleep 15 sec")

		// Read
		_, msgR, errR := ws.ReadMessage()
		if errR != nil {
			ctx.Logger().Error(errR)
			break
		}
		fmt.Printf("%s\n", msgR)
	}
	return nil
}
