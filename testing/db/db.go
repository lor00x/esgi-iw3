package db

import "salut/testing/model"

type Db interface {
	Query(string, int) (*model.User, error)
}
