package Server
import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	db "github.com/eAntillon/SO2_F1_G2/Db"
	ts "github.com/eAntillon/SO2_F1_G2/Types"
)


func PostLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var login ts.Login
	var getLogin ts.Getlogin
	err := json.NewDecoder(r.Body).Decode(&login)
	fmt.Println("usuario: ", login)
	if err != nil {
		fmt.Fprintf(w, "error en kill")
	}
	resp := make(map[string]string)
	getLogin = db.Login(login)
	if login.Correo == getLogin.Correo && login.Password == getLogin.Password {
		resp["status"] = "authorized"
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			fmt.Fprintf(w, "error en kill")
		}
		json.NewEncoder(w).Encode(jsonResp)
		return
	}else{
		resp["status"] = "unauthorized"
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			fmt.Fprintf(w, "error en kill")
		}
		json.NewEncoder(w).Encode(jsonResp)
		return
	}
	
}


func PostRegistro(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var usuario ts.Login
	err := json.NewDecoder(r.Body).Decode(&usuario)
	fmt.Println("usuario: ", usuario)
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
	json.NewEncoder(w).Encode("Buenas")
}



func main() {
	//router
	router := mux.NewRouter()

	router.HandleFunc("/", inicio)
	router.HandleFunc("/login", PostLogin) .Methods("POST", "OPTIONS")
	router.HandleFunc("/registro", PostRegistro).Methods("POST", "OPTIONS")
	fmt.Println("Server running on port 5000")
	// iniciar servidor
	http.ListenAndServe(":5000", router)
}
