package ejemplos

import "fmt"

type Persona struct {
	Nombre   string // ""
	Telefono int    // 0
}

func Saludo() string {

	nombre := "Diego"
	apellido := "Saldias"

	telefono := 78227092

	telefonos := []string{"78227092", "46545454"}
	casas := []int{1, 2, 3, 4}

	persona := Persona{
		Nombre:   "Diego",
		Telefono: 78227092,
	}

	personas := []Persona{}

	personas = append(personas, persona, persona, persona)

	fmt.Println(nombre, apellido, telefono, telefonos, casas)
	fmt.Println(persona)

	fmt.Printf("%+v\n\n\n\n\n", persona)

	res := Suma(5, 8)

	fmt.Println("resultado=", res)

	s1, r1, m1 := Arimetica(5, 8)
	fmt.Println(s1, r1, m1)

	fmt.Println(personas)

	return "Hola mundo a todos"
}

func Suma(parametro1 int, parametro2 int) int {
	resultado := parametro1 + parametro2
	return resultado
}

func Arimetica(parametro1 int, parametro2 int) (int, int, int) {
	suma := parametro1 + parametro2
	resta := parametro1 - parametro2
	multiplicacion := parametro1 * parametro2

	return suma, resta, multiplicacion
}

func Xsaludo() string {
	return "Hola a todos"
}
