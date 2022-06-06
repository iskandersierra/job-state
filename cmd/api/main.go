package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	"github.com/iskandersierra/job-state/endpoints"
	"github.com/iskandersierra/job-state/pb"
	"github.com/iskandersierra/job-state/service"
	transport "github.com/iskandersierra/job-state/transport/grpc"
)

func main() {
	logger := createLogger()

	jobStateServer := createJobStateServer(logger)

	errorsChannel := createErrorsChannel()
	defer level.Error(logger).Log("exit", <-errorsChannel)

	err := startServer(logger, jobStateServer)
	if err != nil {
		level.Error(logger).Log("during", "startServer", "err", err)
		os.Exit(1)
	}
}

func createLogger() log.Logger {
	logger := log.NewJSONLogger(os.Stdout)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	return logger
}

func createJobStateServer(logger log.Logger) pb.JobStateServiceServer {
	svc := service.NewJobStateService()
	endpoint := endpoints.NewJobStateEndpoints(svc)
	server := transport.NewJobStateGRPCServer(endpoint, logger)
	return server
}

func createErrorsChannel() <-chan error {
	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	return errs
}

func startServer(
	logger log.Logger,
	jobStateServer pb.JobStateServiceServer) error {

	listener, err := net.Listen("tcp", ":51051")
	if err != nil {
		return err
	}

	go func() {
		baseServer := grpc.NewServer()
		pb.RegisterJobStateServiceServer(baseServer, jobStateServer)
		level.Info(logger).Log("msg", "Server started successfully 🚀")
		baseServer.Serve(listener)
	}()

	return nil
}
