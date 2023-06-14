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
    
    mostrarOpciones()
    var seleccion string
    fmt.Scanf("%s",&seleccion)
    ejecutarOpcion(seleccion)
    
}

func ejecutarOpcion(selec string){    
    for selec !="q"{
		switch selec{
			case "1":
				fmt.Printf("Creando base de datos\n")
				f.CrearBase()
				break
			case "2":
				fmt.Printf("Creando tablas\n")
				f.CrearTablas()
				break
			case "3":
				fmt.Printf("Cargando primary keys y foreign keys\n")
				f.CargarKeys()
				break
			case "4":
				fmt.Printf("Cargando datos\n")
				f.CargarDatos()
				break
			case "5":
				fmt.Printf("Cargando funciones\n")
				f.CargarFunciones()
				break
			case "6":
				fmt.Printf("Generando turnos\n")
				f.GenerarTurnosDisponibles_Mes(2023,6)
				break
			case "7":
				fmt.Printf("Testeando\n")
				f.TestearConTabla()
				break
			case "8":
				fmt.Printf("Eliminando keys\n")
				f.BorrarKeys()
				break
			case "9":
				//fmt.Printf("Cargando BoltDB\n")
				//cargarBaseBoltDB()
<<<<<<< HEAD
				fmt.Printf("Creando Bolt DB\n")
				f.CrearBoltDB()
=======
				f.AtenderTurnos_Dia()
>>>>>>> 839021e911bb02d6859fb99fc5249c5b783c3b78
				break
			case "10":
				fmt.Printf("Eliminando base de datos\n")
				f.BorrarBase()
				break
			default:
				fmt.Printf("La opción elegida no es válida\n")
				
		}
		fmt.Printf("Elija otra opción\n")
		fmt.Scanf("%s",&selec)
	}
	
<<<<<<< HEAD
	fmt.Printf("Adios\n")
    
}



=======
	fmt.Printf("Adios. Gracias por utilizar el sistema!\n")
      
}

func mostrarOpciones() {
	fmt.Print("\n Introduzca la opción que desee realizar:\n")
	fmt.Print("1. Crear base de datos. \n")
	fmt.Print("2. Crear tablas. \n")
	fmt.Print("3. Agregar keys (primary and foreign). \n")
	fmt.Print("4. Cargar los datos.\n")
	fmt.Print("5. Crear stored procedures y Triggers.\n")
	fmt.Print("6. Generar turnos disponibles.\n")
	fmt.Print("7. Testear la base usando la tabla solicitud_reservas.\n")
	fmt.Print("8. Borrar keys (primary y foreign).\n")
	fmt.Print("9. Cargar base Bolt.dbq\n")
	fmt.Print("10. Borrar base de datos. \n")
	fmt.Print("q. Salir. \n\n")

	fmt.Print("Elija una opcion: \n")
}
>>>>>>> 839021e911bb02d6859fb99fc5249c5b783c3b78
