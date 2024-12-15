package core

import (
	"context"
	"log/slog"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"smart_school_for_mirea/internal/jwt"
	"smart_school_for_mirea/internal/model"
)

func (c *Core) SignUp(ctx context.Context, params model.SignUpRequest) (token string, err error) {
	const op = "authCore.SignUp"

	log := c.logger.With(
		slog.String("op", op),
		slog.String("email", params.Email),
	)

	log.Info("registering new user")

	passHash, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error("failed to generate password hash", err.Error())

		return token, err
	}

	user, err := c.storage.CreateUser(ctx, model.User{
		UUID:         uuid.New(),
		Email:        params.Email,
		HashPassword: string(passHash),
	})
	if err != nil {
		log.Error("failed to create user", err.Error())

		return token, err
	}

	log.Info("successfully created user", user)

	token, err = jwt.NewToken(user, c.authSecretKey, time.Duration(c.authTTL)*time.Hour*24)
	if err != nil {
		log.Error("failed to generate token", err.Error())

		return token, err
	}

	return token, nil
}

func (c *Core) SignIn(ctx context.Context, params model.SignInRequest) (token string, err error) {
	const op = "authCore.SignIn"

	log := c.logger.With(
		slog.String("op", op),
		slog.String("email", params.Email),
	)

	log.Info("login for user")

	user, err := c.storage.GetUserByUsername(ctx, params.Email)
	if err != nil {
		if _, ok := err.(*model.ErrNotFound); ok {
			return "", &model.ErrNotFound{BaseError: model.BaseError{Message: "invalid credential"}}
		}
		log.Error("failed to fetch user", err.Error())

		return token, &model.ErrInternal{BaseError: model.BaseError{Message: "Internal error"}}
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.HashPassword), []byte(params.Password)); err != nil {
		log.Info("invalid password", err.Error())

		return "", &model.ErrNotFound{BaseError: model.BaseError{Message: "invalid credential"}}
	}

	token, err = jwt.NewToken(user, c.authSecretKey, time.Duration(c.authTTL)*time.Hour*24)
	if err != nil {
		log.Error("failed to generate token", err.Error())

		return token, err
	}

	return token, nil
}
