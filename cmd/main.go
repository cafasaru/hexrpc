package main

import (
	"hexrpc/internal/adapters/app/api"
	"hexrpc/internal/adapters/core/ping"
	gRPC "hexrpc/internal/adapters/framework/left/grpc"
	"hexrpc/internal/adapters/framework/right/db"
	"hexrpc/internal/ports"
	"hexrpc/util"
	"log"

	"github.com/sirupsen/logrus"
)

// config file
var config util.Config

func main() {
	logger := logrus.New()
	var err error

	// ports
	var core ports.PingPort
	var dbaseAdapter ports.DbPort
	var appAdapter ports.APIPort
	var gRPCAdapter ports.GRPCPort

	dbaseAdapter, err = db.NewAdapter(config.DBDriver, config.DBSource, logger)
	if err != nil {
		logger.Fatal(err)
	}

	core = ping.NewAdapter(logger)

	appAdapter = api.NewAdapter(dbaseAdapter, core, logger)

	gRPCAdapter = gRPC.NewAdapter(appAdapter)

	gRPCAdapter.Run()

}

// init loads config file
func init() {
	conf, err := util.LoadConfig("../")
	if err != nil {
		log.Fatal(err)
	}
	config = conf
}
