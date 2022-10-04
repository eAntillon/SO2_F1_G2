package reportes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)



func LeerSimulaciones() []EntradaSimulacion {

	var simulacion []EntradaSimulacion

	f, err := os.OpenFile("Simulacion.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	_ = f
	file, err := ioutil.ReadFile("Simulacion.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(file, &simulacion)
	if err != nil {
		return []EntradaSimulacion{}
	}
	return simulacion
}

func AgregarSimulacion(procesos string, unidades string) {
	usuario := UsuarioSim
	procesosInt, _ := strconv.Atoi(procesos)
	usuario.Simulaciones = append(usuario.Simulaciones, Simulacion{procesosInt, unidades})
}

func EscribirSimulacion() {
	simulaciones := LeerSimulaciones()
	simulaciones = append(simulaciones, UsuariosSim...)
	json, _ := json.Marshal(simulaciones)
	ioutil.WriteFile("Simulacion.json", json, 0644)
	UsuariosSim = nil
	UsuarioSim = nil
	fmt.Println("Estado de simulación generado con éxito.")
}
