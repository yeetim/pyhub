package models

import (
	"fmt"
	"time"

	"github.com/eirini/test/app/helpers"
	"github.com/russross/meddler"
)

const stmtGetSetting = `SELECT * FROM settings WHERE key = ?`
const stmtSaveSelect = `SELECT id FROM settings WHERE KEY = ?`

type Setting struct {
	Id        int        `meddler:"id,pk"`
	Key       string     `meddler:"key"`
	Value     string     `meddler:"value"`
	Type      string     `meddler:"type"`
	CreatedAt *time.Time `meddler:"created_at"`
	CreatedBy int64      `meddler:"created_by"`
	UpdatedAt *time.Time `meddler:"updated_at"`
	UpdatedBy int64      `meddler:"updated_by"`
}

func (m *Setting) GetSetting() error {
	err := meddler.QueryRow(db, m, stmtGetSetting, m.Key)
	return err
}

func (m *Setting) Save() error {
	var id int
	row := db.QueryRow(stmtGetSetting, m.Key)
	if err := row.Scan(&id); err != nil {
		m.Id = 0
	} else {
		m.Id = id
	}
	err := meddler.Save(db, "settings", m)
	return err
}

func NewSetting(k, v, t string) *Setting {
	return &Setting{
		Key:       k,
		Value:     v,
		Type:      t,
		CreatedAt: helpers.Now(),
	}
}

func SetSettingIfNotExists(k, v, t string) error {
	s := NewSetting(k, v, t)
	err := s.GetSetting()
	if err != nil {
		fmt.Println("save")
		return s.Save()
	}
	return err
}
