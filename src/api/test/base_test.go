package test

import (
	"os"
	"testing"

	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/selva2474/golang_testing/src/api/app"
)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	go app.StartApp()
	os.Exit(m.Run())
}
