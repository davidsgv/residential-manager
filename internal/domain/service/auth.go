package service

import (
	"errors"
	"residential-manager/internal/common/response"
	"residential-manager/internal/common/validator"
	"residential-manager/internal/domain/entities"
	repo "residential-manager/internal/domain/ports/repository"

	"github.com/golang-jwt/jwt/v5"
)

const ErrInvalidCredentials = "Invalid credentials"
const ErrNotLogedIn = "You must be logged"
const ErrNotVerifiedMail = "Your account is not verified"

type AuthSrv interface {
	//CreateUser(rol, mail, username string) (response response.Response[entities.User])
	Login(mail, password string) (response response.Response[string])
	CheckLogin(tokenstr string) CheckLoginFuncData
	CheckPermission(userMail string, permissions ...int) (bool, error)
}

// checkPermission
type CheckPermissionFunc func(userMail string, permissions ...int) (bool, error)

type CheckLoginFuncData struct {
	Mail  string
	Rol   string
	Error error
}

type AuthConfig struct {
	SecretKey string
	Domain    string
}

type AuthService struct {
	repo   repo.UserRepo
	config AuthConfig
}

func NewAuthService(repo repo.UserRepo, config AuthConfig) *AuthService {
	return &AuthService{
		repo:   repo,
		config: config,
	}
}

func (srv *AuthService) Login(mail, password string) (response response.Response[string]) {
	v := validator.New()

	//validate email
	entities.ValidateMail(v, mail)

	//check validations
	if !v.Valid() {
		return response.ResponseFail(v.Errors)
	}

	//get user data
	user, err := srv.repo.GetUserWithCredentials(mail)
	if err != nil {
		return response.ResponseError(err.Error())
	}

	//hash password
	hashPass := user.HashPassword(password)

	//compare passwords
	if user == nil || user.Password != hashPass {
		return response.ResponseFailMessage(ErrInvalidCredentials)
	}
	if !user.MailVerified {
		return response.ResponseFailMessage(ErrNotVerifiedMail)
	}

	//firm token
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": srv.config.Domain,
		"sub": mail,
		"rol": user.Rol,
	})

	token, err := t.SignedString([]byte(srv.config.SecretKey))
	if err != nil {
		return response.ResponseError(err.Error())
	}

	return response.ResponseSuccess(token)
}

func (srv *AuthService) CheckLogin(tokenstr string) CheckLoginFuncData {
	jwt.WithIssuer(srv.config.Domain)
	token, err := jwt.Parse(tokenstr, func(t *jwt.Token) (interface{}, error) {
		// if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
		// 	return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		// }
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(srv.config.SecretKey), nil
	})

	if err != nil {
		return CheckLoginFuncData{
			Error: err,
		}
	}

	if !token.Valid {
		return CheckLoginFuncData{
			Error: errors.New(ErrNotLogedIn),
		}
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return CheckLoginFuncData{
			Error: errors.New(ErrNotLogedIn),
		}
	}

	return CheckLoginFuncData{
		Rol:   claims["rol"].(string),
		Mail:  claims["sub"].(string),
		Error: nil,
	}
}

// check in the database if the user hace the permissions
// if not return false and the error
func (srv *AuthService) CheckPermission(userMail string, permissions ...int) (bool, error) {
	userPermissions, err := srv.repo.GetUserPermissions(userMail)
	if err != nil {
		return false, err
	}

	userCalcPermit := 0
	for _, permit := range userPermissions {
		userCalcPermit = userCalcPermit | permit
	}

	for _, permit := range permissions {
		if userCalcPermit&permit == 0 {
			return false, nil
		}
	}

	return true, nil
}
