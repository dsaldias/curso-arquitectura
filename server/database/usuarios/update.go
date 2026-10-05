package usuarios

import (
	"database/sql"
	"errors"
)

func Actualizar(db *sql.DB, input NewUsuario) (*Usuario, error) {
	if input.Id == nil {
		return nil, errors.New("falta el id del usuario")
	}
	query := `
	update usuarios set 
		nombre = ?, 
		apellidos = ?, 
		correo=?, 
		username=?, 
		password = IF(
        NULLIF(?, '') IS NULL,
        password,
        SHA2(?, 256)
    )
		where id = ?;
	`
	_, err := db.Exec(query, input.Nombre, input.Apellidos, input.Correo, input.Username, input.Password, input.Password, input.Id)
	if err != nil {
		return nil, err
	}

	return Get(db, *input.Id)
}
