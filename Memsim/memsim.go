package Memsim

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
	rep "github.com/eAntillon/SO2_F1_G2/Reportes"
)

const (
	InfoColor    = "\033[1;32m%s\033[0m"
	NoticeColor  = "\033[1;36m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
	DebugColor   = "\033[0;36m%s\033[0m"
	colorReset = "\033[0m"
	colorYellow = "\033[33m"
	colorGreen = "\033[32m"
	colorBlue = "\033[34m"
)

func trabajador(parte string) {
	log.Printf(DebugColor, parte+" es la parte asignada")
	time.Sleep(time.Duration(rand.Int63n(1e6)))
	log.Printf(ErrorColor, "la parte "+parte+" Completada")

	wg.Done()
}

var (
	arreglo       string
	arregloPartes []string
	wg            sync.WaitGroup
)

func Run(nombre string) {
	//verificar que sea un entero, en ciclos de trabajo
	execute(nombre)
	for {
		fmt.Println(string(colorReset), "")
		fmt.Println("|----------------------------------------------------|")
		fmt.Println("|                Ejecutar de nuevo?                  |")
		fmt.Println("|  1)  Si                                            |")
		fmt.Println("|  2)  No, regresar                                  |")
		fmt.Println("|----------------------------------------------------|")
		fmt.Println("|  Seleccione una opcion:                            |")
		fmt.Print(string(colorYellow), ">> ", string(colorReset))
		com := bufio.NewScanner(os.Stdin)
		if com.Scan() {
			if com.Text() == "1" {
				execute(nombre)
			} else if com.Text() == "2" {
				break
			}
		}
	}
}


func execute(nombre string){
	fmt.Println(string(colorGreen), " Ingrese cantidad de ciclos de trabajo ", string(colorReset))
	fmt.Print(string(colorYellow), ">> ", string(colorReset))
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	ciclos := scanner.Text()
	if _, err := strconv.Atoi(ciclos); err == nil {
		fmt.Println(string(colorGreen), " Ingrese unidades de memoria ", string(colorReset))
		fmt.Print(string(colorYellow), ">> ", string(colorReset))
		scanner2 := bufio.NewScanner(os.Stdin)
		scanner2.Scan()
		arreglo := scanner2.Text()
		rep.AgregarSimulacion(ciclos, arreglo)
		arregloPartes = strings.Split(arreglo, ",")
		i, _ := strconv.Atoi(ciclos)
		fmt.Println("|----------------------------------------------------|")
		fmt.Print(string(colorBlue), ">> Nombre: ", nombre, "\n")
			//ejecucion
		rand.Seed(time.Now().UnixNano())
			for c := 1; c <= i; c++ {
				log.Printf(WarningColor, "----------  INICIA EL CICLO DE TRABAJO "+strconv.Itoa(c)+" ----------")
				wg.Add(len(arregloPartes))
				for _, parte := range arregloPartes {
					go trabajador(parte) // < ====  "FORK"
				}
				wg.Wait() // < ==== "JOIN"
				log.Printf(InfoColor, "---------- FINALIZA EL CICLO DE TRABAJO "+strconv.Itoa(c)+"----------")
			}

	} else {
		fmt.Printf(ErrorColor, "Debio ingresar un numero entero en 'ciclos de trabajo'")
	}
	
}

func colores() {
	fmt.Printf(InfoColor, "Info")
	fmt.Println("")
	fmt.Printf(NoticeColor, "Notice")
	fmt.Println("")
	fmt.Printf(WarningColor, "Warning")
	fmt.Println("")
	fmt.Printf(ErrorColor, "Error")
	fmt.Println("")
	fmt.Printf(DebugColor, "Debug")
	fmt.Println("")
}
