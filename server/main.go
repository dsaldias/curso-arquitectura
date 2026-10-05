package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"gestor_tareas/database/actividades"
	"gestor_tareas/database/usuarios"
	"gestor_tareas/ejemplos"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type LoginRequest struct {
	Usuario  string `json:"usuario1"`
	Password string `json:"password1"`
}

func conexion() *sql.DB {
	dsn := "root:S1nclave@tcp(localhost:3306)/gestor_tareas?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	// defer db.Close()

	// Verificar que realmente existe conexión
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conexión a MySQL exitosa")

	return db
}

type Usuario struct {
	Id        string
	Nombre    string
	Apellidos string
	Correo    *string
	Username  string
	Password  string
}

func ListarUsuarios(db *sql.DB) {
	query := `
	select 
		id,
		nombre,
		apellidos,
		correo,
		username,
		password
		from usuarios;
	`

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	defer rows.Close()

	/* for indice, objeto := range []string{} {

	} */

	for rows.Next() {

		u := Usuario{}

		rows.Scan(
			&u.Id,
			&u.Nombre,
			&u.Apellidos,
			&u.Correo,
			&u.Username,
			&u.Password,
		)
		fmt.Println(">>>>>", u)
		fmt.Printf("%+v\n", u)
	}

}

func VerificarAcceso(db *sql.DB, username string, password string) (string, error) {
	query := `
	select 
		id,
		nombre,
		apellidos,
		correo,
		username,
		password
		from usuarios
		where username = ? and password = sha2(?,256);
	`

	row := db.QueryRow(query, username, password)

	user := Usuario{}
	err := row.Scan(
		&user.Id,
		&user.Nombre,
		&user.Apellidos,
		&user.Correo,
		&user.Username,
		&user.Password,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			t := "no hay usuarios"
			return "", errors.New(t)
		}
		return "", err
	}

	return "Acceso concedido", nil
}

func main() {

	db := conexion()

	ListarUsuarios(db)

	// GET
	http.HandleFunc("/hola", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		xn := r.URL.Query().Get("xnombre")
		xa := r.URL.Query().Get("apellidos")

		fmt.Println(">>>", xn, xa)

		resultado := ejemplos.Xsaludo()

		resul := map[string]any{
			"nombre":    "Diego",
			"telefonos": 787,
			"saludo":    resultado,
		}

		json.NewEncoder(w).Encode(&resul)

	})

	http.HandleFunc("/usuarios", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		listado, err := usuarios.Listar(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		json.NewEncoder(w).Encode(&listado)

	})

	http.HandleFunc("/new-usuario", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")

		// Preflight
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		if r.Method != http.MethodPost {
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
			return
		}

		// Recibir datos enviados desde el frontend
		input := usuarios.NewUsuario{}

		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			http.Error(w, "JSON inválido: "+err.Error(), http.StatusBadRequest)
			return
		}

		// Crear usuario
		us, err := usuarios.Crear(db, input)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Responder al frontend
		json.NewEncoder(w).Encode(us)
	})

	http.HandleFunc("/actualizar-usuario", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")

		// Preflight
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		if r.Method != http.MethodPut {
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
			return
		}

		// Recibir datos enviados desde el frontend
		input := usuarios.NewUsuario{}

		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			http.Error(w, "JSON inválido: "+err.Error(), http.StatusBadRequest)
			return
		}

		// Actualizar usuario
		us, err := usuarios.Actualizar(db, input)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Responder al frontend
		json.NewEncoder(w).Encode(us)
	})

	http.HandleFunc("/eliminar-usuario", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		xn := r.URL.Query().Get("idusuario")

		respuesta, err := usuarios.Eliminar(db, xn)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		resul := map[string]any{
			"resp": respuesta,
		}

		json.NewEncoder(w).Encode(&resul)

	})

	http.HandleFunc("/actividades", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		listado, err := actividades.Listar(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		json.NewEncoder(w).Encode(&listado)

	})

	http.HandleFunc("/new-actividad", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")

		// Preflight
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		if r.Method != http.MethodPost {
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
			return
		}

		// Recibir datos enviados desde el frontend
		input := actividades.NewActividad{}

		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			http.Error(w, "JSON inválido: "+err.Error(), http.StatusBadRequest)
			return
		}

		// Crear actividad
		act, err := actividades.Crear(db, input)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Responder al frontend
		json.NewEncoder(w).Encode(act)
	})

	http.HandleFunc("/actualizar-actividad", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")

		// Preflight
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		if r.Method != http.MethodPut {
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
			return
		}

		// Recibir datos enviados desde el frontend
		input := actividades.NewActividad{}

		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			http.Error(w, "JSON inválido: "+err.Error(), http.StatusBadRequest)
			return
		}

		// Actualizar actividad
		act, err := actividades.Actualizar(db, input)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Responder al frontend
		json.NewEncoder(w).Encode(act)
	})

	http.HandleFunc("/eliminar-actividad", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		xn := r.URL.Query().Get("idactividad")

		respuesta, err := actividades.Eliminar(db, xn)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		resul := map[string]any{
			"resp": respuesta,
		}

		json.NewEncoder(w).Encode(&resul)

	})

	// POST
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {

		// Permitir solicitudes desde otros puertos
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")

		// El navegador puede enviar primero una solicitud OPTIONS
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method != "POST" {
			fmt.Fprintln(w, "Metodo no permitido")
			return
		}

		var login LoginRequest

		err := json.NewDecoder(r.Body).Decode(&login)

		if err != nil {
			http.Error(w, "Body invalido", http.StatusBadRequest)
			return
		}

		fmt.Println("Usuario:", login.Usuario)
		fmt.Println("Password:", login.Password)

		_, err = VerificarAcceso(db, login.Usuario, login.Password)
		if err != nil {
			fmt.Println("el error de la db es:", err)
			http.Error(w, "usuario o clave incorrectos", http.StatusBadRequest)
		} else {

			// fmt.Println("el resultado es:", resultado)
			// if login.Usuario == "admin" && login.Password == "123456" {
			fmt.Fprintln(w, "Acceso concedido")
			return

		}
		/* else {
			http.Error(w, "usuario o clave incorrectos", http.StatusBadRequest)
			return
		} */

	})

	fmt.Println("Servidor iniciado en http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}
