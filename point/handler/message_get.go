package handler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func TextMsg(msg string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hit point v1.0"))
	}
}

func DataSend(file string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		data := []byte("test")
		err := ioutil.WriteFile("bbdata.csv", data, 0600)
		fmt.Printf("file:  %s\n", file)
		if err != nil {
			log.Fatalf("err: %s\n", err)
		}

		w.Write([]byte(ReadFile(file)))
	}
}

func ReadFile(file string) string {
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("File not found")
		return "File not found"
	}
	return string(dat)
}
