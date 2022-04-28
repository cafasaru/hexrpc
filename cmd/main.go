package main

import (
	"fmt"
	"hexrpc/internal/core/ping"

	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()
	pingAdapter := ping.NewAdapter(logger)
	message, err := pingAdapter.Ping()
	if err != nil {
		logger.Error(err)
	}

	fmt.Println(message)
}
