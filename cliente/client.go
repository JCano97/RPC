package main

import (
	"fmt"
	"net/rpc"
)

func client() {
	c, err := rpc.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	var opc int64
	for {
		fmt.Println("-----------------------------------")
		fmt.Println("1.- Agregar Calificación")
		fmt.Println("2.- Mostrar promedio de alumno")
		fmt.Println("3.- Mostrar promedio general")
		fmt.Println("4.- Mostrar promedio de manteria")
		fmt.Println("0.- Salir")
		fmt.Print("Opción: ")
		fmt.Scanln(&opc)

		switch opc {
		case 1:
			var nombre string
			var materia string
			var calif string
			fmt.Print("Nombre del Alumno: ")
			fmt.Scanln(&nombre)
			fmt.Print("Materia: ")
			fmt.Scanln(&materia)
			fmt.Print("Calificación: ")
			fmt.Scanln(&calif)
			var datos string
			datos = nombre + "," + materia + "," + calif
			var result string
			err = c.Call("Server.Agregar",datos,&result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Server.Agregar", result)
			}
		case 2:
			var nombreAlumno string
			fmt.Print("Nombre del alumno:")
			fmt.Scanln(&nombreAlumno)
			var result float64
			err = c.Call("Server.PromedioAlumno",nombreAlumno,&result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Server.PromedioAlumno: ", result)
			}
		case 3:
			var result float64
			err = c.Call("Server.PromedioGeneral",0.0,&result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Server.PromedioGeneral: ", result)
			}
		case 4:
			var nombreMateria string
			fmt.Print("Nombre de la materia:")
			fmt.Scanln(&nombreMateria)
			var result float64
			err = c.Call("Server.PromedioMateria",nombreMateria,&result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Server.PromedioMateria: ", result)
			}
		case 0:
			return
		}
	}
}

func main() {
	client()
}
