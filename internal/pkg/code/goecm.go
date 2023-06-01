package code

const (
	// ErrUserNotFound - 404: user not found.
	ErrUserNotFound int = iota + 110001

	// ErrUserAlreadyExist - 400: user already exist. code: 110002
	ErrUserAlreadyExist

	// ErrPrivateKeyNotFound - 500: private key not found. code: 110003
	ErrPrivateKeyNotFound

	// ErrSystemStorageIssue - 500: save file failed. code: 110004
	ErrSystemStorageIssue

	// ErrCommandExec - 500: command exec failed. code: 110005
	ErrCommandExec

	// ErrAlreadyRegister - 200: agent already registered. code: 110006
	ErrAlreadyRegister

	// ErrImageUpload - 200: can not save multipart file data. code: 110007
	ErrImageUpload

	ErrImageImportMerge
)
