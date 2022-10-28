package sqlstore_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("TEST_DATABASE_URL")

	if databaseURL == "" {
		databaseURL = "host=localhost port=5432 user=kshanti password=qwerty1234 dbname=mydb sslmode=disable"
	}

	res := m.Run()

	os.Exit(res)
}
