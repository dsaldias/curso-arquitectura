package actividades

import "database/sql"

type NewActividad struct {
	Id     *string
	Nombre string
}

type Actividad struct {
	Id     string
	Nombre string
}

func parseRow(r *sql.Row, t *Actividad) error {
	return r.Scan(
		&t.Id,
		&t.Nombre,
	)
}

func parseRows(r *sql.Rows, t *Actividad) error {
	return r.Scan(
		&t.Id,
		&t.Nombre,
	)
}
