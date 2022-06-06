package service

import (
	"time"

	"github.com/gofrs/uuid"
)

// jobStateServiceImpl is the default implementation of the JobStateService interface.
type jobStateServiceImpl struct {
}

// NewJobStateService returns a new instance of the default JobStateService implementation.
func NewJobStateService() JobStateService {
    return &jobStateServiceImpl{}
}

// CreateJobState creates a new job state.
func (service *jobStateServiceImpl) CreateJobState(command CreateJobState) (JobState, error) {
    id, err := uuid.NewV4()
    if err != nil {
        return JobState{}, err
    }

    result := JobState{
        Id: id.String(),
        Title: command.Title,
        CreatedAt: time.Now(),
    }

    return result, nil
}
