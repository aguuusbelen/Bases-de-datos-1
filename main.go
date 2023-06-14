package main

import (
    _"database/sql"
    "fmt"
    _ "github.com/lib/pq"
    _"log"
	_"io/ioutil"
	f "buffagna-curcio-mendez-tp/funciones"
)   



func main() { 
    
    fmt.Printf("Introduzca la opción que desee realizar:\n1. Crear las bases de datos.\n2. Crear tablas. \n3. Agregar keys (primary y foreign).\n4. Cargar los datos. \n5. Cargar stored procedures y triggers.\n6. Generar turnos disponibles. \n7. Testear la base usando la tabla solicitud_reservas.\n8. Borrar keys (primary y foreign).\n9. Cargar base Bolt.dbq\n10. Borrar base de datos. \nq. Salir\n")
    var seleccion string
    fmt.Scanf("%s",&seleccion)
    
    for seleccion !="q"{
		switch seleccion{
			case "1":
				fmt.Printf("Creando base de datos")
				f.CrearBase()
				break
			case "2":
				fmt.Printf("Creando tablas")
				f.CrearTablas()
				break
			case "3":
				fmt.Printf("Cargando primary keys y foreign keys")
				f.CargarKeys()
				break
			case "4":
				fmt.Printf("Cargando datos")
				f.CargarDatos()
				break
			case "5":
				fmt.Printf("Cargando funciones")
				f.CargarFunciones()
				break
			case "6":
				fmt.Printf("Generando turnos")
				f.GenerarTurnosDisponibles_Mes(2023,6)
				break
			case "7":
				fmt.Printf("Testeando")
				f.TestearConTabla()
				break
			case "8":
				fmt.Printf("Eliminando keys")
				f.BorrarKeys()
				break
			case "9":
				//cargarBaseBoltDB()
				break
			case "10":
				f.BorrarBase()
				break
			default:
				fmt.Printf("La opción elegida no es válida\n")
				
		}
		fmt.Printf("Elija otra opción\n")
		fmt.Scanf("%s",&seleccion)
	}
	
	fmt.Printf("Adios\n")
    
    
    
}

