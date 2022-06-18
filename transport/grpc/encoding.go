package grpc

import (
	"context"

	"github.com/iskandersierra/job-state/pb"
	"github.com/iskandersierra/job-state/service/models"
)

func decodeCreateJobStateRequest(_ context.Context, request interface{}) (interface{}, error) {
	request, err := fromProtoCreateJobStateRequest(request.(*pb.CreateJobStateRequest))
	if err != nil {
		return nil, err
	}
	return request, nil
}

func fromProtoCreateJobStateRequest(req *pb.CreateJobStateRequest) (*models.CreateJobState, error) {
	stepError := map[string]string(nil)
	if req.StepError != nil {
		stepError = req.StepError.Metadata
	}

	return &models.CreateJobState{
		Title:    req.Title,
		JobType:  req.JobType,
		Metadata: req.Metadata,

		StepTitle:    req.StepTitle,
		StepType:     req.StepType,
		Progress:     int(req.Progress),
		Status:       models.JobStateStatus(req.Status),
		StepMetadata: req.StepMetadata,
		Error:        stepError,
	}, nil
}

func encodeCreateJobStateResponse(_ context.Context, response interface{}) (interface{}, error) {
	response, err := toProtoCreateJobStateResponse(response.(*models.CreateJobStateResult))
	if err != nil {
		return nil, err
	}
	return response, nil
}

func toProtoCreateJobStateResponse(res *models.CreateJobStateResult) (*pb.CreateJobStateResponse, error) {
	return &pb.CreateJobStateResponse{
		JobId: res.JobId,
	}, nil
}
