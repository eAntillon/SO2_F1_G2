package main

/*  Paquetes que se instalaron
    go get -u github.com/gin-gonic/gin
    npm install -g nodemon
    go get go.mongodb.org/mongo-driver/mongo
    go get github.com/joho/godotenv
*/

import (
	"fmt"
	"os"
	 //p "github.com/eAntillon/SO2_F1_G2/Menu"
	 s "github.com/eAntillon/SO2_F1_G2/Server"
	 "github.com/joho/godotenv"
	
)



func main() {
	//p.Menu_principal()
   if os.Getenv("DB")==""{
        err := godotenv.Load(".env")
        if err!=nil{
            fmt.Println("Error loading enviroment variables")
        }
   }
    s.back()                           
	//nodemon --exec go run main.go --signal SIGTERM
	
}

