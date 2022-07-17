package rest_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"

	"github.com/google/uuid"
	"github.com/lht102/message-playground/jobworker"
	"github.com/lht102/message-playground/jobworker/api"
	"github.com/stretchr/testify/mock"
)

func (s *RESTTestSuite) TestCreateJobHandler() {
	s.Run("Success", func() {
		createJobRequest := api.CreateJobRequest{
			RequestUUID: uuid.MustParse("baa1e965-2881-4f39-acb5-8316bf79b457"),
			Description: "todo",
		}
		s.jobServiceMock.
			EXPECT().
			CreateJob(mock.Anything, jobworker.CreateJobCommand{
				RequestUUID: createJobRequest.RequestUUID,
				Description: createJobRequest.Description,
			}).
			Return(jobworker.Job{
				UUID: createJobRequest.RequestUUID,
			}, nil)
		requestBody, err := json.Marshal(createJobRequest)
		s.Require().NoError(err)
		request := httptest.NewRequest(http.MethodPost, "/api/job", bytes.NewBuffer(requestBody))
		response, b := testHTTPHandler(s.server, request)
		s.Equal(http.StatusCreated, response.StatusCode)

		var actualResponse api.CreateJobResponse
		err = json.Unmarshal(b, &actualResponse)
		s.NoError(err)
		s.Equal(api.CreateJobResponse{
			JobUUID: createJobRequest.RequestUUID,
		}, actualResponse)
	})

	s.Run("Some internal error with the creation", func() {
		createJobRequest := api.CreateJobRequest{
			RequestUUID: uuid.MustParse("987f2081-e2d0-4853-a358-a64af591b339"),
			Description: "todo",
		}
		s.jobServiceMock.
			EXPECT().
			CreateJob(mock.Anything, jobworker.CreateJobCommand{
				RequestUUID: createJobRequest.RequestUUID,
				Description: createJobRequest.Description,
			}).
			Return(jobworker.Job{}, errors.New("some error"))
		requestBody, err := json.Marshal(createJobRequest)
		s.Require().NoError(err)
		request := httptest.NewRequest(http.MethodPost, "/api/job", bytes.NewBuffer(requestBody))
		response, _ := testHTTPHandler(s.server, request)
		s.Equal(http.StatusInternalServerError, response.StatusCode)
	})

	s.Run("Invalid request body format", func() {
		request := httptest.NewRequest(http.MethodPost, "/api/job", bytes.NewBuffer([]byte("some bytes")))
		response, _ := testHTTPHandler(s.server, request)
		s.Equal(http.StatusUnprocessableEntity, response.StatusCode)
	})
}
