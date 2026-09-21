package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LoginRequest struct {
	Usuario  string `json:"usuario1"`
	Password string `json:"password1"`
}

func main() {

	// GET
	http.HandleFunc("/hola", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hola desde el servidor Go")
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

		// fmt.Fprintln(w, "Login recibido correctamente")

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
