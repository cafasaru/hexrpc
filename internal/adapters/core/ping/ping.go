package ping

import (
	"github.com/sirupsen/logrus"
)

type Adapter struct {
	log *logrus.Logger
}

func NewAdapter(log *logrus.Logger) *Adapter {
	return &Adapter{log: log}
}
func (ping Adapter) Ping() (string, error) {
	return "pong", nil
}
