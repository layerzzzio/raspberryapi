package transport

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/raspibuddy/rpi/pkg/api/metrics/process"
)

// HTTP is a struct implementing a core application service.
type HTTP struct {
	svc process.Service
}

// NewHTTP creates new user http service
func NewHTTP(svc process.Service, r *echo.Group) {
	h := HTTP{svc}
	cr := r.Group("/processes")
	cr.GET("", h.list)
	cr.GET("-ws", h.listws)
	cr.GET("/:id", h.view)
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
		// Compose message to write to client
		msg, _ := h.svc.List()
		// if errP != nil {
		// 	ctx.Logger().Error(errP)
		// }

		err := ws.WriteJSON(msg)
		if err != nil {
			ctx.Logger().Error(err)
		}

		time.Sleep(30 * time.Second)
		// fmt.Println(fmt.Sprint(time.Now()) + " : process sleep 30 sec")

		// Read incoming message from client including closing message
		_, _, errR := ws.ReadMessage()
		if errR != nil {
			ctx.Logger().Error(errR)
			break
		}
		// fmt.Printf("%s\n", msgR)
	}
	return nil
}

func (h *HTTP) view(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request due to an invalid process id - should be an integer")
	}

	result, err := h.svc.View(int32(id))
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, result)
}
