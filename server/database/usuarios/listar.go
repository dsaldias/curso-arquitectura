package usuarios

import "database/sql"

func Get(db *sql.DB, id string) (*Usuario, error) {
	query := `
	select 
		id,
		nombre,
		apellidos,
		correo,
		username,
		password
		from usuarios
		where id = ? 
	`
	row := db.QueryRow(query, id)
	user := Usuario{}
	err := parseRow(row, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func Listar(db *sql.DB) ([]*Usuario, error) {
	query := `
	select 
		id,
		nombre,
		apellidos,
		correo,
		username,
		password
		from usuarios
		where estado = 1
	`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	rs := []*Usuario{}
	for rows.Next() {
		user := Usuario{}
		err := parseRows(rows, &user)
		if err != nil {
			return nil, err
		}

		rs = append(rs, &user)

	}

	return rs, nil
}
