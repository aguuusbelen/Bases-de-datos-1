package funciones

import (
    "database/sql"
    "fmt"
    _"time"
    _ "github.com/lib/pq"
    "log"
	"io/ioutil"
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
	
	_, err = db.Exec(`drop database if exists turnos_medicos;`)
	if err != nil {
		log.Fatal(err)
		fmt.Println("Error al eliminar la base si ya existia")
	}

	_, err = db.Exec(`create database turnos_medicos;`)
	if err != nil {
		log.Fatal(err)
		fmt.Println("Error al crear la base prueba")
	}
	
}

func BorrarBase(){
	db, err := sql.Open("postgres", "user=postgres host=localhost dbname=postgres sslmode=disable")
	defer db.Close()
	_, err = db.Exec(`drop database if exists turnos_medicos;`)
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
	
	var err error
	_, err = db.Query(`select generar_turnos_disponibles($1,$2);`,anio,mes)  
	if err != nil {
		log.Fatal(err)
		fmt.Println("Error al generar los turnos del mes ")
	}
	db.Close()
}

func AtenderTurnos_Dia(){ 
	db := conexionBase()
	defer db.Close()
	rows, err := db.Query(`select * from turno where estado='reservado'`) 
	if err != nil {
		log.Fatal(err)
		fmt.Println("Error")
	}
	defer rows.Close()
	
	var t Turno 
	
	for rows.Next(){
		if err := rows.Scan(&t.Nro_turno, &t.Fecha,&t.Nro_consultorio,&t.Dni_medique,&t.Nro_paciente,&t.Nro_obra_social_consulta,&t.Nro_afiliade_consulta,&t.Monto_paciente,&t.Monto_obra_social,&t.F_reserva,&t.Estado); 
		err != nil {
			log.Fatal(err)
		}
		_, err = db.Query(`select atencion_de_turno($1);`,t.Nro_turno) 
	}
	if err = rows.Err(); 
	err != nil {
		log.Fatal(err)
	}
}

func Liquidar_obra_social (anio, mes, nro_OS int) {
	db:= conexionBase()
	defer db.Close()
	_, err := db.Query(`select liquidacion_para_obras_sociales($1, $2, $3);`,anio,mes,nro_OS)
	if err != nil {
		log.Fatal(err)
		fmt.Println("Error al liquidar obra social")
	}
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
	db, err := sql.Open("postgres", "user=postgres host=localhost dbname=turnos_medicos sslmode=disable")
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



