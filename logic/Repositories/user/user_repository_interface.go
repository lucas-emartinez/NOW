package Repositories

import (
	"NOW/logic/Entities"
	"context"
)

type UserRepository interface {
	Create(ctx context.Context, user *Entities.User) error
	GetByDNI(ctx context.Context, dni int) (*Entities.User, error)
	Update(ctx context.Context, user *Entities.User) error
	Delete(ctx context.Context, dni int) error
	// Add methods here
}
