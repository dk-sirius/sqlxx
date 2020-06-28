package sqlxx

import "github.com/jmoiron/sqlx"

type DbHandler interface {
	Conn() *sqlx.DB
}
