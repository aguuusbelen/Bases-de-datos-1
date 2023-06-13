package main

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
    "log"
	"io/ioutil"
)   


func generarTurnosDisponibles_Mes(anio, mes int){
	db := conexionBase()
	defer db.Close()
	var err error
	_, err = db.Query(`select generar_turnos_disponibles($1,$2);`,anio,mes)  
	if err != nil {
		log.Fatal(err)
		fmt.Println("Error al generar los turnos del mes ")
	}
	fmt.Printf("Creacion turnos %d-%d\n",anio,mes)
}

func Liquidar_obra_social (anio, mes, nro_OS int) {
	db:= conexionBase()
	defer db.Close()
	_, err := db.Query(`select liquidacion_para_obras_sociales($1, $2, $3);`,anio,mes,nro_OS)
	if err != nil {
		log.Fatal(err)
		fmt.Println("Error al liquidar obra social")
	}
	fmt.Printf("Liquidacion de %d-%d\n",anio,mes)
}


