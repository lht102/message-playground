package ent

import "github.com/lht102/message-playground/jobworker"

func ParseJobFromModel(job *Job) jobworker.Job {
	return jobworker.Job{
		UUID:        job.ID,
		RequestUUID: job.RequestUUID,
		State:       job.State,
		Description: job.Description,
		CompletedAt: job.CompletedAt,
		CreatedAt:   job.CreatedAt,
		UpdatedAt:   job.UpdatedAt,
	}
}
