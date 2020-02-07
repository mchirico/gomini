package point

import (
	"context"
	"fmt"
	"github.com/mchirico/tlib/util"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

/*
See the following for better testing:
https://golang.org/src/net/http/httptest/example_test.go
*/
func TestMainListen(t *testing.T) {
	tlib := &util.Tlib{FindFunc: util.FindFile, MockDir: "../test-fixtures", SubDir: "TestPoint"}
	defer util.NewTlib(tlib).ConstructDir()()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		api := NewPointFile("./data.csv")
		api.MainListen(ctx, ":3020")
		for {
			select {

			case <-ctx.Done():
				return
			}
		}
	}()

	res, err := http.Get("http://localhost:3020/")
	if err != nil {
		log.Fatal(err)
	}
	result, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	if !strings.Contains(string(result), "hit point") {
		t.Fatalf(string(result))
	}

	fmt.Printf("%s\n", result)

	time.Sleep(7 * time.Second)
}

func TestNewMux(t *testing.T) {
	r := NewMux("../test-fixtures/data.csv")
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	if !strings.Contains(response.Body.String(), "hit point") {
		t.Fatalf(response.Body.String())
	}

	request, _ = http.NewRequest("GET", "/data", nil)
	r.ServeHTTP(response, request)

	if !strings.Contains(response.Body.String(), "2,two") {
		t.Fatalf(response.Body.String())
	}
}
