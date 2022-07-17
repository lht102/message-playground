package rest_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/lht102/message-playground/jobworker/http/rest"
	"github.com/lht102/message-playground/jobworker/mocks"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type RESTTestSuite struct {
	suite.Suite

	server         *rest.Server
	jobServiceMock *mocks.Service
}

func TestRest(t *testing.T) {
	logger, err := zap.NewDevelopment()
	require.NoError(t, err)

	jobServiceMock := mocks.NewService(t)
	server := rest.NewServer(8081, jobServiceMock, logger)

	s := &RESTTestSuite{
		server:         server,
		jobServiceMock: jobServiceMock,
	}

	suite.Run(t, s)
}

func testHTTPHandler(handler http.Handler, r *http.Request) (*http.Response, []byte) {
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, r)

	resp := w.Result()
	b, _ := io.ReadAll(resp.Body)

	return resp, b
}
