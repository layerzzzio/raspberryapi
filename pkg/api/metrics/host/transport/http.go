package transport

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
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
		fmt.Print(time.Now().Date())
		// Write data to client
		msg, errR := h.svc.List()
		if errR != nil {
			ctx.Logger().Error(err)
			break
		}

		fmt.Print(" ---> host data in \n")

		err := ws.WriteJSON(msg)

		fmt.Print("host data out \n")

		if err != nil {
			ctx.Logger().Error(err)
			break
		}

		time.Sleep(15 * time.Second)
		fmt.Print("process sleep 15 sec")
	}
	return nil
}
