package persist

import (
	"errors"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type sqlDB struct {
	*sqlx.DB
}

type dataAccessObject struct {
	sqlDB
}

type Persist interface {
	ServerInfo() (string, error)
}

func Open(args string) (Persist, error) {
	var err error
	conn := &sqlx.DB{}
	if err != nil {
		return dataAccessObject{}, errors.New("Unable to open")
	}
	return dataAccessObject{sqlDB{conn}}, nil
}

func (dao dataAccessObject) ServerInfo() (string, error) {
	return "info", nil
}
