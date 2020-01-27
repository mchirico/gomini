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

func (api *POINT) MainListen(ctx context.Context) {

	r := chi.NewRouter()
	r.Get("/", handler.TextMsg("hit point v1.0"))
	r.Get("/data", handler.DataSend(api.File))

	server := &http.Server{Addr: ":3000", Handler: r}

	server.ListenAndServe()
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("error")
	}
}
