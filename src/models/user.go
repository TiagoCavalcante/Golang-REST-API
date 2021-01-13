package models

import (
	"database/sql"

	"entities"
)

type UserModel struct {
	Database *sql.DB
}

func (userModel UserModel) FindAllUsers() ([]entities.User, error) {
	rows, err := userModel.Database.Query(
		"SELECT * FROM users",
	)

	if err != nil {
		return nil, err
	}
		
	users := []entities.User{}

	for rows.Next() {
		var id int64
		var name string
		var email string

		err = rows.Scan(&id, &name, &email)

		if (err != nil) {
			return nil, err
		}

		users = append(
			users,
			entities.User{id, name, email},
		)
	}

	return users, nil
}

func (userModel UserModel) FindUserById(id int64) (entities.User, error) {
	rows, err := userModel.Database.Query(
		"SELECT name, email FROM users WHERE id = ?",
		id,
	)

	if err != nil {
		return entities.User{}, err
	}
		
	for rows.Next() {
		var name string
		var email string

		err = rows.Scan(&name, &email)

		if (err != nil) {
			return entities.User{}, err
		}
		
		return entities.User{id, name, email}, nil
	}

	return entities.User{}, nil
}

func (userModel UserModel) CreateUser(user *entities.User) error {
	result, err := userModel.Database.Exec(
		"INSERT INTO users (name, email) VALUES (?, ?)",
		user.Name,
		user.Email,
	)

	if err != nil {
		return err
	}

	user.Id, _ = result.LastInsertId()

	return nil
}

func (userModel UserModel) UpdateUser(user *entities.User) (int64, error) {
	result, err := userModel.Database.Exec(
		"UPDATE users SET name = ?, email = ? WHERE id = ?",
		user.Name,
		user.Email,
		user.Id,
	)

	if err != nil {
		return 0, err
	}

	rows, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rows, nil
}

func (userModel UserModel) DeleteUser(id int64) (int64, error) {
	result, err := userModel.Database.Exec(
		"DELETE FROM users WHERE id = ?",
		id,
	)

	if err != nil {
		return 0, err
	}

	rows, err := result.RowsAffected()
		
	if err != nil {
		return 0, err
	}

	return rows, nil
}