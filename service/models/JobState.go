package models

import (
	"time"

	"github.com/iskandersierra/job-state/db/models"
)

// JobState is the struct that represents the job state.
type JobState struct {
	JobId string `json:"jobId"`

	Title     string            `json:"title"`
	JobType   string            `json:"jobType"`
	StepCount int               `json:"stepCount"`
	Metadata  map[string]string `json:"metadata"`
	Steps     []JobStateStep    `json:"steps"`

	Progress int               `json:"progress"`
	Status   JobStateStatus    `json:"status"`
	Error    map[string]string `json:"error"`

	CreatedAt  time.Time  `json:"createdAt"`
	UpdatedAt  time.Time  `json:"updatedAt"`
	FinishedAt *time.Time `json:"finishedAt"`
}

func FromJobStateModel(model *models.JobState) (*JobState, error) {
	metadata, err := DeserializeMetadata(model.Metadata)
	if err != nil {
		return nil, err
	}

	steps := []JobStateStep(nil)
	if model.Steps != nil {
		for _, step := range model.Steps {
			step, err := FromJobStateStepModel(&step)
			if err != nil {
				return nil, err
			}
			steps = append(steps, *step)
		}
	}

	return &JobState{
		JobId:      model.JobId,
		Title:      model.Title,
		JobType:    model.JobType,
		StepCount:  model.StepCount,
		Metadata:   metadata,
		Steps:      steps,
		CreatedAt:  model.CreatedAt,
		UpdatedAt:  model.UpdatedAt,
		FinishedAt: model.FinishedAt,
	}, nil
}
