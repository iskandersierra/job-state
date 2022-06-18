package models

import (
	"bytes"
	"encoding/json"
	"errors"
)

// JobStateStatus is the status of a job
type JobStateStatus int

const (
	JobUnknown   JobStateStatus = iota
	JobCreated                  = 1
	JobUpdated                  = 2
	JobCompleted                = 3
	JobFailed                   = 4

	MinJobStateStatus = JobCreated
	MaxJobStateStatus = JobFailed
)

// MarshalJSON implements json.Marshaler
func (status JobStateStatus) String() string {
	switch status {
	case JobCreated:
		return "created"
	case JobUpdated:
		return "updated"
	case JobCompleted:
		return "completed"
	case JobFailed:
		return "failed"
	default:
		return "unknown"
	}
}

// IsFinished returns true if the status is finished
func (status JobStateStatus) IsDefined() bool {
	return status >= MinJobStateStatus && status <= MaxJobStateStatus
}

// IsFinished returns true if the status is finished
func (status JobStateStatus) IsFinished() bool {
	switch status {
	case JobCompleted:
	case JobFailed:
		return true
	}
	return false
}

// MarshalJSON implements json.Marshaler
func (status JobStateStatus) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(status.String())
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON implements json.Unmarshaler
func (status *JobStateStatus) UnmarshalJSON(bytes []byte) error {
	var text string
	err := json.Unmarshal(bytes, &text)
	if err != nil {
		return err
	}
	switch text {
	case "created":
		*status = JobCreated
		return nil
	case "updated":
		*status = JobUpdated
		return nil
	case "completed":
		*status = JobCompleted
		return nil
	case "failed":
		*status = JobFailed
		return nil
	default:
		*status = JobUnknown
		return errors.New("unknown job state status")
	}
}
