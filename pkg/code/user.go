package code

const (
	// ErrUserNotFound - 404: User not found.
	ErrUserNotFound int = iota + 100401

	// ErrUserAlreadyExists - 400: User already exists.
	ErrUserAlreadyExists

	// ErrUserPasswordIncorrect - 400: User password incorrect.
	ErrUserPasswordIncorrect

	// ErrSmsSend - 400: Send sms error.
	ErrSmsSend

	// ErrCodeNotExist - 400: Sms code incorrect or expired.
	ErrCodeNotExist

	// ErrCodeIncorrect - 400: Verify code incorrect.
	ErrCodeIncorrect
)
