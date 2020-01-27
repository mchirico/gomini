package cmd

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

func TestEntryPt(t *testing.T) {

	tlib := &util.Tlib{FindFunc: util.FindFile, MockDir: "../test-fixtures", SubDir: "TestEntryPt"}
	defer util.NewTlib(tlib).ConstructDir()()
	expected_value := "2,two"

	file := util.PWD() + "/data.csv"
	util.WriteString(file, expected_value, 0644)
	fmt.Printf("here: %v\n", file)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		EntryPt(ctx, file, ":3001")
		for {
			select {

			case <-ctx.Done():
				return
			}
		}
	}()

	res, err := http.Get("http://localhost:3001/data")
	if err != nil {
		log.Fatal(err)
	}
	result, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	if !strings.Contains(string(result), expected_value) {
		t.Fatalf(string(result))
	}

	fmt.Printf("%s\n", result)

	time.Sleep(9 * time.Second)

}
