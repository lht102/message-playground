package rest_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/google/uuid"
	"github.com/lht102/message-playground/jobworker"
	"github.com/lht102/message-playground/jobworker/api"
	"github.com/stretchr/testify/mock"
)

func (s *RESTTestSuite) TestGetJobHandler() {
	s.Run("Success", func() {
		jobUUID := uuid.MustParse("26013f20-4ede-46ef-bc08-693cc4a63ddb")
		completedAt := time.Date(2022, 1, 1, 2, 0, 0, 0, time.UTC)
		expectedResponse := api.JobResponse{
			UUID:        jobUUID,
			RequestUUID: uuid.New(),
			State:       api.JobStateCompleted,
			Description: "description",
			CompletedAt: &completedAt,
			CreatedAt:   time.Date(2022, 1, 1, 1, 0, 0, 0, time.UTC),
			UpdatedAt:   time.Date(2022, 1, 1, 2, 0, 0, 0, time.UTC),
		}
		s.jobServiceMock.
			EXPECT().
			GetJob(mock.Anything, jobUUID).
			Return(jobworker.Job{
				UUID:        jobUUID,
				RequestUUID: expectedResponse.RequestUUID,
				State:       jobworker.JobState(expectedResponse.State),
				Description: expectedResponse.Description,
				CompletedAt: &completedAt,
				CreatedAt:   expectedResponse.CreatedAt,
				UpdatedAt:   expectedResponse.UpdatedAt,
			}, nil)
		request := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/job/%s", jobUUID.String()), nil)
		response, b := testHTTPHandler(s.server, request)
		s.Equal(http.StatusOK, response.StatusCode)

		var actualResponse api.JobResponse
		err := json.Unmarshal(b, &actualResponse)
		s.NoError(err)
		s.Equal(expectedResponse, actualResponse)
	})

	s.Run("Some internal error with the query", func() {
		jobUUID := uuid.MustParse("46bfd468-574d-48d3-b7fb-4c28a02a00ce")
		s.jobServiceMock.
			EXPECT().
			GetJob(mock.Anything, jobUUID).
			Return(jobworker.Job{}, errors.New("some error"))
		request := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/job/%s", jobUUID.String()), nil)
		response, _ := testHTTPHandler(s.server, request)
		s.Equal(http.StatusInternalServerError, response.StatusCode)
	})

	s.Run("Invalid uuid format", func() {
		request := httptest.NewRequest(http.MethodGet, "/api/job/wrong-uuid", nil)
		response, _ := testHTTPHandler(s.server, request)
		s.Equal(http.StatusUnprocessableEntity, response.StatusCode)
	})
}
