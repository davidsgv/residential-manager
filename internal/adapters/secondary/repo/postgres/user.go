package postgres

import (
	"database/sql"
	"errors"
	"residential-manager/internal/domain/entities"

	"github.com/google/uuid"
)

func (repo *PostgresRepo) GetUsers() ([]entities.User, error) {
	var query string = `
		SELECT 
			usr.uuid, 
			rol.name, 
			usr.mail,
			usr.verified,
			apa.number,
			blo.block
		FROM user_account usr
		INNER JOIN rol
			ON usr.rol_id = rol.id
		INNER JOIN apartment apa
			ON usr.apartment_id = apa.id
		INNER JOIN block blo
			ON apa.block_id = blo.id
	`

	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []entities.User{}
	for rows.Next() {
		user := entities.User{
			Apartment: &entities.Apartment{},
		}

		rows.Scan(&user.Id, &user.Rol, &user.Mail, &user.MailVerified, &user.Apartment.Number, &user.Apartment.Block)
		users = append(users, user)
	}

	return users, nil
}

func (repo *PostgresRepo) CreateUser(user *entities.User) error {
	query := `call sp_create_user($1, $2, $3, $4, $5, $6, $7, $8, $9);`
	args := []any{
		user.Id.String(),
		user.Mail,
		user.Password,
		user.MailVerified,
		user.Rol,
		user.Apartment.Block,
		user.Apartment.Number,
		user.Token.Token,
		user.Token.Expire,
	}

	result, err := repo.db.Exec(query, args...)
	if err != nil {
		return err
	}
	affectedRows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affectedRows < 0 {
		return errors.New(ErrNotAffectedRows)
	}
	return nil
}

func (repo *PostgresRepo) GetUserByMail(mail string) (*entities.User, error) {
	var query string = `
		SELECT 
			usr.uuid, 
			rol.name, 
			usr.mail,
			usr.verified,
			apa.number,
			blo.block
		FROM user_account usr
		INNER JOIN rol
			ON usr.rol_id = rol.id
		INNER JOIN apartment apa
			ON usr.apartment_id = apa.id
		INNER JOIN block blo
			ON apa.block_id = blo.id
		WHERE usr.mail = $1
	`

	row := repo.db.QueryRow(query, mail)

	user := entities.User{Apartment: &entities.Apartment{}}
	err := row.Scan(&user.Id, &user.Rol, &user.Mail, &user.MailVerified,
		&user.Apartment.Number, &user.Apartment.Block)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (repo *PostgresRepo) GetUserById(id string) (*entities.User, error) {
	var query string = `
		SELECT 
			usr.uuid, 
			rol.name, 
			usr.mail,
			usr.verified,
			apa.number,
			blo.block
		FROM user_account usr
		INNER JOIN rol
			ON usr.rol_id = rol.id
		INNER JOIN apartment apa
			ON usr.apartment_id = apa.id
		INNER JOIN block blo
			ON apa.block_id = blo.id
		WHERE usr.uuid = $1
	`

	row := repo.db.QueryRow(query, id)

	user := entities.User{Apartment: &entities.Apartment{}}
	err := row.Scan(&user.Id, &user.Rol, &user.Mail, &user.MailVerified, &user.Apartment.Number, &user.Apartment.Block)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (repo *PostgresRepo) GetUserByToken(token string) (*entities.User, error) {
	var query string = `
		SELECT 
			usr.uuid, 
			rol.name, 
			usr.mail,
			usr.verified,
			apa.number,
			blo.block,
			usr.token,
			usr.token_expire
		FROM user_account usr
		INNER JOIN rol
			ON usr.rol_id = rol.id
		INNER JOIN apartment apa
			ON usr.apartment_id = apa.id
		INNER JOIN block blo
			ON apa.block_id = blo.id
		WHERE usr.token = $1
	`

	row := repo.db.QueryRow(query, token)

	user := entities.User{Apartment: &entities.Apartment{}}
	err := row.Scan(&user.Id, &user.Rol, &user.Mail,
		&user.MailVerified, &user.Apartment.Number, &user.Apartment.Block,
		&user.Token.Token, &user.Token.Expire)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (repo *PostgresRepo) GetUserWithCredentials(mail string) (*entities.User, error) {
	var query string = `
		SELECT 
			usr.uuid,
			rol.name, 
			usr.mail,
			usr.password,
			usr.verified
		FROM user_account usr
		INNER JOIN rol
			ON usr.rol_id = rol.id
		INNER JOIN apartment apa
			ON usr.apartment_id = apa.id
		INNER JOIN block blo
			ON apa.block_id = blo.id
		WHERE usr.mail = $1
	`

	row := repo.db.QueryRow(query, mail)

	user := entities.User{Apartment: &entities.Apartment{}}
	err := row.Scan(&user.Id, &user.Rol, &user.Mail, &user.Password, &user.MailVerified)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (repo *PostgresRepo) GetUserPermissions(mail string) ([]int, error) {
	var query string = `
		SELECT per.id
		FROM user_account usr
		INNER JOIN rol
			ON usr.rol_id = rol.id
		INNER JOIN permit_rol per_rol
			ON rol.id = per_rol.rol_id
		INNER JOIN permit per
			ON per_rol.permit_id = per.id
		WHERE usr.mail = $1
	`

	rows, err := repo.db.Query(query, mail)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	permissions := []int{}
	for rows.Next() {
		var id int
		rows.Scan(&id)
		permissions = append(permissions, id)
	}

	return permissions, nil
}

func (repo *PostgresRepo) UpdateUser(user *entities.User) error {
	query := `call sp_update_user($1, $2, $3, $4);`
	args := []any{user.Id.String(), user.Rol, user.Apartment.Block, user.Apartment.Number}

	result, err := repo.db.Exec(query, args...)
	if err != nil {
		return err
	}
	affectedRows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affectedRows < 0 {
		return errors.New(ErrNotAffectedRows)
	}
	return nil
}

func (repo *PostgresRepo) DeleteUser(id uuid.UUID) error {
	query := `
		DELETE 
		FROM user_account
		WHERE uuid = $1
	`

	result, err := repo.db.Exec(query, id)
	if err != nil {
		return err
	}
	_, err = result.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (repo *PostgresRepo) UpdatePassword(id uuid.UUID, password string) error {
	query := `
			UPDATE user_account
			SET password = $1
			WHERE uuid = $2
		`
	args := []any{password, id}

	result, err := repo.db.Exec(query, args...)
	if err != nil {
		return err
	}
	affectedRows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affectedRows < 0 {
		return errors.New(ErrNotAffectedRows)
	}
	return nil
}
