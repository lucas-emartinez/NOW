package Repositories

import (
	dbEntity "NOW/rest_service/logic/entities/db"
	"context"
)

type UserRepository interface {
	Create(ctx context.Context, user *dbEntity.User) error
	GetByDNI(ctx context.Context, dni int) (*dbEntity.User, error)
	Update(ctx context.Context, user *dbEntity.User) error
	Delete(ctx context.Context, dni int) error
	// Add methods here
}
