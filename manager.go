package auth

import (
	"context"
	"net/http"

	"github.com/bla005/auth/models"
	"github.com/golang/protobuf/proto"
)

func (auth *Auth) CreateUser(ctx context.Context, w http.ResponseWriter) (*models.User, error) {
	return nil, nil
}

func (auth *Auth) Login(ctx context.Context, w http.ResponseWriter, email, password string) (*models.User, error) {
	return nil, nil
}

func (auth *Auth) Logout(ctx context.Context, id, session string, w http.ResponseWriter) error {
	return nil
}

func (auth *Auth) ConfirmAccount(ctx context.Context, token string) error {
	return nil
}
func (auth *Auth) UpdateEmail(ctx context.Context, token, newEmail string, user *models.User, session proto.Message, sid string) error {
	return nil
}
