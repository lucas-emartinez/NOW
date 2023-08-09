package Repositories

import (
	"NOW/db"
	dbEntity "NOW/logic/entities/db"
	"context"
	"database/sql"
	"golang.org/x/crypto/bcrypt"
)

type UserRepositoryImplementation struct {
	db *db.Database
}

func NewUserRepositoryImpl(db *db.Database) *UserRepositoryImplementation {
	return &UserRepositoryImplementation{db: db}
}

func (r *UserRepositoryImplementation) GetByDNI(ctx context.Context, dni int) (*dbEntity.User, error) {
	sqlStatement := `SELECT DNI, Name, LastName, IdentityCheck 
					 FROM  User
					 WHERE DNI = ?`

	row := r.db.GetDB().QueryRowContext(ctx, sqlStatement, dni)
	user := &dbEntity.User{}
	err := row.Scan(&user.DNI, &user.Name, &user.LastName, &user.IdentityCheck)

	if err == sql.ErrNoRows {
		return nil, err
	} else {
		return user, err
	}
}
func (r *UserRepositoryImplementation) Update(ctx context.Context, user *dbEntity.User) error {

	// Implementa la lógica para actualizar un usuario existente en la base de datos.
	sqlStatement := `
		UPDATE User
		SET Name = ?, LastName = ?, IdentityCheck = ?
		WHERE DNI = ?
	`

	_, err := r.db.GetDB().ExecContext(ctx, sqlStatement, user.Name, user.LastName, user.IdentityCheck, user.DNI)
	if err != nil {
		return err
	}

	return nil
}
func (r *UserRepositoryImplementation) Delete(ctx context.Context, dni int) error {
	// Implementa la lógica para eliminar un usuario existente en la base de datos.
	sqlStatement := `
		DELETE FROM User
		WHERE DNI = ?
	`

	_, err := r.db.GetDB().ExecContext(ctx, sqlStatement, dni)
	if err != nil {
		return err
	}

	return nil
}
func (r *UserRepositoryImplementation) Create(ctx context.Context, user *dbEntity.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	if err != nil {
		return err
	}

	// Implementa la lógica para insertar un nuevo usuario en la base de datos.
	sqlStatement := `
		INSERT INTO User (DNI, Name, LastName, Password, IdentityCheck)
		VALUES (?, ?, ?, ?, ?)
	`

	_, err = r.db.GetDB().ExecContext(ctx, sqlStatement, user.DNI, user.Name, user.LastName, hashedPassword, user.IdentityCheck)
	if err != nil {
		return err
	}

	return err
}
