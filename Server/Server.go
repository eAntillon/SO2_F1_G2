package Server
import (
	"fmt"
	//p "github.com/eAntillon/SO2_F1_G2/Menu"
	"net/http"
	"github.com/gin-gonic/gin"
	//"strings"
	//"github.com/eAntillon/SO2_F1_G2"
	ts "github.com/eAntillon/SO2_F1_G2/Types"
	db "github.com/eAntillon/SO2_F1_G2/Db"
	
)
func getTest(c *gin.Context) {
    c.JSON(http.StatusOK, "conexion exitosa")
}


//-- Peticion de Usuarios   
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

func PostRegistro(c *gin.Context) {
	var registro ts.RegistroUsuario
	
	if err := c.ShouldBind(&registro); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	db.CrearRegistro(registro)
	c.JSON(http.StatusCreated, "it has been created successfully")

}












func Backend(){
	fmt.Println("")
    fmt.Println(" ==========================  SERVIDOR  ========================== ")
    fmt.Println("")
	router := gin.Default()
    router.Use(gin.Recovery()) // Para recuperarse de Errores y enviar un 500
	router.GET("/test", getTest) 
	router.POST("/login", PostLogin) 
	router.POST("/registro", PostRegistro) 
	                              
	router.Run(":4000")   
}

