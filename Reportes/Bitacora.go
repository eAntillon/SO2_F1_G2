package reportes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Usuario struct {
	Usuario   string      `json:"usuario"`
	Actividad []Actividad `json:"actividad"`
}

type Actividad struct {
	Funcion     string `json:"funcion"`
	Ejecuciones int    `json:"ejecuciones"`
}

var UsuarioActivo *Usuario
var UsuariosSesion []Usuario

func LeerBitacora() []Usuario {

	var bitacora []Usuario

	f, err := os.OpenFile("Bitacora.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	_ = f
	file, err := ioutil.ReadFile("Bitacora.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(file, &bitacora)
	// OCULTAR
	if err != nil {
		log.Fatal(err)
		// return  empty array
		return []Usuario{}
	}
	return bitacora
}

func CrearUsuario(nombre string) Usuario {
	actividades := []Actividad{
		{"IOTOP", 0},
		{"TOP", 0},
		{"SYSCALL", 0},
	}
	usuario := Usuario{nombre, actividades}
	UsuarioActivo = &usuario
	return usuario
}

func AgregarActividad(actividad string) {
	usuario := UsuarioActivo
	for i := 0; i < len(usuario.Actividad); i++ {
		if usuario.Actividad[i].Funcion == actividad {
			usuario.Actividad[i].Ejecuciones++
		}
	}
}

func GuardarUsuario() {
	if UsuarioActivo != nil {
		if UsuariosSesion == nil {
			UsuariosSesion = []Usuario{*UsuarioActivo}
		} else {

			UsuariosSesion = append(UsuariosSesion, *UsuarioActivo)
		}
	}
}

func EscribirBitacora() {
	bitacora := LeerBitacora()
	bitacora = append(bitacora, UsuariosSesion...)
	json, _ := json.Marshal(bitacora)
	ioutil.WriteFile("Bitacora.json", json, 0644)
	UsuariosSesion = nil
	UsuarioActivo = nil
	fmt.Println("Bitácora generada con éxito.")
}
