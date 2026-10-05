package actividades

import (
	"database/sql"
	"errors"
)

func Actualizar(db *sql.DB, input NewActividad) (*Actividad, error) {
	if input.Id == nil {
		return nil, errors.New("falta el id de la actividad")
	}
	query := `
	update actividades set 
		nombre = ?
		where id = ?;
	`
	_, err := db.Exec(query, input.Nombre, input.Id)
	if err != nil {
		return nil, err
	}

	return Get(db, *input.Id)
}
