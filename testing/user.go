package testing

import (
	"salut/testing/db"
	"salut/testing/model"
)

func GetUserById(db db.Db, id int) (*model.User, error) {
	return db.Query("SELECT * FROM users WHERE id = ?", id)
}
