package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"residential-manager/internal/adapters/primary/api"
	"residential-manager/internal/adapters/secondary/notification"
	"residential-manager/internal/adapters/secondary/repo/postgres"
	"residential-manager/internal/domain/service"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

const version = "1.0.0"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//repos
	repo := setUpPostgreRepo()
	defer repo.CloseDB()
	if err != nil {
		log.Fatal("Could not connect to DB")
	}
	notificator := setUpEmailSender()

	//services
	// authService := service.NewAuthService(repo, service.AuthConfig{
	authService := service.NewAuthService(repo, service.AuthConfig{
		SecretKey: os.Getenv("SECRET_KEY"),
		Domain:    os.Getenv("DOMAIN"),
	})
	userService := service.NewUserService(repo, notificator, authService.CheckPermission, service.MailData{
		LogoURL:  os.Getenv("REGISTER_MAIL_LOGO"),
		TokenURL: os.Getenv("REGISTER_MAIL_TOKEN_URL"),
		Path:     os.Getenv("REGISTER_MAIL"),
	})

	var cfg = api.Config{
		Env: os.Getenv("ENVIRONMENT"),
	}
	cfg.Port, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := api.NewAplication(&cfg, api.Services{
		AuthService: authService,
		UserService: userService,
	})
	err = app.Fiber.Listen(fmt.Sprintf("0.0.0.0:%d", cfg.Port))
	fmt.Println(err)
}

func setUpPostgreRepo() *postgres.PostgresRepo {
	var maxOpenConns, maxIdleConns *int
	var connMaxIdleTime, connMaxLifetime *time.Duration
	var err error

	maxOpenConns, err = getEnvInt("DB_POOL_MAX_OPEN_CONNS")
	if err != nil {
		log.Fatal(err)
	}
	maxIdleConns, err = getEnvInt("DB_POOL_MAX_IDLE_CONNS")
	if err != nil {
		log.Fatal(err)
	}
	connMaxIdleTime, err = getEnvDuration("DB_POOL_CONN_MAX_IDLE_TIME")
	if err != nil {
		log.Fatal(err)
	}
	connMaxLifetime, err = getEnvDuration("DB_POOL_CONN_MAX_LIFE_TIME")
	if err != nil {
		log.Fatal(err)
	}

	repo, err := postgres.NewPostgresRepo(&postgres.PostgresConfig{
		Host:     os.Getenv("POSTGRES_DB_HOST"),
		Port:     os.Getenv("POSTGRES_DB_PORT"),
		User:     os.Getenv("POSTGRES_DB_USER"),
		Password: os.Getenv("POSTGRES_DB_PASSWORD"),
		DBname:   os.Getenv("POSTGRES_DB_NAME"),
		Pool: postgres.PoolConfig{
			MaxOpenConns:    maxOpenConns,
			MaxIdleConns:    maxIdleConns,
			ConnMaxIdleTime: connMaxIdleTime,
			ConnMaxLifetime: connMaxLifetime,
		},
	})
	if err != nil {
		log.Fatal("Could not connect to DB")
	}

	return repo
}

func setUpEmailSender() *notification.Mail {
	smptPort, err := getEnvInt("SMTP_PORT")
	if err != nil {
		log.Fatal(err)
	}
	mailConfig := notification.EmailConfig{
		Mail:     os.Getenv("SMTP_MAIL"),
		Password: os.Getenv("SMTP_PASSWORD"),
		SmptHost: os.Getenv("SMTP_HOST"),
		SmptPort: *smptPort,
	}
	mail := notification.NewMail(mailConfig)

	return mail
}

func getEnvInt(s string) (*int, error) {
	value, err := strconv.Atoi(os.Getenv(s))
	if err != nil {
		return nil, err
	}
	if value <= 0 {
		errS := fmt.Sprintf("var %s must be greather than 0, value: %d", s, value)
		return nil, errors.New(errS)
	}
	return &value, nil
}

func getEnvDuration(s string) (*time.Duration, error) {
	value, err := time.ParseDuration(os.Getenv(s))
	if err != nil {
		return nil, err
	}
	if value < 0 {
		errS := fmt.Sprintf("var %s must be greather than 0, value: %d", s, value)
		return nil, errors.New(errS)
	}
	return &value, nil
}
