package Repositories

import (
	"NOW/db"
	"NOW/logic/Entities"
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

func (r *UserRepositoryImplementation) GetByDNI(ctx context.Context, dni int) (*Entities.User, error) {
	//TODO implement me
	sqlStatement := `SELECT * 
					 FROM  User
					 WHERE DNI = ?`

	row := r.db.GetDB().QueryRowContext(ctx, sqlStatement, dni)
	user := &Entities.User{}
	err := row.Scan(&user.DNI, &user.Name, &user.LastName, &user.Password, &user.IdentityCheck)

	if err == sql.ErrNoRows {
		return nil, nil
	} else {
		return nil, err
	}

	return user, nil
}

func (r *UserRepositoryImplementation) Update(ctx context.Context, user *Entities.User) error {
	//TODO implement me

	// Implementa la lógica para actualizar un usuario existente en la base de datos.
	sqlStatement := `
		UPDATE User
		SET Name = ?, LastName = ?, Password = ?, IdentityCheck = ?
		WHERE DNI = ?
	`

	_, err := r.db.GetDB().ExecContext(ctx, sqlStatement, user.Name, user.LastName, user.Password, user.IdentityCheck, user.DNI)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepositoryImplementation) Delete(ctx context.Context, dni int) error {
	//TODO implement me
	panic("implement me")
}

func (r *UserRepositoryImplementation) Create(ctx context.Context, user *Entities.User) error {
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
