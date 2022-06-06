package grpc

import (
	"context"
	"time"

	"github.com/iskandersierra/job-state/pb"
	"github.com/iskandersierra/job-state/service"
)

func decodeCreateJobStateRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.CreateJobStateRequest)
	result := service.CreateJobState{
		Title: req.Title,
	}
	return result, nil
}

func encodeCreateJobStateResponse(_ context.Context, response interface{}) (interface{}, error) {
    res := response.(service.JobState)
    result := pb.CreateJobStateResponse{
        JobState: &pb.JobStateModel{
            Id: res.Id,
            Title: res.Title,
            CreatedAt: res.CreatedAt.UTC().Format(time.RFC3339),
        },
    }
    return result, nil
}
