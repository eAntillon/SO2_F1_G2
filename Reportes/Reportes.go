package reportes

type Usuario struct {
	Usuario   string      `json:"usuario"`
	Actividad []Actividad `json:"actividad"`
}

type EntradaSimulacion struct {
	Usuario      string       `json:"usuario"`
	Simulaciones []Simulacion `json:"actividad"`
}

type Actividad struct {
	Funcion     string `json:"funcion"`
	Ejecuciones int    `json:"ejecuciones"`
}

type Simulacion struct {
	Procesos int    `json:"procesos"`
	Unidades string `json:"unidades"`
}

var UsuarioActivo *Usuario
var UsuariosSesion []Usuario

var UsuarioSim *EntradaSimulacion
var UsuariosSim []EntradaSimulacion

func CrearUsuario(nombre string) {
	actividad := []Actividad{
		{"IOTOP", 0},
		{"TOP", 0},
		{"SYSCALL", 0},
	}
	usuario := Usuario{nombre, actividad}
	UsuarioActivo = &usuario

	simulaciones := []Simulacion{}
	usuarioSimulacion := EntradaSimulacion{nombre, simulaciones}
	UsuarioSim = &usuarioSimulacion

}

func GuardarUsuario() {
	if UsuarioActivo != nil {
		if UsuariosSesion == nil {
			UsuariosSesion = []Usuario{*UsuarioActivo}
		} else {

			UsuariosSesion = append(UsuariosSesion, *UsuarioActivo)
		}
	}

	if UsuarioSim != nil {
		if UsuariosSim == nil {
			UsuariosSim = []EntradaSimulacion{*UsuarioSim}
		} else {

			UsuariosSim = append(UsuariosSim, *UsuarioSim)
		}
	}

}
