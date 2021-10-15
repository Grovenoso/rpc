package main

import "fmt"

func main() {
	// -- mapa de materias --
	materias := make(map[string]map[string]float64)

	// creación de un alumno - calificacion
	alumno := make(map[string]float64)
	alumno["michel"] = 90

	fmt.Println(materias)

	// creacion de una materia
	materias["Algoritmia"] = alumno
	fmt.Println(materias)
	alumno["grover"] = 20
	materias["Algoritmia"] = alumno
	fmt.Println(materias)

	// agregar alumno a materia existente
	materias["Algoritmia"]["davalos"] = 95.5
	fmt.Println(materias)

	// consultar calificacion de un alumno para una materia
	fmt.Println(materias["Algoritmia"]["michel"])

	// mostrar el mapa de materias
	fmt.Println(materias)

	// mostrar los alumnos de una materia existente
	for alumno, calificion := range materias["Algoritmia"] {
		fmt.Println(alumno+":", calificion)
	}

	// -- mapa de alumnos --
	alumnos := make(map[string]map[string]float64)

	// creacion materia - calificion
	materia := make(map[string]float64)
	materia["Distribuidos"] = 100

	// creación de un alumno
	alumnos["boites"] = materia

	// agregar materia a un alumno existente
	alumnos["boites"]["Algoritmia"] = 60

	// consultar calificacion de una materia para un alumno y materia existente
	fmt.Println(alumnos["boites"]["Distribuidos"])

	// consultar las calificaciones de un alumno existente
	for materia, calificacion := range alumnos["boites"] {
		fmt.Println(materia+":", calificacion)
	}
}
