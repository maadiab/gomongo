package database

import (
	"context"

	"github.com/maadiab/gomongo/core"
)

type UserDB interface {
	AddUser(context.Context, core.User) (core.User, error)
}

const addUser = `INSERT INTO users (id, email) VALUES ($1, $2);`

func (d *database) AddUser(ctx context.Context, user core.User) (core.User, error) {
	err := d.db.QueryRowContext(
		ctx,
		addUser,
		user.ID,
		user.Email,
	).Scan(&user)
	return user, err
}
