package models

import (
	"time"

	"github.com/iskandersierra/job-state/db/models"
)

type JobStateStep struct {
	JobId  string `json:"jobId"`
	StepId int    `json:"stepId"`

	Title    string            `json:"title"`
	StepType string            `json:"type"`
	Progress int               `json:"progress"`
	Status   JobStateStatus    `json:"status"`
	Metadata map[string]string `json:"metadata"`
	Error    map[string]string `json:"error"`

	CreatedAt time.Time `json:"createdAt"`
}

func FromJobStateStepModel(jobStateStep *models.JobStateStep) (*JobStateStep, error) {
	metadata, err := DeserializeMetadata(jobStateStep.Metadata)
	if err != nil {
		return nil, err
	}

	errorMetadata, err := DeserializeOptionalMetadata(jobStateStep.Error)
	if err != nil {
		return nil, err
	}

	return &JobStateStep{
		JobId:     jobStateStep.JobId,
		StepId:    jobStateStep.StepId,
		Title:     jobStateStep.Title,
		StepType:  jobStateStep.StepType,
		Metadata:  metadata,
		Progress:  jobStateStep.Progress,
		Status:    JobStateStatus(jobStateStep.Status),
		Error:     errorMetadata,
		CreatedAt: jobStateStep.CreatedAt,
	}, nil
}
