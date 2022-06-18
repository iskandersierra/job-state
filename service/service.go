package service

import "github.com/iskandersierra/job-state/service/models"

// JobStateService describes the service operations.
type JobStateService interface {
	CreateJobState(command *models.CreateJobState) (*models.CreateJobStateResult, error)
}
