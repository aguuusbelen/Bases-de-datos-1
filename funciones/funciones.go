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

func AtenderTurnos_Dia(){ 
	db := conexionBase()
	defer db.Close()
	//_var err error
	rows, err := db.Query(`select * from turno where estado='reservado'`) 
	if err != nil {
		log.Fatal(err)
		fmt.Println("Error")
	}
	defer rows.Close()
	
	var t turno 
	
	for rows.Next(){
		if err := rows.Scan(&t.nro_turno, &t.fecha,&t.nro_consultorio,&t.dni_medique,&t.nro_paciente,&t.nro_obra_social_consulta,&t.nro_afiliade_consulta,&t.monto_paciente,&t.monto_obra_social,&t.f_reserva,&t.estado); 
		err != nil {
			log.Fatal(err)
		}
		_, err = db.Query(`select atencion_de_turno($1);`,t.nro_turno) 
	}
	if err = rows.Err(); 
	err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("Atencion de turnos del dìa \n")
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

func EnvioMailsDiarios(){
	db:= conexionBase()
	defer db.Close()
	_, err := db.Query(`select envio_mail_diario()`)
	if err != nil {
		log.Fatal(err)
		fmt.Println("Error al enviar mails")
	}
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

type Medique struct{
	Dni_medique int `json:"dni_medique"`
	Nombre string `json:"nombre"`
	Apellido string `json:"apellido"`
	Especialidad any `json:"especialidad,omitempty"`
	Monto_consulta_privada string `json:"monto_consulta_privada"`
	Telefono string `json:telefono"`
}

type Consultorio struct{
	Nro_consultorio int `json:"nro_consultorio"`
	Nombre string `json:"nombre"`
	Domicilio string `json:"domicilio"`
	Codigo_postal string `json:"codigo_postal"`
	Telefono string `json:telefono"`
}

type Obra_social struct{
	Nro_obra_social int `json:"nro_obra_social"`
	Nombre string `json:"nombre"`
	Contacto_nombre string `json:"contacto_nombre"`
	Contacto_apellido string `json:"contacto_apellido"`
	Contacto_telefono string `json:"contacto_telefono"`
	Contacto_email string `json:"contacto_email"`
}

type Turno struct{
	Nro_turno int `json:"nro_turno"`
	Fecha string `json:"fecha"`
	Nro_consultorio string `json:"nro_consultorio"`
	Dni_medique int `json:"dni_medique"`
	Nro_paciente int `json:"nro_paciente"`
	Nro_obra_social_consulta any `json:"nro_obra_social_consulta, omitempty"`
	Nro_afiliade_consulta any `json:"nro_afiliade_consulta, omitempty"`
	Monto_paciente string `json:"monto_paciente"`
	Monto_obra_social any `json:"monto_obra_social,omitempty"`
	F_reserva string `json:"f_reserva"`
	Estado string `json:"estado"`
}

func CrearBoltDB() {

	// Abrimos la BoltDB
	dbbolt, err := bolt.Open("bolt.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer dbbolt.Close()

	// Abrimos la bd psql 
	db:= conexionBase()
	defer db.Close()

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

		// transformamos al paciente al formato json
		data_pac, err:=json.MarshalIndent(pac, "", "     ")
		if err!=nil{
			log.Fatalf("%s",err)	
		}
		
		// se cargan los pacientes en la BoltDB
		CreateUpdate(dbbolt, "pacientes", []byte(strconv.Itoa(pac.Nro_paciente)), data_pac)
		resultado1, err := ReadUnique(dbbolt, "pacientes", []byte(strconv.Itoa(pac.Nro_paciente)))
		fmt.Printf("%s\n", resultado1)
	}
	
	// Obtenemos todos los mediques
	rows_mediques, err:= db.Query(`select * from medique`)
	if err!=nil{
		log.Fatal(err)
	}

	for rows_mediques.Next(){
		var med Medique
		if err:= rows_mediques.Scan(&med.Dni_medique, &med.Nombre, &med.Apellido, &med.Especialidad, &med.Monto_consulta_privada, &med.Telefono); err!=nil{
			log.Fatal(err)
		}

		// transformamos al medique al formato json
		data_med, err:=json.MarshalIndent(med, "", "     ")
		if err!=nil{
			log.Fatalf("%s",err)	
		}
		
		// se cargan los mediques en la BoltDB
		CreateUpdate(dbbolt, "mediques", []byte(strconv.Itoa(med.Dni_medique)), data_med)
		resultado2, err := ReadUnique(dbbolt, "mediques", []byte(strconv.Itoa(med.Dni_medique)))
		fmt.Printf("%s\n", resultado2)
	
	}
	
	// Obtenemos todos los consultorios
	rows_consultorios, err:= db.Query(`select * from consultorio`)
	if err!=nil{
		log.Fatal(err)
	}

	for rows_consultorios.Next(){
		var consul Consultorio
		if err:= rows_consultorios.Scan(&consul.Nro_consultorio, &consul.Nombre, &consul.Domicilio, &consul.Codigo_postal, &consul.Telefono); err!=nil{
			log.Fatal(err)
		}

		// transformamos al consultorio al formato json
		data_consul, err:=json.MarshalIndent(consul, "", "     ")
		if err!=nil{
			log.Fatalf("%s",err)	
		}
		
		// se cargan los consultorios en la BoltDB
		CreateUpdate(dbbolt, "consultorios", []byte(strconv.Itoa(consul.Nro_consultorio)), data_consul)
		resultado3, err := ReadUnique(dbbolt, "consultorios", []byte(strconv.Itoa(consul.Nro_consultorio)))
		fmt.Printf("%s\n", resultado3)
	
	}
	
	// Obtenemos todas las obras sociales
	rows_obras_sociales, err:= db.Query(`select * from obra_social`)
	if err!=nil{
		log.Fatal(err)
	}

	for rows_obras_sociales.Next(){
		var os Obra_social
		if err:= rows_obras_sociales.Scan(&os.Nro_obra_social, &os.Nombre, &os.Contacto_nombre, &os.Contacto_apellido, &os.Contacto_telefono, &os.Contacto_email); err!=nil{
			log.Fatal(err)
		}

		// transformamos al consultorio al formato json
		data_os, err:=json.MarshalIndent(os, "", "     ")
		if err!=nil{
			log.Fatalf("%s",err)	
		}
		
		// se cargan los consultorios en la BoltDB
		CreateUpdate(dbbolt, "obras_sociales", []byte(strconv.Itoa(os.Nro_obra_social)), data_os)
		resultado4, err := ReadUnique(dbbolt, "obras_sociales", []byte(strconv.Itoa(os.Nro_obra_social)))
		fmt.Printf("%s\n", resultado4)
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

type turno struct{
	nro_turno int
	fecha string  
	nro_consultorio int
	dni_medique int
	nro_paciente int
	nro_obra_social_consulta sql.NullInt64
	nro_afiliade_consulta sql.NullInt64
	monto_paciente float64 
	monto_obra_social float64
	f_reserva string
	estado string 
}
