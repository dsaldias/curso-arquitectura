package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
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

func VerificarAcceso(db *sql.DB, username string, password string) {
	/* select
	id,
	nombre,
	apellidos,
	correo,
	username,
	password
	from usuarios
	where username = ? and password = ?; */
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

		if login.Usuario == "admin" && login.Password == "123456" {
			fmt.Fprintln(w, "Acceso concedido")
			return

		} else {
			http.Error(w, "usuario o clave incorrectos", http.StatusBadRequest)
			return
		}

	})

	fmt.Println("Servidor iniciado en http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}
