package model

type BaseError struct {
	Message string
}

func (e *BaseError) Error() string {
	return e.Message
}

type ErrNotFound struct {
	BaseError
}

type ErrInvalidToken struct {
	BaseError
}

type ErrInvalidTokenClaims struct {
	BaseError
}
