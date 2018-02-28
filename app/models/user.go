package models

import (
	"time"

	"github.com/russross/meddler"
)

const stmtGetUserByEmail = `SELECT * FROM users WHERE email = ?`
const stmtGetUsersCountByEmail = `SELECT count(*) FROM users where email = ?`
const stmtGetNumberOfUsers = `SELECT COUNT(*) FROM users`

type User struct {
	Id             int64      `meddler:"id,pk"`
	Name           string     `meddler:"name"`
	Slug           string     `meddler:"slug"`
	HashedPassword string     `meddler:"password"`
	Email          string     `meddler:"email"`
	Image          string     `meddler:"image"`
	Cover          string     `meddler:"cover"`
	Bio            string     `meddler:"bio"`
	Website        string     `meddler:"website"`
	Location       string     `meddler:"location"`
	Accessibility  string     `meddler:"accessibility"`
	Status         string     `meddler:"status"`
	Language       string     `meddler:"language"`
	Lastlogin      *time.Time `meddler:"last_login"`
	CreatedAt      *time.Time `meddler:"created_at"`
	CreatedBy      int        `meddler:"created_by"`
	UpdatedAt      *time.Time `meddler:"updated_at"`
	UpdatedBy      int        `meddler:"updated_by"`
	Role           int        `meddler:"-"`
}

func (m *User) GetUserByEmail() error {
	err := meddler.QueryRow(db, m, stmtGetUserByEmail, m.Email)
	return err
}

func GetNumberOfUsers() (int64, error) {
	var count int64
	row := db.QueryRow(stmtGetNumberOfUsers)
	err := row.Scan(&count)
	return count, err
}
