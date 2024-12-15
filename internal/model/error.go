package model

type BaseError struct {
	Message string
}

func (e *BaseError) Error() string {
	return e.Message
}

type ErrBadRequest struct {
	BaseError
}

type ErrNotFound struct {
	BaseError
}

type ErrInternal struct {
	BaseError
}

type ErrInvalidToken struct {
	BaseError
}

type ErrInvalidTokenClaims struct {
	BaseError
}
