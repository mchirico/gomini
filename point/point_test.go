package point

import (
	"context"
	"fmt"
	"github.com/mchirico/tlib/util"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestMainListen(t *testing.T) {
	defer util.NewTlib().ConstructDir()()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		api := NewPointFile("./data.csv")
		api.MainListen(ctx, ":3000")
		for {
			select {

			case <-ctx.Done():
				return
			}
		}
	}()

	res, err := http.Get("http://localhost:3000/")
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
