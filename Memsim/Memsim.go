package Memsim

import (
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"
)

func trabajador(parte string, c int ) (string){
	var procesos = ""
	procesos += "No." +strconv.Itoa(c)+": "+ parte+" es la parte asignada,"
	time.Sleep(time.Duration(rand.Int63n(1e6)))
	procesos += "No." +strconv.Itoa(c)+": la parte "+parte+" Completada,"
	wg.Done()
	return procesos
}

var (
	arreglo       string
	arregloPartes []string
	wg            sync.WaitGroup
)

func Run(ciclos int, arreglo string ) (map[string]string){
	arregloPartes = strings.Split(arreglo, ",")
	rand.Seed(time.Now().UnixNano())
	resp := make(map[string]string)
	var i = 1;
	for c := 1; c <= ciclos; c++ {
		var data = "----------  INICIA EL CICLO DE TRABAJO "+strconv.Itoa(c)+" ----------,"
		wg.Add(len(arregloPartes))
		for _, parte := range arregloPartes {
			data += trabajador(parte, i) // < ====  "FORK"
			i++
		}
		wg.Wait() // < ==== "JOIN"
		data += "---------- FINALIZA EL CICLO DE TRABAJO "+strconv.Itoa(c)+"----------"	
		var name = "data" +strconv.Itoa(c)
		resp[name] = data
	}
	return resp
}