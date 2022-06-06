package grpc

import (
	"context"

	"github.com/go-kit/kit/log"
	grpcTransport "github.com/go-kit/kit/transport/grpc"

	"github.com/iskandersierra/job-state/endpoints"
	"github.com/iskandersierra/job-state/pb"
)

// jobStateGRPCServer is a gRPC server that decodes requests to JobStateService
type jobStateGRPCServer struct {
	pb.UnimplementedJobStateServiceServer
	createJobState grpcTransport.Handler
}

func NewJobStateGRPCServer(endpoints endpoints.JobStateEndpoints, logger log.Logger) *jobStateGRPCServer {
	options := []grpcTransport.ServerOption{
		grpcTransport.ServerErrorLogger(logger),
	}

	return &jobStateGRPCServer{
		createJobState: grpcTransport.NewServer(
			endpoints.CreateJobState,
			decodeCreateJobStateRequest,
			encodeCreateJobStateResponse,
			options...,
		),
	}
}

// Force implementation of interface
var _ pb.JobStateServiceServer = &jobStateGRPCServer{}

// CreateJobState implements pb.JobStateServiceServer
func (server *jobStateGRPCServer) CreateJobState(ctx context.Context, request *pb.CreateJobStateRequest) (*pb.CreateJobStateResponse, error) {
	_, resp, err := server.createJobState.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}

	return resp.(*pb.CreateJobStateResponse), nil
}
