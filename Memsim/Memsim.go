package Memsim

import (
	"fmt"
	"log"
	"math/rand"
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
	colorReset = "\033[0m"
	colorYellow = "\033[33m"
	colorGreen = "\033[32m"
	colorBlue = "\033[34m"
)

func trabajador(parte string, c int ) {
	log.Printf(DebugColor,"No." +strconv.Itoa(c)+":"+ parte+" es la parte asignada")
	time.Sleep(time.Duration(rand.Int63n(1e6)))
	log.Printf(ErrorColor,"No." +strconv.Itoa(c)+": la parte "+parte+" Completada")

	wg.Done()
}

var (
	arreglo       string
	arregloPartes []string
	wg            sync.WaitGroup
)

func execute(ciclos int, arreglo string ){
	arregloPartes = strings.Split(arreglo, ",")
	rand.Seed(time.Now().UnixNano())
		for c := 1; c <= ciclos; c++ {
			log.Printf(WarningColor, "----------  INICIA EL CICLO DE TRABAJO "+strconv.Itoa(c)+" ----------")
			wg.Add(len(arregloPartes))
			for _, parte := range arregloPartes {
				go trabajador(parte, c) // < ====  "FORK"
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
