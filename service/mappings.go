package service

import "github.com/iskandersierra/job-state/db/models"

func fromJobStateModel(model *models.JobState) JobState {
	return JobState{
		Id:        model.ID,
		Title:     model.Title,
		CreatedAt: model.CreatedAt,
	}
}
