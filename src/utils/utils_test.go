package util

import (
	"flag"
	"fmt"
	"os"
	"testing"
)

func Test_Convert(t *testing.T) {
	const (
		expected ="321"
	)
	oStr :="123"
	resp := Convert([]byte(oStr))
	if resp!=expected{
		t.Errorf("Convert(%s) = %s; expected %s", oStr, resp, expected)
	}
}



func Test_Sum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	expected := 15
	actual := Sum(numbers)
	if actual != expected {
		t.Errorf("Expected the sum of %v to be %d but instead got %d!", numbers, expected, actual)
	}
}

func ExampleSum() {
	numbers := []int{1, 2, 3, 4, 5}
	Sum(numbers)

	// Output:
	// map[string]string{}{"A":"2"}
}

var db struct {
	Url string
}
func TestMain(m *testing.M) {
	// Pretend to open our DB connection
	db.Url = os.Getenv("DATABASE_URL")
	if db.Url == "" {
		db.Url = "localhost:5432"
	}
	flag.Parse()
	exitCode := m.Run()
	// Pretend to close our DB connection
	db.Url = ""
	// Exit
	os.Exit(exitCode)
}
func TestDatabase(t *testing.T) {
	// Pretend to use the db
	fmt.Println(db.Url)
}