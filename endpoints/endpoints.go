package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"github.com/iskandersierra/job-state/service"
)

// JobStateEndpoints wraps all of the service's endpoints.
type JobStateEndpoints struct {
    CreateJobState endpoint.Endpoint
}

// NewJobStateEndpoints returns an initialized JobStateEndpoints struct.
func NewJobStateEndpoints(svc service.JobStateService) JobStateEndpoints {
    return JobStateEndpoints{
        CreateJobState: makeCreateJobStateEndpoint(svc),
    }
}

// makeCreateJobStateEndpoint returns an endpoint that invokes CreateJobState on the service.
func makeCreateJobStateEndpoint(svc service.JobStateService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
        req := request.(service.CreateJobState)

        result, err := svc.CreateJobState(req)
        if err != nil {
            return nil, err
        }

        return result, nil
    }
}

