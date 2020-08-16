package auth

import (
	"context"
	"database/sql"
	"time"

	"github.com/gofrs/uuid"
)

const (
	SessionLen = 32
	TokenLen   = 32
)

type AuthStore interface {
	SaveSession(prefix, userID, sessionID string, sessionObject interface{}, sessionDuration time.Duration) error
	GetSession(sid string) ([]byte, error)
	GetOneSession(prefix, id string) (string, error)
	CountSessions(prefix, userID string) (int64, error)
	RemoveSession(prefix, id, session string) error
	GetAllSessions(id string) ([]string, error)
	UpdateSessions(sessions []string, session []byte) error
	UpdateCurrentSession(prefix, sessionID string, sessionObject []byte, sessionDuration time.Duration) error

	RemoveOneSession(prefix, userID string) error

	NewAuthFail(prefix, userID string) error
	GetAuthFails(prefix, userID string) (int, error)
	ExistsAuthFail(prefix, userID string) (int64, error)
	IncrAuthFail(prefix, userID string) error
	DelAuthFail(prefix, userID string) error
}
type AuthRepository interface {
	GetUserByID(ctx context.Context, id uuid.UUID) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetUserByUsername(ctx context.Context, username string) (*User, error)
	CreateUser(ctx context.Context, user *User) (*User, error)
	CreateUserTx(ctx context.Context, user *User, tx *sql.Tx) (*User, error)
	GetPasswordByID(ctx context.Context, id uuid.UUID) ([]byte, error)
	UpdatePasswordByID(ctx context.Context, newPassword []byte, id uuid.UUID) error
	UpdateConfirmedByID(ctx context.Context, id uuid.UUID) (bool, error)
	UpdateEmailByID(ctx context.Context, newEmail string, id uuid.UUID) error
	UpdateUsernameByID(ctx context.Context, newUsername string, id uuid.UUID) error
	CreateToken(ctx context.Context, token *Token) (*Token, error)
	CreateTokenTx(ctx context.Context, token *Token, tx *sql.Tx) (*Token, error)
	GetToken(ctx context.Context, value string) (*Token, error)
	DelToken(ctx context.Context, value string) error
}

type Auth struct {
}

func New() *Auth {
	return &Auth{}
}
