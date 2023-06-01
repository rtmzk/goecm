package code

const (
	ErrSuccess int = iota + 100001

	ErrUnknown

	ErrBind

	ErrValidation

	ErrTokenInvalid

	ErrPageNotFound
)

const (
	ErrDatabase int = iota + 100101
	ErrSignatureInvalid
	ErrPermissionDenied

	ErrSSHConnectionFailed
)

const (
	ErrUnmarshal int = iota + 100010
)
