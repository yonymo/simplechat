package code

const (
	// ErrNicknameEmpty - 400: Nickname is empty.
	ErrNicknameEmpty int = iota + 100601
	// ErrFriendAlreadyExist - 400: Friend already exist.
	ErrFriendAlreadyExist
	// ErrFriendReqAlreadyCommit - 400: Friend request already commit.
	ErrFriendReqAlreadyCommit
)
