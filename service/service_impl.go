package service

import (
	"time"

	"gorm.io/gorm"

	"github.com/go-playground/validator/v10"

	"github.com/iskandersierra/job-state/db/models"
	"github.com/iskandersierra/job-state/utils"
)

var validate = validator.New()

// jobStateServiceImpl is the default implementation of the JobStateService interface.
type jobStateServiceImpl struct {
	db *gorm.DB
}

// NewJobStateService returns a new instance of the default JobStateService implementation.
func NewJobStateService(db *gorm.DB) JobStateService {
	return &jobStateServiceImpl{
		db: db,
	}
}

// CreateJobState creates a new job state.
func (service *jobStateServiceImpl) CreateJobState(command *CreateJobState) (*JobState, error) {
	err := validate.Struct(command)
	if err != nil {
		return nil, err
	}

	id := utils.NewId()

	model := models.JobState{
		ID:        id,
		Title:     command.Title,
		CreatedAt: time.Now(),
	}

	created := service.db.Create(&model)

	if created.Error != nil {
		return nil, created.Error
	}

	result := fromJobStateModel(&model)

	return &result, nil
}
