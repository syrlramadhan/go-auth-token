package repository

import (
	"context"
	"database/sql"
	"errors"
	"tugas-pemrograman-web/model"
	"tugas-pemrograman-web/util"
)

type userRepositoryImpl struct {
}

func NewUserRepositoryImpl() UserRepository {
	return &userRepositoryImpl{}
}

func (repository *userRepositoryImpl) CreateUser(ctx context.Context, tx *sql.Tx, user model.User) (model.User, error) {
	query := `INSERT INTO register(id, name, email, password) VALUES(?, ?, ?, ?)`

	_, err := tx.ExecContext(ctx, query, user.Id, user.Name, user.Email, user.Password)
	util.SentPanicIfError(err)

	return user, nil
}

func (repository *userRepositoryImpl) ReadUser(ctx context.Context, tx *sql.Tx) []model.User {
	query := `SELECT id, name, email, password, photo FROM register`

	rows, err := tx.QueryContext(ctx, query)
	util.SentPanicIfError(err)
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		user := model.User{}
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Photo)
		util.SentPanicIfError(err)
		users = append(users, user)
	}

	return users
}

func (repository *userRepositoryImpl) UpdateUser(ctx context.Context, tx *sql.Tx, user model.User) model.User {
	query := `UPDATE register SET name = ?, email = ? WHERE id = ?`

	_, err := tx.ExecContext(ctx, query, user.Name, user.Email, user.Id)
	util.SentPanicIfError(err)

	return user
}

func (repository *userRepositoryImpl) UpdatePhoto(ctx context.Context, tx *sql.Tx, user model.User) model.User {
	query := `UPDATE register SET photo = ? WHERE id = ?`

	_, err := tx.ExecContext(ctx, query, user.Photo, user.Id)
	util.SentPanicIfError(err)

	return user
}

func (repository *userRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, idUser string) (model.User, error) {
	query := `SELECT id, name, email, password, photo FROM register WHERE id = ?`

	rows, err := tx.QueryContext(ctx, query, idUser)
	util.SentPanicIfError(err)

	defer rows.Close()
	users := model.User{}
	if rows.Next() {
		err := rows.Scan(&users.Id, &users.Name, &users.Email, &users.Password, &users.Photo)
		util.SentPanicIfError(err)
		return users, err
	} else {
		return users, errors.New("id user not found")
	}
}

func (repository *userRepositoryImpl) FindByEmail(ctx context.Context, tx *sql.Tx, email string) (model.User, error) {
	query := `SELECT id, name, email, password, photo FROM register WHERE email = ?`

	rows, err := tx.QueryContext(ctx, query, email)
	util.SentPanicIfError(err)

	defer rows.Close()
	users := model.User{}
	if rows.Next() {
		err := rows.Scan(&users.Id, &users.Name, &users.Email, &users.Password, &users.Photo)
		util.SentPanicIfError(err)
		return users, err
	} else {
		return users, errors.New("id user not found")
	}
}

func (repository *userRepositoryImpl) DeleteUser(ctx context.Context, tx *sql.Tx, user model.User) model.User {
	query := `DELETE FROM register WHERE id = ?`

	_, err := tx.ExecContext(ctx, query, user.Id)
	util.SentPanicIfError(err)

	return user
}
