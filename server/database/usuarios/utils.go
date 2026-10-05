package usuarios

import "database/sql"

type NewUsuario struct {
	Id        *string
	Nombre    string
	Apellidos string
	Correo    *string
	Username  string
	Password  string
}

type Usuario struct {
	Id        string
	Nombre    string
	Apellidos string
	Correo    *string
	Username  string
	Password  string
}

func parseRow(r *sql.Row, t *Usuario) error {
	return r.Scan(
		&t.Id,
		&t.Nombre,
		&t.Apellidos,
		&t.Correo,
		&t.Username,
		&t.Password,
	)
}

func parseRows(r *sql.Rows, t *Usuario) error {
	return r.Scan(
		&t.Id,
		&t.Nombre,
		&t.Apellidos,
		&t.Correo,
		&t.Username,
		&t.Password,
	)
}
