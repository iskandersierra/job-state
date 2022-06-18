package service

import (
	"math"
	"time"

	"gorm.io/gorm"

	"github.com/go-playground/validator/v10"

	db "github.com/iskandersierra/job-state/db/models"
	models "github.com/iskandersierra/job-state/service/models"
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
func (service *jobStateServiceImpl) CreateJobState(command *models.CreateJobState) (*models.CreateJobStateResult, error) {
	err := validate.Struct(command)
	if err != nil {
		return nil, err
	}

	jobId := utils.NewId()
	now := time.Now()

	dao, err := jobStateDaoFromCreateCommand(command, jobId, now)
	if err != nil {
		return nil, err
	}

	created := service.db.Create(&dao)
	err = created.Error
	if err != nil {
		return nil, err
	}

	return &models.CreateJobStateResult{
		JobId: jobId,
	}, nil
}

func jobStateDaoFromCreateCommand(command *models.CreateJobState, jobId string, now time.Time) (*db.JobState, error) {
	var status models.JobStateStatus = models.JobCreated
	if command.Status.IsDefined() {
		status = command.Status
	}

	var progress = int(math.Max(0, math.Min(100, float64(command.Progress))))
	if status.IsFinished() {
		progress = 100
	}

	jobMetadata, err := models.SerializeMetadata(command.Metadata)
	if err != nil {
		return nil, err
	}

	stepMetadata, err := models.SerializeMetadata(command.StepMetadata)
	if err != nil {
		return nil, err
	}

	errorMetadata, err := models.SerializeOptionalMetadata(command.Error)
	if err != nil {
		return nil, err
	}

	stepDao := db.JobStateStep{
		JobId:  jobId,
		StepId: 1,

		Title:    command.StepTitle,
		StepType: command.StepType,
		Progress: progress,
		Status:   int(status),
		Metadata: stepMetadata,
		Error:    errorMetadata,

		CreatedAt: now,
	}

	steps := []db.JobStateStep{stepDao}

	dao := db.JobState{
		JobId: jobId,

		Title:     command.Title,
		JobType:   command.JobType,
		StepCount: len(steps),
		Metadata:  jobMetadata,
		Steps:     steps,

		Progress: stepDao.Progress,
		Status:   stepDao.Status,
		Error:    stepDao.Error,

		CreatedAt: now,
		UpdatedAt: now,
	}

	return &dao, nil
}
