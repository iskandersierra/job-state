package service

import "time"

// JobState is the struct that represents the job state.
type JobState struct {
    Id string `json:"id"`
    Title string `json:"title"`
    CreatedAt time.Time `json:"created_at"`
}

// CreateJobState is the struct that represents the command to create a job state.
type CreateJobState struct {
    Title string `json:"title"`
}
