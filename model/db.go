package model

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var Mgr *Manager

func init() {
	db, err := sql.Open("mysql", "root:root@/test1")
	CheckErr(err)
	err = db.Ping()
	CheckErr(err)
	Mgr = &Manager{db: db}
}

type Manager struct {
	db *sql.DB
}

func CheckErr(err error) {
	if err != nil {
		Mgr.db.Close()
		panic(err)
	}
}
