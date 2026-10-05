package actividades

import "database/sql"

func Get(db *sql.DB, id string) (*Actividad, error) {
	query := `
	select 
		id,
		nombre
		from actividades
		where id = ? 
	`
	row := db.QueryRow(query, id)
	actividad := Actividad{}
	err := parseRow(row, &actividad)
	if err != nil {
		return nil, err
	}

	return &actividad, nil
}

func Listar(db *sql.DB) ([]*Actividad, error) {
	query := `
	select 
		id,
		nombre
		from actividades
	`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	rs := []*Actividad{}
	for rows.Next() {
		actividad := Actividad{}
		err := parseRows(rows, &actividad)
		if err != nil {
			return nil, err
		}

		rs = append(rs, &actividad)

	}

	return rs, nil
}
