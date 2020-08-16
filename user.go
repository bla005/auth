package auth

import (
	"time"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
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

func NewUser() (*User, error) {
	userID, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	return &User{
		ID:          userID,
		Email:       "",
		Username:    "",
		Password:    nil,
		IsStaff:     false,
		IsConfirmed: false,
		IsActive:    true,
		CreatedAt:   time.Now(),
	}, nil
}

func hashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func checkPassword(passwordHash []byte, password string) error {
	return bcrypt.CompareHashAndPassword(passwordHash, []byte(password))
}

func (user *User) SetPassword(password string) error {
	hashedPassword, err := hashPassword(password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	return nil
}

func (user *User) ToSessionObject() ([]byte, error) {
	sessionPb := &Session{
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
