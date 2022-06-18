package repl

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	// "github.com/AlecAivazis/survey/v2"

	pb "github.com/iskandersierra/job-state/pb"
)

type ReplService struct {
	addr string
	conn *grpc.ClientConn
}

func New(addr string) (*ReplService, error) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to JobState gRPC service: " + addr)

	return &ReplService{
		addr: addr,
		conn: conn,
	}, nil
}

func (repl *ReplService) Start() error {
	client := pb.NewJobStateServiceClient(repl.conn)

	fmt.Println("Creating new job...")
	ctx := context.Background()
	response, err := client.CreateJobState(ctx, &pb.CreateJobStateRequest{
		Title: "New job",
		JobType: "create-client",
	})
	if err != nil {
		return err
	}

	jobId := response.JobId
	fmt.Println("Created JobState: " + jobId)

	return nil
}

func (repl *ReplService) Close() error {
	fmt.Println("Disconnecting from JobState gRPC service: " + repl.addr)
	defer repl.conn.Close()
	return nil
}
