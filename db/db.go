package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

func Conn(table string,file string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
	"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	"content" TEXT,
    timeEnter DATE
	);

`, table)

	stmt, err := db.Prepare(query)
	if err != nil {
		fmt.Fprintf(os.Stderr, "db.Prepare(query): %s", err)
		return nil, err
	}
	_, err = stmt.Exec()
	if err != nil {
		fmt.Fprintf(os.Stderr, "stmt.Exec(): %s", err)
		return nil, err
	}

	query = fmt.Sprintf(`
	CREATE TRIGGER IF NOT EXISTS insert_%s_timeEnter AFTER  INSERT ON %s
     BEGIN
      UPDATE %s SET timeEnter = DATETIME('NOW')  WHERE rowid = new.rowid;
     END;
`, table, table, table)

	stmt, err = db.Prepare(query)

	if err != nil {
		return nil, err
	}
	stmt.Exec()

	return db, nil

}

func Insert(db *sql.DB, table string, content string) (sql.Result, error) {

	query := fmt.Sprintf(`
	INSERT INTO %s (content) values (?)
`, table)
	stmt, err := db.Prepare(query)

	if err != nil {
		return nil, err
	}
	r, err := stmt.Exec(content)
	return r, err

}

func Query(db *sql.DB, table string) (*sql.Rows, error) {
	query := fmt.Sprintf("select id,content,timeEnter from %s", table)
	return db.Query(query)

}
