package main

import (
	"os"

	"github.com/iskandersierra/job-state/api"
	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load(".env")

	err := api.StartApi()
	if err != nil {
		os.Exit(1)
	}
}
