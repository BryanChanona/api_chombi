package infrastructure

import (
	"database/sql"
	"errors"

	"github.com/BryanChanona/api_chombi.git/app/src/login/domain/entities"
)

type MySQL struct {
	db *sql.DB
}

func NewMySQL(db *sql.DB) *MySQL {
	return &MySQL{db: db}
}


func (mysql *MySQL) LogIn(email string) (entities.UserWithPassword, error) {
	query := `
		SELECT id, name, lastName, email, passwordHash
		FROM users
		WHERE email = ?
	`

	var user entities.UserWithPassword

	err := mysql.db.QueryRow(query, email).
		Scan(&user.Id, &user.Name, &user.LastName, &user.Email, &user.PasswordHash)

	if err != nil {
		return entities.UserWithPassword{}, errors.New("user not found")
	}

	return user, nil
}