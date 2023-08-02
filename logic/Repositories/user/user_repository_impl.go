package Repositories

import (
	"NOW/db"
	"NOW/logic/Entities"
)

type UserRepositoryImplementation struct {
	db *db.Database
}

func (r *UserRepositoryImplementation) GetByDNI(dni int) (*Entities.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r *UserRepositoryImplementation) Update(user *Entities.User) error {
	//TODO implement me
	panic("implement me")
}

func (r *UserRepositoryImplementation) Delete(dni int) error {
	//TODO implement me
	panic("implement me")
}

func NewUserRepositoryImpl(db *db.Database) *UserRepositoryImplementation {
	return &UserRepositoryImplementation{db: db}
}

func (r *UserRepositoryImplementation) Create(user *Entities.User) error {
	// Implementa la lógica para insertar un nuevo usuario en la base de datos.
	sqlStatement := `
		INSERT INTO User (DNI, Name, LastName, PasswordHash, IdentityCheck)
		VALUES (?, ?, ?, ?, ?)
	`

	_, err := r.db.GetDB().Exec(sqlStatement, user.DNI, user.Name, user.LastName, user.PasswordHash, user.IdentityCheck)
	if err != nil {
		return err
	}

	return nil
}

// Not implemented

//func (r *UserRepositoryImplementation) GetByDNI(dni int) (*models.User, error) {
// Implementa la lógica para buscar un usuario por DNI.
//}
// Not implemented
//func (r *UserRepositoryImplementation) Update(user *models.User) error {
// Implementa la lógica para actualizar un usuario existente.
//}
// Not implemented
//func (r *UserRepositoryImplementation) Delete(dni int) error {
// Implementa la lógica para eliminar un usuario.

//}
