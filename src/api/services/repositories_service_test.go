package services

import (
	"os"
	"testing"

	"github.com/Washington-Ksea/golang-microservices/src/api/clients/restclient"
)

func TestMain(m *testing.M) {
	restclient.StartMockups()
	os.Exit(m.Run())
}
