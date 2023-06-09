package main

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
    "log"
	"io/ioutil"
)   

func main() {
    crearBase()
    cargarDatos()
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

func cargarDatos() {
	db, err := sql.Open("postgres", "user=postgres host=localhost dbname=prueba2 sslmode=disable")
	if err != nil {
		fmt.Println("Error al abrir la base de datos ya creada")
		log.Fatal(err)
	}
	defer db.Close()

	cargarTablas(db, "creacion_tablas.sql")
	cargarTablas(db, "add_PKs_FKs.sql")
	cargarTablas(db, "carga_valores.sql")
	
}


func cargarTablas(db *sql.DB, path string){
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
