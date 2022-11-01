package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	db "github.com/eAntillon/SO2_F1_G2/Fase 4/API/Db"
	ts "github.com/eAntillon/SO2_F1_G2/Fase 4/API/Types"
)


func PostLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	login ts.Login
	getLogin ts.Getlogin
	err := json.NewDecoder(r.Body).Decode(&login)
	fmt.Println("usuario: ", login)
	if err != nil {
		fmt.Fprintf(w, "error en kill")
	}
	getLogin = db.Login(login)

	if login.Correo == getLogin.Correo && login.Password == getLogin.Password {
		json.NewEncoder(w).Encode("status": "authorized")
		return
	}else{
		json.NewEncoder(w).Encode("status": "unauthorized")
		return
	}
	
}


func PostRegistro(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	usuario ts.Login
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
