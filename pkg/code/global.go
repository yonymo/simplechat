package code

//go:generate codegen -type=int
//go:generate codegen -type=int -doc -output ./error_code_generated.md

const (
	// ErrParam - 400: Param error.
	ErrParam int = iota + 100501
	// ErrTokenCreate - 400: Token create failed.
	ErrTokenCreate
	// ErrServerInternal - 500: Server internal error.
	ErrServerInternal
)
