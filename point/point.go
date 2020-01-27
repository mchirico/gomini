package point

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/mchirico/gomini/point/handler"
	"log"
	"net/http"
)

type POINT struct {
	File string
}

func NewPoint() *POINT {
	return &POINT{File: "/data.csv"}
}

func NewPointFile(file string) *POINT {
	return &POINT{File: file}
}

func (api *POINT) MainListen(ctx context.Context, port string) {

	r := chi.NewRouter()
	r.Get("/", handler.TextMsg("hit point v1.0"))
	r.Get("/data", handler.DataSend(api.File))

	// ":3000"
	server := &http.Server{Addr: port, Handler: r}

	server.ListenAndServe()
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("error")
	}
}
