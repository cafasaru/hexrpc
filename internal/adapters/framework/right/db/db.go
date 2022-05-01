package db

import (
	"database/sql"
	db "hexrpc/internal/adapters/framework/right/db/sqlc"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type Adapter struct {
	db db.Store
}

func NewAdapter(driverName, dataSource string, log *logrus.Logger) (*Adapter, error) {
	conn, err := sql.Open(driverName, dataSource)
	if err != nil {
		log.WithFields(logrus.Fields{
			"driver: ":   driverName,
			"dataSource": dataSource,
		}).Error(err)
		return nil, err
	}

	if err := conn.Ping(); err != nil {
		return nil, err
	}

	store := db.NewStore(conn)

	return &Adapter{
		db: store,
	}, nil
}


func (dba Adapter) Ping() (string, error) {
	return "pong", nil
}