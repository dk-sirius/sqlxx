package confpostgres

import (
	"fmt"

	"github.com/dk-sirius/sqlxx"
	"github.com/jmoiron/sqlx"
	"github.com/kelseyhightower/envconfig"
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
	// 环境变量存在则依据环境变量为主
	err = envconfig.Process(sqlxx.DB_TYPE_POSTGRES.String(), des)
	if err != nil {
		fmt.Println(err)
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
	// 环境变量存在则依据环境变量为主
	des := &executor.DBDescriptor
	err := envconfig.Process(sqlxx.DB_TYPE_POSTGRES.String(), des)
	if err != nil {
		fmt.Println(err)
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
