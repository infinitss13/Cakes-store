package repository

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/infinitss13/Cakes-store/config"
	"github.com/infinitss13/Cakes-store/entities"
	_ "github.com/lib/pq"
)

type Database struct {
	*sql.DB
	migrations *migrate.Migrate
}

func NewDatabase(config config.DBConfig) (*Database, error) {
	dataBase, err := sql.Open("postgres", config.ConnectionDbData())
	if err != nil {
		return nil, err
	}
	err = dataBase.Ping()
	if err != nil {
		return nil, err
	}
	driver, err := postgres.WithInstance(dataBase, &postgres.Config{})
	if err != nil {
		return nil, fmt.Errorf("with instance failed: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://schema", "postgres", driver)
	if err != nil {
		return nil, fmt.Errorf("new with database instance failed: %w", err)
	}
	err = m.Up()
	if err != migrate.ErrNoChange && err != nil {
		return nil, fmt.Errorf("migrate up failed: %w", err)
	}
	return &Database{
		dataBase,
		m,
	}, nil
}

type Repository interface {
	CreateUser(userData *entities.UserPersonalData) error
	GetPasswordByNumber(phoneNumber string) (string, error)
}

func (db *Database) IsUserExist(email, phoneNumber string) (bool, error) {
	id := 0
	query := "SELECT id FROM usersData WHERE email=$1 OR phoneNumber=$2;"
	err := db.QueryRow(query, email, phoneNumber).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil

}
func (db *Database) CreateUser(userData *entities.UserPersonalData) error {
	query := "INSERT INTO usersData (firstName, lastName, email, phoneNumber,hashPassword, dateOfBirth, role, imgUrl) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id"
	row := db.QueryRow(query,
		userData.FirstName,
		userData.LastName,
		userData.Email,
		userData.PhoneNumber,
		userData.Password,
		userData.DateOfBirth,
		userData.Role,
		userData.ImgUrl)
	var id = 0
	if err := row.Scan(&id); err != nil {
		return err
	}
	return nil
}

func (db *Database) GetPasswordByNumber(phoneNumber string) (string, error) {
	passwordHash := ""
	query := "SELECT hashPassword FROM usersData WHERE phoneNumber=$1"
	err := db.QueryRow(query, phoneNumber).Scan(&passwordHash)
	if err != nil {
		return "", err
	}
	return passwordHash, nil
}
