package models

import (
	"time"
)

type JobState struct {
	ID        string
	Title     string
	CreatedAt time.Time
}
