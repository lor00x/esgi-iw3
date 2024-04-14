package testing

import (
	"go.uber.org/mock/gomock"
	"salut/testing/db/mock"
	"salut/testing/model"
	"testing"
)

func TestGetUserById(t *testing.T) {

	ctrl := gomock.NewController(t)
	db := mock.NewMockDb(ctrl)
	db.EXPECT().
		Query("SELECT * FROM users WHERE id = ?", 18).
		Return(&model.User{}, nil)

	GetUserById(db, 18)

}
