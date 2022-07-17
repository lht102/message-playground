package rest

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/lht102/message-playground/jobworker"
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

		job, err := jobService.GetJob(ctx, jobUUID)
		if err != nil {
			respondErr(w, http.StatusInternalServerError, err.Error())

			return
		}

		respond(w, http.StatusOK, jobworker.ParseJobAPIResponse(job))
	}
}
