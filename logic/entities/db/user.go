package db

type User struct {
	DNI           int
	Name          string
	LastName      string
	Password      string
	IdentityCheck bool
}
