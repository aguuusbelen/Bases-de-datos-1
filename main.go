package main

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
    "log"
	"io/ioutil"
)   

var db *sql.DB

func main() { 
    crearBase()
    cargarDatos()
    cargarFunciones()
}

func crearBase() {
	db, err := sql.Open("postgres", "user=postgres host=localhost dbname=postgres sslmode=disable")
	if err !=nil {
		log.Fatal(err)
		fmt.Println("Error al abrir la base de datos creada")
	}
	defer db.Close()
	
	_, err = db.Exec(`drop database if exists prueba2;`)
	if err != nil {
		log.Fatal(err)
		fmt.Println("Error al eliminar la base si ya existia")
	}

	_, err = db.Exec(`create database prueba2;`)
	if err != nil {
		log.Fatal(err)
		fmt.Println("Error al crear la base prueba")
	}
	
}

func conexionBase() *sql.DB{
	db, err := sql.Open("postgres", "user=postgres host=localhost dbname=prueba2 sslmode=disable")
	if err !=nil {
		log.Fatal(err)
		fmt.Println("Error al abrir la base de datos ya creada")
	}
	//defer db.Close()
	
	return db
}

func cargarDatos() {
	db:= conexionBase()
	defer db.Close()
	ejecutar_sql(db, "creacion_tablas.sql")
	ejecutar_sql(db, "add_PKs_FKs.sql")
	ejecutar_sql(db, "carga_valores.sql")
	
}

func ejecutar_sql(db *sql.DB, path string){
	file, err := ioutil.ReadFile(path)
	
	if err !=nil {
		log.Fatal(err)
	}
	
	request := string(file)
	
	_, err = db.Exec(request)
	if err != nil {
		log.Fatal(err)
	}
}


func cargarFunciones() {
	db:= conexionBase()
	defer db.Close()
	ejecutar_sql(db, "generacion_de_turnos_disponibles.sql")
	ejecutar_sql(db, "reservar_turno.sql")
	ejecutar_sql(db, "atencion_de_turnos.sql")
	ejecutar_sql(db, "cancelar_turnos.sql")
	ejecutar_sql(db, "envio_mails.sql")
	ejecutar_sql(db, "liquidacion_para_obras_sociales.sql")
	
}

func borrarPKs_FKs () {
	db:= conexionBase()
	defer db.Close()
	ejecutar_sql(db, "borrarPK_FK.sql")
	
}
