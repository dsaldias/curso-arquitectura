package actividades

import (
	"database/sql"
	"strconv"
)

func Crear(db *sql.DB, input NewActividad) (*Actividad, error) {
	query := `
	insert into actividades(nombre) values(?);
	`
	res, err := db.Exec(query, input.Nombre)
	if err != nil {
		return nil, err
	}

	idd, _ := res.LastInsertId()
	id := strconv.FormatInt(idd, 10)

	return Get(db, id)
}
