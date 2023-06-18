package apperrors

const (
	ErrInvalidTokenType = "invalid token type provided"
	ErrUpgradeRequired  = "sorry, you need to upgrade your current plan to access this resource"
	ErrValidation       = "one or more fields are invalid"
	ErrUnauthorized     = "you need to be logged in to access this resource"
	ErrForbidden        = "you don't have the permission to access this resource"
	ErrBadRequest       = "oops, the server received a bad request, try again"
)

type CustomError struct {
	message string
}

func New(message string) error {
	return &CustomError{message}
}

func (e *CustomError) Error() string {
	return e.message
}
