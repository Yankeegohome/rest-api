package sqlnlab

import (
	"database/sql"
	"fmt"
	"testing"
)

func TestDB(t *testing.T, databaseURL string) (*sql.DB, func(...string)) {
	t.Helper()

	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		t.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}
	return db, func(tables ...string) {
		if len(tables) > 0 {
			db.Exec(fmt.Sprintf("delete from nlab.usr where date_trunc('day', createdate) = date_trunc('day',current_timestamp)"))
		}
		db.Close()

	}
}
