package auth

import (
	"time"

	wtf "github.com/bla005/auth/models"
	"github.com/gofrs/uuid"
	"google.golang.org/protobuf/proto"
)

type User struct {
	ID          uuid.UUID `json:"-"`
	Email       string    `json:"email"`
	Username    string    `json:"username"`
	Password    []byte    `json:"-"`
	IsStaff     bool      `json:"is_staff"`
	IsConfirmed bool      `json:"is_confirmed"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
}

func (auth *Auth) NewUser() (*User, error) {
	userId, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	return &User{
		ID:          userId,
		Email:       "",
		Username:    "",
		Password:    nil,
		IsStaff:     false,
		IsConfirmed: false,
		IsActive:    true,
		CreatedAt:   time.Now(),
	}, nil
}

func (user *User) GetUsername() string {
	return user.Username
}
func (user *User) GetEmail() string {
	return user.Email
}
func (user *User) SetPassword(password string) error {
	hashedPassword, err := hashPassword(password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	return nil
}
func (user *User) SetUsername(username string) {
	user.Username = username
}
func (user *User) SetEmail(email string) {
	user.Email = email
}
func (user *User) CheckPassword(password string) error {
	return checkPassword(user.Password, password)
}
func (user *User) ToSessionObject() ([]byte, error) {
	sessionPb := &wtf.Session{
		Id:           "",
		Email:        user.Email,
		Username:     user.Username,
		IsStaff:      user.IsStaff,
		IsConfirmed:  user.IsConfirmed,
		IsActive:     user.IsActive,
		LastLogin:    0,
		Subscription: 0,
	}
	return proto.Marshal(sessionPb)
}
