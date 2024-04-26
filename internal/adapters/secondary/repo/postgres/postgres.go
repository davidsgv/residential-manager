package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

const ErrNotAffectedRows = "not affected Rows"

type PoolConfig struct {
	ConnMaxLifetime *time.Duration //SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	ConnMaxIdleTime *time.Duration //SetConnMaxIdleTime sets the maximum amount of time a connection may be idle.
	MaxIdleConns    *int           //SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	MaxOpenConns    *int           //SetMaxOpenConns sets the maximum number of open connections to the database.
}

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBname   string
	Pool     PoolConfig
}

type PostgresRepo struct {
	db *sql.DB
}

func NewPostgresRepo(config *PostgresConfig) (*PostgresRepo, error) {
	//TODO add to postgres config an option to set the sslmode, dint hardcode this option.
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DBname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	if config.Pool.ConnMaxLifetime != nil {
		db.SetConnMaxLifetime(*config.Pool.ConnMaxLifetime)
	}
	if config.Pool.ConnMaxIdleTime != nil {
		db.SetConnMaxIdleTime(*config.Pool.ConnMaxIdleTime)
	}
	if config.Pool.MaxIdleConns != nil {
		db.SetMaxIdleConns(*config.Pool.MaxIdleConns)
	}
	if config.Pool.MaxOpenConns != nil {
		db.SetMaxOpenConns(*config.Pool.MaxOpenConns)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	pingErr := db.PingContext(ctx)
	if pingErr != nil {
		panic(pingErr)
	}

	return &PostgresRepo{
		db: db,
	}, nil
}

func (repo *PostgresRepo) CloseDB() {
	repo.db.Close()
}
