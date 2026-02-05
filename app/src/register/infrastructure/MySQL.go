package infrastructure

import (
	"database/sql"

	"github.com/BryanChanona/api_chombi.git/app/src/register/domain/entities"
)

type MySQL struct{
	db *sql.DB
}

func NewMySQL(db *sql.DB)*MySQL{
	return &MySQL{db: db}
}

func (mysql *MySQL) Save(user entities.User) error {
	query := "INSERT INTO users (name, lastName, email, passwordHash) VALUES (?, ?, ?, ?)"
	_, err := mysql.db.Exec(query, user.Name, user.LastName, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil

}
