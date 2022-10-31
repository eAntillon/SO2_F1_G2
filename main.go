package main
// go get -u github.com/gin-gonic/gin
//npm install -g nodemon
//go get go.mongodb.org/mongo-driver/mongo
//go get github.com/joho/godotenv
import (
	"fmt"
	"os"
	//"bufio"
	 //p "github.com/eAntillon/SO2_F1_G2/Menu"
	 s "github.com/eAntillon/SO2_F1_G2/Server"
	 "github.com/joho/godotenv"
	//"net/http"
	//"github.com/gin-gonic/gin"
  	//"strings"
	//"github.com/eAntillon/SO2_F1_G2"
	
)



func main() {
	//p.Menu_principal()

	if os.Getenv("DB")==""{
        err := godotenv.Load(".env")
        if err!=nil{
            fmt.Println("Error loading enviroment variables")
        }
   }

	 s.Backend()                           
	//nodemon --exec go run main.go --signal SIGTERM
	//router.Run(":3000")
}

