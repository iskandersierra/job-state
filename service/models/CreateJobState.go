package models

// CreateJobState is the struct that represents the command to create a job state.
type CreateJobState struct {
	Title    string            `json:"title" validate:"required,min=1,max=100"`
	JobType  string            `json:"jobType" validate:"required,min=1,max=100"`
	Metadata map[string]string `json:"metadata"`

	StepTitle    string            `json:"stepTitle" validate:"min=0,max=100"`
	StepType     string            `json:"stepType" validate:"min=0,max=100"`
	Progress     int               `json:"progress" validate:"min=0,max=100"`
	Status       JobStateStatus    `json:"status" validate:"min=0,max=4"`
	StepMetadata map[string]string `json:"stepMetadata"`
	Error        map[string]string `json:"error"`
}

type CreateJobStateResult struct {
	JobId string `json:"jobId"`
}
