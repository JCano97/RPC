package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"strconv"
	"strings"
)

var materias = make(map[string]map[string]float64)
var alumnos = make(map[string]map[string]float64)

type Server struct{}

func (this *Server) Agregar(datos string, reply *string) error {
	arregloDatos := strings.Split(datos, ",")
	var calif float64
	calif, err := strconv.ParseFloat(arregloDatos[2], 64)
	if err != nil {
		log.Fatal(err)
	}
	alumnoNuevo := make(map[string]float64)
	alumnoNuevo[arregloDatos[0]] = calif
	if el, ok := materias[arregloDatos[1]]; ok {
		if _, ok := el[arregloDatos[0]]; ok {
			return errors.New("Error! ya existe calificacion para este alumno y materia especificos")
		}else{
			materias[arregloDatos[1]][arregloDatos[0]] = calif
		}
	} else {
		materias[arregloDatos[1]] = alumnoNuevo
	}
	materiaNueva := make(map[string]float64)
	materiaNueva[arregloDatos[1]] = calif
	if _, ok := alumnos[arregloDatos[0]]; ok {
		alumnos[arregloDatos[0]][arregloDatos[1]] = calif
	} else {
		alumnos[arregloDatos[0]] = materiaNueva
	}
	*reply = "Hecho "+ arregloDatos[0]
	return nil
}
func (this *Server) PromedioAlumno(nombreAlumno string, reply *float64) error {
	cont := 0.0
	suma := 0.0
	for _, calificacion := range alumnos[nombreAlumno] {
		cont++
		suma = suma+calificacion
	}
	promedio := suma/cont
	*reply = promedio
	return nil
}
func (this *Server) PromedioMateria(nombreMateria string, reply *float64) error {
	cont := 0.0
	suma := 0.0
	for _, calificacion := range materias[nombreMateria] {
		cont++
		suma = suma+calificacion
	}
	promedio := suma/cont
	*reply = promedio
	return nil
}
func (this *Server) PromedioGeneral(cont float64, reply *float64) error {
	suma := 0.0
	for _,alumnoCalif := range materias {
		
		for _,calificacion := range alumnoCalif {
			cont++
			suma = suma+calificacion
		}
	}
	promedio := suma/cont
	*reply = promedio
	return nil
}
func server() {
	rpc.Register(new(Server))
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go rpc.ServeConn(c)
	}
}

func main() {
	go server()

	var input string
	fmt.Scanln(&input)
}
