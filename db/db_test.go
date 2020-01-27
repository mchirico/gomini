package db

import (
	"fmt"
	"github.com/mchirico/tlib/util"
	"testing"
)

func TestConn(t *testing.T) {

	tlib := &util.Tlib{FindFunc: util.FindFile, MockDir: "../test-fixtures", SubDir: "DB"}
	defer util.NewTlib(tlib).ConstructDir()()

	table := "test"
	file := "test.sqlite"
	db, err := Conn(table, file)
	defer db.Close()

	if err != nil {
		t.Fatalf("No connection")
	}

	_, err = Insert(db, table, "stuff")
	if err != nil {
		t.Fatalf("insert")
	}

	rows, err := Query(db, table)
	if err != nil {
		t.Fatalf("failed Query(db)")
	}
	var id int
	var content string
	var timeEnter string
	for rows.Next() {
		rows.Scan(&id, &content, &timeEnter)
		fmt.Printf("id:%d, content: %s, timeEnter: %s\n", id, content, timeEnter)

	}

}
