package user_mysql_repository

import (
	"database/sql"
)

const tableName = "users"

type UserRepositoryMysql struct {
	DB *sql.DB
}

func (r *UserRepositoryMysql) GetTableName() string {
	return tableName
}

func (r *UserRepositoryMysql) GetDb() *sql.DB {
	return r.DB
}
