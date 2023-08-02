package Entities

type User struct {
	DNI           int
	Name          string
	LastName      string
	PasswordHash  string
	RefreshToken  string
	IdentityCheck bool
}
