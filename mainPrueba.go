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
    generarTurnosDisponibles()
    cargarFunciones()
    generarTurnosDisponibles_Mes(2023,6)
    //Liquidar_obra_social(2023,6,400)
}

func crearBase() {
	db, err := sql.Open("postgres", "user=postgres host=localhost dbname=postgres sslmode=disable")
	if err !=nil {
		log.Fatal(err)
		fmt.Println("Error al abrir la base de datos creada")
	}
	defer db.Close()
	
	_, err = db.Exec(`drop database if exists prueba3;`)
	if err != nil {
		log.Fatal(err)
		fmt.Println("Error al eliminar la base si ya existia")
	}

	_, err = db.Exec(`create database prueba3;`)
	if err != nil {
		log.Fatal(err)
		fmt.Println("Error al crear la base prueba")
	}
	
}

func conexionBase(){
	db, err := sql.Open("postgres", "user=postgres host=localhost dbname=prueba3 sslmode=disable")
	if err !=nil {
		log.Fatal(err)
		fmt.Println("Error al abrir la base de datos ya creada")
	}
	defer db.Close()
}

func cargarDatos() {
	db, err := sql.Open("postgres", "user=postgres host=localhost dbname=prueba3 sslmode=disable")
	if err !=nil {
		log.Fatal(err)
		fmt.Println("Error al abrir la base de datos ya creada")
	}
	defer db.Close()
	//conexionBase()
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

func generarTurnosDisponibles() {
	//conexionBase()
	db, err := sql.Open("postgres", "user=postgres host=localhost dbname=prueba3 sslmode=disable")
	if err !=nil {
		log.Fatal(err)
		fmt.Println("Error al abrir la base de datos ya creada")
	}
	defer db.Close()
	ejecutar_sql(db, "generacion_de_turnos_disponibles.sql")
	
}

func generarTurnosDisponibles_Mes(a int, m int){
	//conexionBase()
	db, err := sql.Open("postgres", "user=postgres host=localhost dbname=prueba3 sslmode=disable")
	if err !=nil {
		log.Fatal(err)
		fmt.Println("Error al abrir la base de datos ya creada")
	}
	defer db.Close()
	_, err = db.Query(`select generar_trunos_disponibles(&a, &m);`)  
	if err != nil {
		log.Fatal(err)
		fmt.Println("Error al generar los turnos del mes ")
	}
}

func cargarFunciones() {
	//conexionBase()
	db, err := sql.Open("postgres", "user=postgres host=localhost dbname=prueba3 sslmode=disable")
	if err !=nil {
		log.Fatal(err)
		fmt.Println("Error al abrir la base de datos ya creada")
	}
	defer db.Close()
	ejecutar_sql(db, "reservar_turno.sql")
	ejecutar_sql(db, "atencion_de_turnos.sql")
	ejecutar_sql(db, "cancelar_turnos.sql")
	ejecutar_sql(db, "envio_mails.sql")
	ejecutar_sql(db, "liquidacion_para_obras_sociales.sql")
	
}

/*func borrarPKs_FKs () {
	conexionBase()
	ejecutar_sql(db, "borrarPK_FK.sql")
	
}*/

/*func Liquidar_obra_social (anio, mes, OS int){
	//conexionBase()
	db, err := sql.Open("postgres", "user=postgres host=localhost dbname=prueba3 sslmode=disable")
	if err !=nil {
		log.Fatal(err)
		fmt.Println("Error al abrir la base de datos ya creada")
	}
	defer db.Close()
	_, err = db.Query(`select liquidacion_para_obras_sociales(anio, mes, OS);`)
	if err != nil {
		log.Fatal(err)
		fmt.Println("Error al liquidar obra social")
	}
	
}*/
