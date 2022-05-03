package api

import (
	"hexrpc/internal/ports"

	"github.com/sirupsen/logrus"
)

type Adapter struct {
	db   ports.DbPort
	ping ports.PingPort
	log  *logrus.Logger
}

func NewAdapter(db ports.DbPort, ping ports.PingPort, log *logrus.Logger) *Adapter {
	return &Adapter{
		db:   db,
		ping: ping,
		log:  log,
	}
}

// Logic layer, unmarshal, marshal perform logic from ping port
// then call database to perform insert, ready update delete
func (apia Adapter) Ping() (string, error) {
	message, err := apia.ping.Ping()
	if err != nil {
		return message, err
	}

	apia.log.Info("INSIDE API ADAPTER")

	return message, nil

}
