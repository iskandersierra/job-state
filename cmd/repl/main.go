package main

import (
	"flag"
	"log"
	"os"

	"github.com/iskandersierra/job-state/repl"
	"github.com/subosito/gotenv"
)

var (
	addr = flag.String("addr", "", "JobState gRPC service address")
)

func main() {
	gotenv.Load(".env")

	flag.Parse()

	address := *addr
	if address == "" {
		address = os.Getenv("JOBSTATE_ADDR")
	}

	replService, err := repl.New(address)
	if err != nil {
		log.Fatalln(err)
	}

	err = replService.Start()
	if err != nil {
		log.Fatalln(err)
	}
	defer replService.Close()
}
