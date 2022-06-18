package models

import (
	"time"
)

type JobState struct {
	JobId string `gorm:"primaryKey"`

	Title     string
	JobType   string
	StepCount int
	Metadata  string
	Steps     []JobStateStep `gorm:"foreignKey:FK__JobStateStep_Job"`

	Progress int
	Status   int
	Error    *string

	CreatedAt  time.Time
	UpdatedAt  time.Time
	FinishedAt *time.Time
}

type JobStateStep struct {
	JobId  string   `gorm:"primaryKey"`
	StepId int      `gorm:"primaryKey"`
	Job    JobState `gorm:"foreignKey:FK__JobStateStep_Job"`

	Title    string
	StepType string
	Progress int
	Status   int
	Metadata string
	Error    *string

	CreatedAt time.Time
}
