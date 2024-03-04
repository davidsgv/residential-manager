package entities

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"residential-manager/internal/common/validator"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/sha3"
)

// errors
const (
	ErrInvalidMail       = "Invalid user mail"
	ErrInvalidRol        = "Rol provided not exists"
	ErrLenPassword       = "Password must contain 8 chars"
	ErrLenAparmentBlock  = "block must contain at least 1 char"
	ErrLenAparmentNumber = "number must contain at least 1 char"
)

// validations
const (
	MinPasswordLen       = 8
	MinAparmentBlockLen  = 1
	MinAparmentNumberLen = 1
)

// configurations
const (
	tokenNonce          = "5050"
	expirationTimeToken = time.Hour * 24 * 7
)

type Apartment struct {
	Block  string
	Number string
}

type Token struct {
	Token  string
	Expire time.Time
}

type User struct {
	Id           uuid.UUID
	Rol          *string
	Mail         string
	Password     string
	MailVerified bool
	Apartment    *Apartment
	Token        Token
}

func CreateUser(rol, mail, block, number string) (*User, map[string]string) {
	user := User{
		Id:           uuid.New(),
		Rol:          &rol,
		Mail:         mail,
		MailVerified: false,
		Apartment: &Apartment{
			Block:  block,
			Number: number,
		},
		Token: Token{
			Expire: time.Now().Add(expirationTimeToken),
		},
	}

	byteData := []byte(tokenNonce + user.Id.String() + user.Mail + *user.Rol)
	hash := sha256.Sum256(byteData)
	user.Token.Token = base64.URLEncoding.EncodeToString(hash[:])

	if block == "" && number == "" {
		user.Apartment = nil
	}

	v := validator.New()
	user.ValidateUserInfo(v)

	//return nil, errors.New(ErrInvalidUserInfo)
	//if not valid data return error
	if !v.Valid() {
		return nil, v.Errors
	}

	return &user, nil
}

func (user *User) HashPassword(password string) string {
	hash := sha3.New256()
	// _, err := hash.Write([]byte(input))
	// if err != nil {
	// 	return "", nil
	// }

	sha3 := hash.Sum([]byte(password))
	return fmt.Sprintf("%x", sha3)
}

func (user *User) ValidateUserInfo(v *validator.Validator) {
	//checkRoles
	ValidateRol(v, *user.Rol)

	//check email
	ValidateMail(v, user.Mail)
	if *user.Rol == RolEncargadoApartamento || *user.Rol == RolResidente {
		ValidateApartment(v, user.Apartment.Block, user.Apartment.Number)
	}
}

func ValidateRol(v *validator.Validator, rol string) {
	roles := GetAllRoles()
	if !validator.PermittedValue(rol, roles...) {
		v.AddError("rol", ErrInvalidRol)
	}
}

func ValidateMail(v *validator.Validator, mail string) {
	validMail := validator.Matches(mail, validator.EmailRX)
	v.Check(validMail, "mail", ErrInvalidMail)
}

func ValidatePassword(v *validator.Validator, password string) {
	v.Check(MinPasswordLen <= len(password), "password", ErrLenPassword)
}

func ValidateApartment(v *validator.Validator, block, number string) {
	v.Check(len(block) >= MinAparmentBlockLen, "apartment block", ErrLenAparmentBlock)
	v.Check(len(number) >= MinAparmentNumberLen, "apartment number", ErrLenAparmentNumber)
}
