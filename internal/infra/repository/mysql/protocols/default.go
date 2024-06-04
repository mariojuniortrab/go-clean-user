package protocols_mysql_repository

import (
	"database/sql"
)

type DefaultRepositoryProtocol interface {
	GetDb() *sql.DB
	GetTableName() string
}

type DefaultSelectRepositoryProtocol interface {
	DefaultRepositoryProtocol
}
