package point

import (
	"context"
	"github.com/go-chi/chi"
	"io/ioutil"
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
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hit point v1.0"))
	})
	r.Get("/data", func(w http.ResponseWriter, r *http.Request) {

		data := []byte("test")
		err := ioutil.WriteFile("bbdata.csv", data, 0600)
		if err != nil {
			log.Fatalf("err: %s\n", err)
		}

		w.Write([]byte(ReadFile(api.File)))
	})

	server := &http.Server{Addr: ":3000", Handler: r}

	server.ListenAndServe()
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("error")
	}
}

func ReadFile(file string) string {
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return string(dat)
}
