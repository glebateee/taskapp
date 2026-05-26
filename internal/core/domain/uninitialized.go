package domain

var (
	UninitilizedID      = -1
	UninitilizedVersion = -1
)

func NewUserUninitialized(
	fullName string,
	phoneNumber *string,
) User {
	return NewUser(
		UninitilizedID,
		UninitilizedVersion,
		fullName,
		phoneNumber,
	)
}
