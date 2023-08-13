package code

const (
	// ErrInvalidAuthHeader - 400: Invalid auth header.
	ErrInvalidAuthHeader int = iota + 100701
	// ErrSignatureInvalid - 400: Signature invalid.
	ErrSignatureInvalid
)
