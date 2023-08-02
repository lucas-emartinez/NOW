package Repositories

import "NOW/logic/Entities"

type UserRepository interface {
	Create(user *Entities.User) error
	GetByDNI(dni int) (*Entities.User, error)
	Update(user *Entities.User) error
	Delete(dni int) error
	// Add methods here
}
