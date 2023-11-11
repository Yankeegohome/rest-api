package nlab

import (
	"fmt"
	"testing"
)

func Testnlab(t *testing.T, databaseURL string) (*Nlab, func(...string)) {
	t.Helper()

	config := NewConfig()
	config.DatabaseURL = databaseURL
	s := New(config)
	if err := s.Open(); err != nil {
		t.Fatal(err)
	}
	return s, func(tables ...string) {
		if len(tables) > 0 {
			if _, err := s.db.Exec(fmt.Sprintf("delete from nlab.usr where date_trunc('day', createdate) = date_trunc('day',current_timestamp)")); err != nil {
				t.Fatal(err)
			}
		}

		s.Close()
	}
}
