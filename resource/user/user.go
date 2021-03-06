package user

import (
	"github.com/albertwidi/kothak/service/user"
	"github.com/jmoiron/sqlx"
)

type Resource struct {
	masterDB   *sqlx.DB
	followerDB *sqlx.DB
}

func New(masterDB, followerDB *sqlx.DB) *Resource {
	r := Resource{
		masterDB:   masterDB,
		followerDB: followerDB,
	}
	return &r
}

func (r *Resource) GetUser(userid int64) (user.User, error) {
	u := user.User{}
	return u, nil
}
