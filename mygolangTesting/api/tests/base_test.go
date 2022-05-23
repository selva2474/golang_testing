package tests

import (
	"fmt"
	"os"
	"testing"

	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/selva/golang-testing/src/api/app"
)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	fmt.Println("now starting application")
	go app.StartApp()
	os.Exit(m.Run())
	fmt.Println("Application STarted")
}
