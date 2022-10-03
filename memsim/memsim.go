package main

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
)

const (
	InfoColor    = "\033[1;32m%s\033[0m"
	NoticeColor  = "\033[1;36m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
	DebugColor   = "\033[0;36m%s\033[0m"
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

func main() {
	//limpiar pantalla

	fmt.Print("\033[H\033[2J")
	//memsim(ciclos, arregloPartes)

	fmt.Printf(InfoColor, "Ingrese cantidad de ciclos de trabajo>> ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	ciclos := scanner.Text()

	fmt.Printf("")
	fmt.Println("ingreso:", ciclos)

	//verificar que sea un entero, en ciclos de trabajo
	if _, err := strconv.Atoi(ciclos); err == nil {
		fmt.Printf(InfoColor, " Ingrese unidades de memoria>> ")
		scanner2 := bufio.NewScanner(os.Stdin)
		scanner2.Scan()
		arreglo := scanner2.Text()
		arregloPartes = strings.Split(arreglo, ",")

		i, _ := strconv.Atoi(ciclos)
		memsim(i, arregloPartes)

	} else {
		fmt.Printf(ErrorColor, "Debio ingresar un numero entero en 'ciclos de trabajo'")
	}

}

func memsim(ciclos int, arreglo []string) {

	rand.Seed(time.Now().UnixNano())
	for c := 1; c <= ciclos; c++ {
		log.Printf(WarningColor, "----------  INICIA EL CICLO DE TRABAJO "+strconv.Itoa(c)+" ----------")
		wg.Add(len(arreglo))
		for _, parte := range arreglo {
			go trabajador(parte) // < ====  "FORK"
		}
		wg.Wait() // < ==== "JOIN"
		log.Printf(InfoColor, "---------- FINALIZA EL CICLO DE TRABAJO "+strconv.Itoa(c)+"----------")

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
