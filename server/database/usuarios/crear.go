package usuarios

import (
	"database/sql"
	"strconv"
)

func Crear(db *sql.DB, input NewUsuario) (*Usuario, error) {
	query := `
	insert into usuarios(nombre,apellidos, correo, username, password) values(?,?,?,?,?);
	`
	res, err := db.Exec(query, input.Nombre, input.Apellidos, input.Correo, input.Username, input.Password)
	if err != nil {
		return nil, err
	}

	idd, _ := res.LastInsertId()
	id := strconv.FormatInt(idd, 10)

	return Get(db, id)
}
