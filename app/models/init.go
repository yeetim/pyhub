package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Row interface {
	Scan(dest ...interface{}) error
}

func Initialize(dbPath string) error {
	if err := initConnection(dbPath); err != nil {
		return err
	}
	return nil
}

func initConnection(dbPath string) error {
	var err error
	db, err = sql.Open("mysql", dbPath)
	if err != nil {
		return err
	}
	//checkCmsSettings()
	return nil
}

func checkCmsSettings() {
	SetSettingIfNotExists("theme", "default", "cms")
	SetSettingIfNotExists("title", "The Go lang Cms", "cms")
	SetSettingIfNotExists("description", "The Go lang Cms", "cms")
}
