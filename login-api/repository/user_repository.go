package repository

import (
	"context"
	"database/sql"
	"tugas-pemrograman-web/model"
)

type UserRepository interface {
	CreateUser(ctx context.Context, tx *sql.Tx, user model.User) (model.User, error)
	ReadUser(ctx context.Context, tx *sql.Tx) []model.User
	UpdateUser(ctx context.Context, tx *sql.Tx, user model.User) model.User
	FindById(ctx context.Context, tx *sql.Tx, idUser string) (model.User, error)
	FindByEmail(ctx context.Context, tx *sql.Tx, email string) (model.User, error)
	DeleteUser(ctx context.Context, tx *sql.Tx, user model.User) model.User
}
