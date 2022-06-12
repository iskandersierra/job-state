package main

import (
	"os"

	"github.com/iskandersierra/job-state/api"
)

func main() {
	err := api.StartApi()
	if err != nil {
		os.Exit(1)
	}
}
