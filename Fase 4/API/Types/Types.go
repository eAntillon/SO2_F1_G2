package types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	
)


type Login struct {
	Correo 	 string `json:"correo"`
    Password string `json:"password"`
}

type Getlogin struct {
    ID primitive.ObjectID `bson:"_id"`
	Correo 	 string 	  `bson:"correo"`
    Password string 	  `bson:"password"`
}

