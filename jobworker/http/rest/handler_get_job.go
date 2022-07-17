package rest

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/lht102/message-playground/jobworker"
	"github.com/lht102/message-playground/jobworker/job"
	"github.com/uptrace/bunrouter"
)

func getJobHandler(
	jobService jobworker.JobService,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		params := bunrouter.ParamsFromContext(ctx)
		uuidStr := params.ByName("uuid")

		jobUUID, err := uuid.Parse(uuidStr)
		if err != nil {
			respondErr(w, http.StatusUnprocessableEntity, err.Error())

			return
		}

		queriedJob, err := jobService.GetJob(ctx, jobUUID)
		if errors.Is(err, job.ErrNotFound) {
			respondErr(w, http.StatusNotFound, err.Error())

			return
		}

		if err != nil {
			respondErr(w, http.StatusInternalServerError, err.Error())

			return
		}

		respond(w, http.StatusOK, jobworker.ParseJobAPIResponse(queriedJob))
	}
}
