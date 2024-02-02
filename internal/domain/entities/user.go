package entities

const ErrInvalidUserInfo = "Invalid user info"

type Apartment struct {
	Block  string
	Number string
}

type User struct {
	Rol          string
	Mail         string
	User         string
	Password     string
	permissions  Permission
	MailVerified bool
}

func CreateUser(rol, mail, user string, permissions Permission) (*User, error) {
	//check rol
	//check email
	//check user

	//return nil, errors.New(ErrInvalidUserInfo)
	//if not valid data return error
	return nil, nil
}
