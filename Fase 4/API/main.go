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
)

type Login struct {
	Correo 	 string `json:"correo"`
    Password string `json:"password"`
}

type Getlogin struct {

    ID primitive.ObjectID `bson:"_id"`
	Correo string `bson:"correo"`
    Password string `bson:"password"`

}

func PostLogin(c *gin.Context) {
	var login ts.Login
	var getLogin ts.Getlogin
	if err := c.ShouldBind(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	getLogin =db.Login(login)
	if login.Correo == getLogin.Correo && login.Password == getLogin.Password {
		c.JSON(http.StatusOK, gin.H{"status": "authorized"})
		return
	}else{
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}
	
}

func PostRegistro(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	ingreso := Login{}
	err := json.NewDecoder(r.Body).Decode(&ingreso)
	fmt.Println("INGRESO: ", ingreso)
	if err != nil {
		fmt.Fprintf(w, "error en kill")
	}

	
	if err := c.ShouldBind(&registro); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	db.CrearRegistro(registro)
	c.JSON(http.StatusCreated, "it has been created successfully")

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
