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
    
    fmt.Printf("Introduzca la opción que desee realizar:\n1. Crear las bases de datos.\n2. Crear tablas. \n3. Agregar keys (primary y foreign).\n4. Cargar los datos. \n5. Cargar stored procedures y triggers.\n6. Testear la base usando la tabla solicitud_reservas.\n7. Borrar keys (primary y foreign).\n8. Cargar base Bolt.dbq\nq. Salir\n")
    
    
    var seleccion string
    fmt.Scanf("%s",&seleccion)
    
    switch seleccion{
		case "1":
			crearBase()
			break
		case "2":
			crearTablas()
			break
		case "3":
			cargarKeys()
			break
		case "4":
			cargarDatos()
			break
		case "5":
			cargarFunciones()
			break
		case "6":
			//testearConTabla()
			break
		case "7":
			borrarKeys()
			break
		case "8":
			//cargarBaseBoltDB()
			break
		case "q":
			return
			break
		default:
			fmt.Printf("La opción elegida no es válida\n")
	}
    
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
func crearTablas() {
	db:= conexionBase()
	defer db.Close()
	ejecutar_sql(db, "creacion_tablas.sql")
}

func cargarKeys() {
	db:= conexionBase()
	defer db.Close()
	ejecutar_sql(db, "add_PKs_FKs.sql")
}
	
func cargarDatos() {
	db:= conexionBase()
	defer db.Close()
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

func borrarKeys() {
	db:= conexionBase()
	defer db.Close()
	ejecutar_sql(db, "borrarPK_FK.sql")
	
}
