package model

import (
	"go.uber.org/mock/gomock"
	"salut/testing/db/mock"
	"testing"
)

func TestGetUserById(t *testing.T) {

	ctrl := gomock.NewController(t)
	db := mock.NewMockDb(ctrl)
	db.EXPECT().
		Query("SELECT * FROM users WHERE id = ?", 18).
		Return(&User{}, nil)

	GetUserById(db, 18)

}
