package service

// JobStateService describes the service operations.
type JobStateService interface {
    CreateJobState(command CreateJobState) (JobState, error)
}
