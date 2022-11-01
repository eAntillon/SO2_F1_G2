package Server
import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	db "github.com/eAntillon/SO2_F1_G2/Db"
	ts "github.com/eAntillon/SO2_F1_G2/Types"
	mem "github.com/eAntillon/SO2_F1_G2/Memsim"
	
)
func PostLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var login ts.Login
	var getLogin ts.Getlogin
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		fmt.Fprintf(w, "error en kill")
	}
	resp := make(map[string]string)
	getLogin = db.Login(login)
	if login.Correo == getLogin.Correo && login.Password == getLogin.Password {
		resp["status"] = "authorized"
		json.NewEncoder(w).Encode(resp)
		return
	}else{
		resp["status"] = "unauthorized"
		json.NewEncoder(w).Encode(resp)
		return
	}
}


func PostRegistro(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var usuario ts.Login
	err := json.NewDecoder(r.Body).Decode(&usuario)
	if err != nil {
		fmt.Fprintf(w, "error en kill")
	}
	db.CrearRegistro(usuario)
	json.NewEncoder(w).Encode("It has been created successfully")

}

func inicio(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	resp := make(map[string]string)
	resp["data"] = "Buenas"
	json.NewEncoder(w).Encode(resp)
}

func memsim(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var mems ts.MemStruct
	err := json.NewDecoder(r.Body).Decode(&mems)
	if err != nil {
		fmt.Fprintf(w, "error en kill")
	}
	resp := mem.Run(mems.Ciclos, mems.Unidades)
	
	json.NewEncoder(w).Encode(resp)

}

func Backend(){
	router := mux.NewRouter()
	router.HandleFunc("/api", inicio)
	router.HandleFunc("/api/login", PostLogin) .Methods("POST", "OPTIONS")
	router.HandleFunc("/api/registro", PostRegistro).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/memsim", memsim).Methods("POST", "OPTIONS")
	fmt.Println("Server running on port 5000")
	// iniciar servidor
	http.ListenAndServe(":5000", router)   
}

