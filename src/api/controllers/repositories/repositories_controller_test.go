package repositories

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Washington-Ksea/golang-microservices/src/api/utils/errors"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateRepoInvalidJsonRequest(t *testing.T) {
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)

	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(``))
	c.Request = request

	CreateRepo(c)

	assert.EqualValues(t, http.StatusBadRequest, response.Code)
	apiErr, err := errors.NewApiErrFromBytes(response.Body.Bytes())
	assert.Nil(t, err)
	assert.NotNil(t, apiErr)
}
