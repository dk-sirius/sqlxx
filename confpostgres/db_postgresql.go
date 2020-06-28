package confpostgres

import (
	"fmt"

	"github.com/dk-sirius/sqlxx"
	"github.com/jmoiron/sqlx"
)

type PostgresSql struct {
	sqlxx.DBDescriptor
}

func (executor *PostgresSql) ConnYAML(yaml string) *sqlx.DB {
	driverName := fmt.Sprintf("%v", sqlxx.DB_TYPE_POSTGRES)
	des, err := executor.ConfigByYaml(yaml)
	if err != nil {
		panic(err)
	}
	db, err := sqlx.Open(driverName, fmt.Sprintf("%v", des))
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(10)
	return db
}

func (executor *PostgresSql) Conn() *sqlx.DB {
	driverName := fmt.Sprintf("%v", sqlxx.DB_TYPE_POSTGRES)
	db, err := sqlx.Open(driverName, fmt.Sprintf("%v", executor.DBDescriptor))
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(10)
	return db
}
