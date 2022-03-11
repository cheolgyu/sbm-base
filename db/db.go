package db

import (
	"database/sql"

	"github.com/cheolgyu/sbm-base/c"
	"github.com/cheolgyu/sbm-base/env"
	"github.com/cheolgyu/sbm-base/logging"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var Conn *sqlx.DB

type PQ struct {
}

var url string

func init() {
	url = env.MyEnv["DB_URL"]
	pg := PQ{}
	Conn = pg.conn_sqlx()
	Conn.SetMaxIdleConns(c.DB_MAX_CONN)
	Conn.SetMaxOpenConns(c.DB_MAX_CONN)
	logging.Log.Debug("db init 실행: ", url)
}

func (o *PQ) conn() *sql.DB {

	db, err := sql.Open("postgres", url)
	if err != nil {
		logging.Log.Error("커넥션 오류:", err)
		panic(err)
	}
	return db
}

func (o *PQ) conn_sqlx() *sqlx.DB {

	db, err := sqlx.Connect("postgres", url)
	if err != nil {
		logging.Log.Error("커넥션 오류222:", url, err, url)
		panic(err)
	}
	return db
}

func Connect() *sqlx.DB {
	pg := PQ{}
	return pg.conn_sqlx()
}
