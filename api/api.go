package api

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	"github.com/iskandersierra/job-state/endpoints"
	pb "github.com/iskandersierra/job-state/pb"
	"github.com/iskandersierra/job-state/service"
	transport "github.com/iskandersierra/job-state/transport/grpc"
)

func StartApi() error {
	logger := createLogger()

	db, err := connectDatabase()
	if err != nil {
		return err
	}

	jobStateServer := createJobStateServer(db, logger)

	errorsChannel := createErrorsChannel()

	err = startServer(logger, jobStateServer)
	if err != nil {
		level.Error(logger).Log("during", "startServer", "err", err)
		return err
	}

	level.Error(logger).Log("exit", <-errorsChannel)

	return nil
}

func createLogger() log.Logger {
	logger := log.NewJSONLogger(os.Stdout)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	return logger
}

type nameReplacerNoop struct{}

// Replace implements schema.Replacer
func (nameReplacerNoop) Replace(name string) string {
	return name
}

func connectDatabase() (*gorm.DB, error) {
	connectionString := os.Getenv("SQLSERVER_CONNECTION_STRING")

	sqlConn := sqlserver.Open(connectionString)

	config := gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "jobs.",
			NameReplacer:  nameReplacerNoop{},
			SingularTable: true,
			NoLowerCase:   true,
		},
	}

	return gorm.Open(sqlConn, &config)
}

func createJobStateServer(
	db *gorm.DB,
	logger log.Logger) pb.JobStateServiceServer {
	svc := service.NewJobStateService(db)
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

	addr := os.Getenv("JOBSTATE_ADDR")

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	fmt.Printf("Listening on port %s\n", addr)

	go func() {
		baseServer := grpc.NewServer()
		pb.RegisterJobStateServiceServer(baseServer, jobStateServer)
		level.Info(logger).Log("msg", "Server started successfully 🚀")
		baseServer.Serve(listener)
	}()

	return nil
}
