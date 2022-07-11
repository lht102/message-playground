package rest

import (
	"net/http"

	jobworker "github.com/lht102/message-playground/jobworker"
	"github.com/lht102/message-playground/jobworker/api"
)

func createJobHandler(
	jobService jobworker.Service,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var req api.CreateJobRequest
		if err := decode(r, &req); err != nil {
			respondErr(w, http.StatusUnprocessableEntity, err.Error())

			return
		}

		createdJob, err := jobService.CreateJob(ctx, jobworker.CreateJobCommand{
			RequestUUID: req.RequestUUID,
			Description: req.Description,
		})
		if err != nil {
			respondErr(w, http.StatusInternalServerError, err.Error())

			return
		}

		respond(w, http.StatusCreated, api.CreateJobResponse{
			JobUUID: createdJob.UUID,
		})
	}
}
