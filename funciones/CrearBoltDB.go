package funciones

import (
    "fmt"
    _"time"
    _ "github.com/lib/pq"
    "log"
	"encoding/json"
	bolt "go.etcd.io/bbolt"
	"strconv"
)   

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
	Nro_consultorio int `json:"nro_consultorio"`
	Dni_medique int `json:"dni_medique"`
	Nro_paciente int `json:"nro_paciente"`
	Nro_obra_social_consulta any `json:"nro_obra_social_consulta, omitempty"`
	Nro_afiliade_consulta any `json:"nro_afiliade_consulta, omitempty"`
	Monto_paciente float64 `json:"monto_paciente"`
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
	
	turno1 := Turno {1, "2023-06-15 12:00", 5, 31759846, 1, 100, 1001, 0, 514.5, "2023-06-14 22:00","reservado"}
	
	data,err:= json.MarshalIndent(turno1, "", "     ")
	if err!=nil{
		log.Fatal(err)
	}
	
	CreateUpdate(dbbolt,"turno", []byte(strconv.Itoa(turno1.Nro_turno)) ,data)
	t1,err:=ReadUnique(dbbolt,"turno",[]byte(strconv.Itoa(turno1.Nro_turno)))
	fmt.Printf("%s\n",t1)
	
	turno2 := Turno {2, "2023-06-15 17:20", 5, 31759846, 6, 300, 3001, 200, 314.5, "2023-06-14 22:00","reservado"}
  
	data,err= json.MarshalIndent(turno2, "", "     ")
	if err!=nil{
		log.Fatal(err)
	}
	
	CreateUpdate(dbbolt,"turno", []byte(strconv.Itoa(turno2.Nro_turno)) ,data)
	t2,err:=ReadUnique(dbbolt,"turno",[]byte(strconv.Itoa(turno2.Nro_turno)))
	fmt.Printf("%s\n",t2)
	
	turno3 := Turno {3, "2023-06-15 17:20", 2, 20147852, 5, 200, 2002, 1000, 3347, "2023-06-14 22:00","reservado"}
  
	data,err= json.MarshalIndent(turno3, "", "     ")
	if err!=nil{
		log.Fatal(err)
	}
	
	CreateUpdate(dbbolt,"turno", []byte(strconv.Itoa(turno3.Nro_turno)) ,data)
	t3,err:=ReadUnique(dbbolt,"turno",[]byte(strconv.Itoa(turno3.Nro_turno)))
	fmt.Printf("%s\n",t3)
	
	turno4 := Turno {4, "2023-06-22 10:20", 2, 20147852, 14, 200, 2004, 1000, 3347, "2023-06-14 22:00","reservado"}
  
	data,err= json.MarshalIndent(turno4, "", "     ")
	if err!=nil{
		log.Fatal(err)
	}
	
	CreateUpdate(dbbolt,"turno", []byte(strconv.Itoa(turno4.Nro_turno)) ,data)
	t4,err:=ReadUnique(dbbolt,"turno",[]byte(strconv.Itoa(turno4.Nro_turno)))
	fmt.Printf("%s\n",t4)
	
	turno5 := Turno {5, "2023-06-16 13:00", 6, 29541019, 8, 400, 4002, 600, 368.8, "2023-06-14 22:00","reservado"}
	  data,err= json.MarshalIndent(turno5, "", "     ")
	if err!=nil{
		log.Fatal(err)
	}
	
	CreateUpdate(dbbolt,"turno", []byte(strconv.Itoa(turno5.Nro_turno)) ,data)
	t5,err:=ReadUnique(dbbolt,"turno",[]byte(strconv.Itoa(turno5.Nro_turno)))
	fmt.Printf("%s\n",t5)
	
	turno6 := Turno {6, "2023-06-16 13:00", 1, 31759846, 13, 100, 1003, 0, 514.5, "2023-06-14 22:00","reservado"}
	data,err= json.MarshalIndent(turno6, "", "     ")
	if err!=nil{
		log.Fatal(err)
	}
	
	CreateUpdate(dbbolt,"turno", []byte(strconv.Itoa(turno6.Nro_turno)) ,data)
	t6,err:=ReadUnique(dbbolt,"turno",[]byte(strconv.Itoa(turno6.Nro_turno)))
	fmt.Printf("%s\n",t6)
	
	turno7 := Turno {7, "2023-06-19 15:15", 3, 30668951, 4, 400, 4001, 400, 623.5, "2023-06-14 22:00","reservado"}
	data,err= json.MarshalIndent(turno7, "", "     ")
	if err!=nil{
		log.Fatal(err)
	}
	
	CreateUpdate(dbbolt,"turno", []byte(strconv.Itoa(turno7.Nro_turno)) ,data)
	t7,err:=ReadUnique(dbbolt,"turno",[]byte(strconv.Itoa(turno7.Nro_turno)))
	fmt.Printf("%s\n",t7)
	
	turno8 := Turno {8, "2023-06-21 12:00", 3, 30668951, 15, 100, 1004, 0, 1023.5, "2023-06-14 22:00","reservado"}
	data,err= json.MarshalIndent(turno8, "", "     ")
	if err!=nil{
		log.Fatal(err)
	}
	
	CreateUpdate(dbbolt,"turno", []byte(strconv.Itoa(turno8.Nro_turno)) ,data)
	t8,err:=ReadUnique(dbbolt,"turno",[]byte(strconv.Itoa(turno8.Nro_turno)))
	fmt.Printf("%s\n",t8)
	
	turno9 := Turno {9, "2023-06-20 10:00", 7, 20147852, 3, 200, 1004, 1000, 3347, "2023-06-14 22:00","reservado"}
	data,err= json.MarshalIndent(turno9, "", "     ")
	if err!=nil{
		log.Fatal(err)
	}
	
	CreateUpdate(dbbolt,"turno", []byte(strconv.Itoa(turno9.Nro_turno)) ,data)
	t9,err:=ReadUnique(dbbolt,"turno",[]byte(strconv.Itoa(turno9.Nro_turno)))
	fmt.Printf("%s\n",t9)
	
	turno10 := Turno {10, "2023-06-16 10:00", 6, 29541019, 8, 200, 4002, 600, 368.8, "2023-06-14 22:00","reservado"}
	data,err= json.MarshalIndent(turno10, "", "     ")
	if err!=nil{
		log.Fatal(err)
	}
	
	CreateUpdate(dbbolt,"turno", []byte(strconv.Itoa(turno10.Nro_turno)) ,data)
	t10,err:=ReadUnique(dbbolt,"turno",[]byte(strconv.Itoa(turno10.Nro_turno)))
	fmt.Printf("%s\n",t10)
	
	turno11 := Turno {11, "2023-06-16 10:30", 6, 29541019, 16, nil, nil, 600, 0, "2023-06-14 22:00","reservado"}
	data,err= json.MarshalIndent(turno11, "", "     ")
	if err!=nil{
		log.Fatal(err)
	}
	
	CreateUpdate(dbbolt,"turno", []byte(strconv.Itoa(turno11.Nro_turno)) ,data)
	t11,err:=ReadUnique(dbbolt,"turno",[]byte(strconv.Itoa(turno11.Nro_turno)))
	fmt.Printf("%s\n",t11)
	
	turno12 := Turno {12, "2023-06-21 09:00", 3, 30668951, 12, 100, 1002, 0, 1023.5, "2023-06-14 22:00","reservado"}
	data,err= json.MarshalIndent(turno12, "", "     ")
	if err!=nil{
		log.Fatal(err)
	}
	
	CreateUpdate(dbbolt,"turno", []byte(strconv.Itoa(turno12.Nro_turno)) ,data)
	t12,err:=ReadUnique(dbbolt,"turno",[]byte(strconv.Itoa(turno12.Nro_turno)))
	fmt.Printf("%s\n",t12)
	
	turno13 := Turno {13, "2023-06-15 17:00", 4, 25013965, 9, 300, 3002, 561.23, 400, "2023-06-14 22:00","reservado"}
	data,err= json.MarshalIndent(turno13, "", "     ")
	if err!=nil{
		log.Fatal(err)
	}
	
	CreateUpdate(dbbolt,"turno", []byte(strconv.Itoa(turno13.Nro_turno)) ,data)
	t13,err:=ReadUnique(dbbolt,"turno",[]byte(strconv.Itoa(turno13.Nro_turno)))
	fmt.Printf("%s\n",t13)
	
	turno14:= Turno{14,"2023-06-15 17:15", 4, 25013965, 18, 300, 3003, 561.23, 400, "2023-06-14 22:00","reservado"}
	data,err= json.MarshalIndent(turno14, "", "     ")
	if err!=nil{
		log.Fatal(err)
	}
	
	CreateUpdate(dbbolt,"turno", []byte(strconv.Itoa(turno14.Nro_turno)) ,data)
	t14,err:=ReadUnique(dbbolt,"turno",[]byte(strconv.Itoa(turno14.Nro_turno)))
	fmt.Printf("%s\n",t14)
	
	turno15:= Turno{15,"2023-06-15 17:30", 4, 25013965, 19, 300, 3004, 561.23, 400, "2023-06-14 22:00","reservado"}
	data,err= json.MarshalIndent(turno15, "", "     ")
	if err!=nil{
		log.Fatal(err)
	}
	
	CreateUpdate(dbbolt,"turno", []byte(strconv.Itoa(turno15.Nro_turno)) ,data)
	t15,err:=ReadUnique(dbbolt,"turno",[]byte(strconv.Itoa(turno15.Nro_turno)))
	fmt.Printf("%s\n",t15)
	
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
