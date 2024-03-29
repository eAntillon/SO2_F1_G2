package Menu

import (
	"bufio"
	"fmt"
	iot "github.com/eAntillon/SO2_F1_G2/Iotop"
	rep "github.com/eAntillon/SO2_F1_G2/Reportes"
	top "github.com/eAntillon/SO2_F1_G2/Top"
	strace "github.com/eAntillon/SO2_F1_G2/strace"
	mem "github.com/eAntillon/SO2_F1_G2/Memsim"
	"os"
)

func Menu_principal() {

	colorReset := "\033[0m"

	colorRed := "\033[31m"
	//colorGreen := "\033[32m"
	colorYellow := "\033[33m"
	colorBlue := "\033[34m"
	//colorPurple := "\033[35m"
	//colorCyan := "\033[36m"
	//colorWhite := "\033[37m"

	fmt.Println(string(colorRed), "######## BIENVENIDOS AL SYSTEMA DE METAOS ############")
	for {
		fmt.Println(string(colorBlue), "|----------------------------------------------------|")
		fmt.Println(string(colorBlue), "|                    MENU                            |")
		fmt.Println(string(colorBlue), "|----------------------------------------------------|")
		fmt.Println(string(colorBlue), "|  1)  Nueva ejecucion                               |")
		fmt.Println(string(colorBlue), "|  2)  Reporte                                       |")
		fmt.Println(string(colorBlue), "|  3)  Exit                                          |")
		fmt.Println(string(colorBlue), "|----------------------------------------------------|")
		fmt.Println(string(colorBlue), "|  Seleccione una opcion:                            |")
		fmt.Println("")

		fmt.Print(string(colorYellow), ">> ", string(colorReset))
		com := bufio.NewScanner(os.Stdin)
		if com.Scan() {

			if com.Text() == "1" {
				nuevaEjecucion()
			} else if com.Text() == "2" {
				mostrarReportes()
			} else if com.Text() == "3" {
				break
			}
		}

	}
}

func nuevaEjecucion() {
	colorReset := "\033[0m"
	colorYellow := "\033[33m"
	for {
		fmt.Println("|----------------------------------------------------|")
		fmt.Println("|                SUB MENU                            |")
		fmt.Println("|----------------------------------------------------|")
		fmt.Println("|  1)  Ingrese su nombre                             |")
		fmt.Println("|  2)  Regresar                                      |")
		fmt.Println("|----------------------------------------------------|")
		fmt.Println("|  Seleccione una opcion:                            |")
		fmt.Println("")

		fmt.Print(string(colorYellow), ">> ", string(colorReset))
		com := bufio.NewScanner(os.Stdin)
		if com.Scan() {
			if com.Text() == "1" {
				monitoreo()
			} else if com.Text() == "2" {
				break
			}
		}
	}
}

func mostrarReportes() {
	colorReset := "\033[0m"
	colorYellow := "\033[33m"
	colorRed := "\033[31m"
	for {
		fmt.Println("|----------------------------------------------------|")
		fmt.Println("|                  Bitacora                          |")
		fmt.Println("|----------------------------------------------------|")
		fmt.Println("|  1)  Bitacora                                      |")
		fmt.Println("|  2)  Estado de Simulacion                          |")
		fmt.Println("|----------------------------------------------------|")
		fmt.Println("|  Seleccione la opcion:                            |")
		fmt.Println("")
		fmt.Print(string(colorYellow), ">> ", string(colorReset))
		com := bufio.NewScanner(os.Stdin)
		if com.Scan() {
			if com.Text() == "1" {
				Bitacora()
				break
			} else if com.Text() == "2" {
				Estado_Simulacion()
				break
			} else {
				fmt.Println(string(colorRed), " Opcion Incorrecta", string(colorReset))
				break
			}
		}
	}
}

func Bitacora() {
	fmt.Println("############## BITACORA ################")
	rep.EscribirBitacora()
}

func Estado_Simulacion() {

	fmt.Println("############## ESTADO DE SIMULACION ################")
	rep.EscribirSimulacion()
}

func monitoreo() {
	colorReset := "\033[0m"
	colorYellow := "\033[33m"
	colorGreen := "\033[32m"
	colorCyan := "\033[36m"
	fmt.Println(string(colorGreen), "  !Por favor ingrese su nombre o alias!", string(colorReset))
	fmt.Print(string(colorYellow), ">> ", string(colorReset))
	com := bufio.NewScanner(os.Stdin)
	com.Scan()
	nombre := com.Text()
	fmt.Println(" ########  ", nombre)
	rep.CrearUsuario(nombre)

	for {
		fmt.Println("")
		fmt.Println(string(colorCyan), "|----------------------------------------------------|")
		fmt.Println(string(colorCyan), "|  1)  IOTOP                                         |")
		fmt.Println(string(colorCyan), "|  2)  TOP                                           |")
		fmt.Println(string(colorCyan), "|  3)  SYSCALL                                       |")
		fmt.Println(string(colorCyan), "|  4)  MEMSIM                                        |")
		fmt.Println(string(colorCyan), "|  5)  Regresar                                      |")
		fmt.Println(string(colorCyan), "|----------------------------------------------------|")
		fmt.Println(string(colorCyan), "|  Seleccione una opcion:                            |")
		fmt.Println("")
		fmt.Print(string(colorYellow), ">> ", string(colorReset))

		com := bufio.NewScanner(os.Stdin)
		if com.Scan() {
			if com.Text() == "1" {
				rep.AgregarActividad("IOTOP")
				IOTOP(nombre)
			} else if com.Text() == "2" {
				rep.AgregarActividad("TOP")
				TOP(nombre)
			} else if com.Text() == "3" {
				rep.AgregarActividad("SYSCALL")
				SYSCALL(nombre)
			} else if com.Text() == "4" {
				MEMSIM(nombre)
			} else if com.Text() == "5" {
				rep.GuardarUsuario()
				break
			}

		}
	}

}
func IOTOP(name string) {
	fmt.Println("############## IOTOP ################")
	iot.Run(name)
}
func TOP(name string) {
	fmt.Println("############## TOP ################")
	top.Run(name)
}

func SYSCALL(name string) {
	fmt.Println("############## SYSCALL ################")
	strace.Run(name)
	/*
		fmt.Println(string(colorGreen),"  !Ingrese el comando !",string(colorReset))
		fmt.Print(string(colorYellow),">> ",string(colorReset))
			com := bufio.NewScanner(os.Stdin)
			com.Scan()
			comando:=com.Text()
			fmt.Println(" ########  ",comando)
	*/

}
var (
	arreglo       string
	arregloPartes []string
)

func MEMSIM(nombre string) {
	fmt.Println("############## MEMSIM  ################")
	mem.Run(nombre)
}
