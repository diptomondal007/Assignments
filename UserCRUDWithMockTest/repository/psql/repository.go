package psql

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"usercrudwithmocktest/crud"
	"usercrudwithmocktest/crud/models"
)

type psqlRepository struct {
	db *sql.DB
}

func (p psqlRepository) Get(page, limit int) ([]models.User, error) {
	rows, err := p.db.Query("SELECT * from users limit $1 offset $2", limit, limit*(page-1))
	if err != nil {
		return nil, err
	}
	var users []models.User
	users = make([]models.User, 0)
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Username, &user.Name)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (p psqlRepository) GetByUsername(username string) (*models.User, error) {
	var user models.User
	row, err := p.db.Query("select * from users where username=$1", username)
	if err != nil {
		return nil, err
	} else {
		for row.Next() {
			err := row.Scan(&user.Username, &user.Name)
			if err != nil {
				return nil, err
			}
		}

	}
	return &user, nil
}

func (p psqlRepository) Create(user *models.User) error {
	_, err := p.db.Query("INSERT INTO users VALUES ($1,$2)", user.Username, user.Name)
	if err != nil {
		return err
	}
	return nil
}

func (p psqlRepository) Delete(username string) (int64,error) {
	sqlResult, err := p.db.Exec("delete from users where username=$1", username)
	if err != nil {
		return 0, err
	}
	rowsAffectedCount, err := sqlResult.RowsAffected()
	if rowsAffectedCount < 1 {
		return 0, nil
	}
	return rowsAffectedCount, nil
}

func (p psqlRepository) Update(username string, user *models.User) error {
	if p.rowExists("select username from users where username=$1", username) {
		rows, err := p.db.Query("select username from users where username=$1", user.Username)
		if err != nil {
			return err
		}
		for rows.Next() {
			var username string
			_ = rows.Scan(&username)
			if username == user.Username {
				err = errors.New("username already exists")
				return err
			}
		}
		_, err = p.db.Query("update users set username=$1, name =$2 where username=$3", user.Username, user.Name, username)
		if err != nil {
			err = errors.New("error updating user")
			return err
		}
		return nil
	}
	err := errors.New("user not found")
	return err
}

func NewPsqlRepository(db *sql.DB) (crud.Repository, error) {
	repo := &psqlRepository{}
	repo.db = db
	return repo, nil
}

func(p psqlRepository) rowExists(query string, args ...interface{}) bool {
	var exists bool
	query = fmt.Sprintf("SELECT exists (%s)", query)
	err := p.db.QueryRow(query, args...).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		log.Fatalf("error checking if row exists '%s' %v", args, err)
	}
	return exists
}