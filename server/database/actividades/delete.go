package actividades

import "database/sql"

func Eliminar(db *sql.DB, id string) (string, error) {
	query := `delete from actividades where id = ?;`
	_, err := db.Exec(query, id)
	if err != nil {
		return "", err
	}

	return "ok", nil
}
