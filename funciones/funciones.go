package funciones

import (
    "database/sql"
    "fmt"
    _"time"
    _ "github.com/lib/pq"
    "log"
	"io/ioutil"
	"encoding/json"
	bolt "go.etcd.io/bbolt"
	"strconv"
)   
//Variables
var db *sql.DB

//Funciones
func CrearBase() {
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

func BorrarBase(){
	db, err := sql.Open("postgres", "user=postgres host=localhost dbname=postgres sslmode=disable")
	defer db.Close()
	_, err = db.Query(`drop database if exists prueba2;`)
	if err != nil {
		log.Fatal(err)
		fmt.Println("Error al eliminar la base de datos")
	}
	
}

func CrearTablas() {
	db:= conexionBase()
	defer db.Close()
	ejecutar_sql(db, "archivosSQL/creacion_tablas.sql")
}

func CargarKeys() {
	db:= conexionBase()
	defer db.Close()
	ejecutar_sql(db, "archivosSQL/add_PKs_FKs.sql")
}
	
func CargarDatos() {
	db:= conexionBase()
	defer db.Close()
	ejecutar_sql(db, "archivosSQL/carga_valores.sql")
	
}

func CargarFunciones() {
	db:= conexionBase()
	defer db.Close()
	ejecutar_sql(db, "archivosSQL/generacion_de_turnos_disponibles.sql")
	ejecutar_sql(db, "archivosSQL/reservar_turno.sql")
	ejecutar_sql(db, "archivosSQL/atencion_de_turnos.sql")
	ejecutar_sql(db, "archivosSQL/cancelar_turnos.sql")
	ejecutar_sql(db, "archivosSQL/envio_mails.sql")
	ejecutar_sql(db, "archivosSQL/liquidacion_para_obras_sociales.sql")
	
}

func BorrarKeys() {
	db:= conexionBase()
	defer db.Close()
	ejecutar_sql(db, "archivosSQL/borrarPK_FK.sql")
	
}

func GenerarTurnosDisponibles_Mes(anio, mes int){ 
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

func TestearConTabla() {
	db:= conexionBase()
	defer db.Close()
	ejecutar_sql(db, "archivosSQL/test_reservas.sql")
}


//Funciones auxiliares
func conexionBase() *sql.DB{
	db, err := sql.Open("postgres", "user=postgres host=localhost dbname=prueba2 sslmode=disable")
	if err !=nil {
		log.Fatal(err)
		fmt.Println("Error al abrir la base de datos ya creada")
	}
	
	return db
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

type Paciente struct{
	Nro_paciente int `json:"nro_paciente"`
	Nombre string `json:"nombre"`
	Apellido string `json:"apellido"`
	Dni_paciente int `json:"dni_paciente"`
	F_nac string `json:"f_nac"`
	Nro_obra_social any `json:"nro_obra_social,omitempty"`
	Nro_afiliade any  `json:"nro_afiliade,omitempty"`
	Domicilio string `json:"domicilio"`
	Telefono string `json:"telefono"`
	Email string `json:"email"`
}

func CrearBoltDB() {

	// Abrimos la bolt bd
	dbbolt, err := bolt.Open("bolt.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer dbbolt.Close()

	// Abrimos la bd 
	db:= conexionBase()
	defer db.Close()

	// var lista_pac []Paciente

	// Obtenemos todos los pacientes
	rows_pacientes, err:= db.Query(`select * from paciente`)
	if err!=nil{
		log.Fatal(err)
	}

	for rows_pacientes.Next(){
		var pac Paciente
		if err:= rows_pacientes.Scan(&pac.Nro_paciente, &pac.Nombre, &pac.Apellido, &pac.Dni_paciente, &pac.F_nac, &pac.Nro_obra_social, &pac.Nro_afiliade, &pac.Domicilio, &pac.Telefono, &pac.Email); err!=nil{
			log.Fatal(err)
		}
		// lista_pac=append(lista_pac,pac)
		
		// transformamos al paciente al formato json
		data_pac, err:=json.MarshalIndent(pac, "", "     ")
		if err!=nil{
			log.Fatalf("%s",err)	
		}
		
		// se cargan los pacientes en la db bolt
		CreateUpdate(dbbolt, "pacientes", []byte(strconv.Itoa(pac.Nro_paciente)), data_pac)
		resultado1, err := ReadUnique(dbbolt, "pacientes", []byte(strconv.Itoa(pac.Nro_paciente)))
		fmt.Printf("%s\n", resultado1)
	}
}


func CreateUpdate(db *bolt.DB, bucketName string, key []byte, val []byte) error {
	// abre transacción de escritura
	tx, err := db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	b, _ := tx.CreateBucketIfNotExists([]byte(bucketName))
	err = b.Put(key, val)
	if err != nil {
		return err
	}
	// cierra transacción
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func ReadUnique(db *bolt.DB, bucketName string, key []byte) ([]byte, error) {
	var buf []byte
	// abre una transacción de lectura
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		buf = b.Get(key)
		return nil
	})
	return buf, err
}
