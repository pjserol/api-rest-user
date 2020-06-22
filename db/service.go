package db

import (
	"context"
	"database/sql"
	"errors"

	"github.com/pjserol/api-rest-user/common/logs"
)

//ServiceDB service of the DB package
type ServiceDB struct {
	DB *sql.DB
}

// NewServiceDB return the connection of the DB
func NewServiceDB(ctx context.Context, appEnv string, db *sql.DB) (ServiceDB, error) {
	if !Connected {
		logs.Log(ctx, "database not connected")
		return ServiceDB{}, errors.New("database not connected")
	}

	database := ServiceDB{
		DB: db,
	}

	return database, nil
}
