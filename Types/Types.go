package types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	
)
/*
type Tuit struct {

    ID primitive.ObjectID `json:"_id,omitempty"`
	Nombre string `json:"nombre"`
	Comentario string `json:"comentario"`
	Fecha string `json:"fecha"`
	Hashtags []string `json:"hashtags"`
	Upvotes int `json:"upvotes"`
	Downvotes int `json:"downvotes"`

}

type Log struct {
   // StatusNumber int `json:"statusNumber"`
    Message string `json:"message"`
	Time time.Duration `json:"time"`
}
*/

type RegistroUsuario struct {
	Correo string `json:"correo"`
    Password string `json:"password"`
}


type Login struct {
	Correo string `json:"correo"`
    Password string `json:"password"`
}

type Getlogin struct {

    ID primitive.ObjectID `json:"_id,omitempty"`
	Correo string `json:"correo"`
    Password string `json:"password"`

}


