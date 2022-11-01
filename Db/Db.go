package Db
import (
	"fmt"
	//"errors"
	"os"
	//"time"
	"context"
	"log"
	"encoding/json"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	ts "github.com/eAntillon/SO2_F1_G2/Types"
)

var collection *mongo.Collection
var ctx = context.TODO()

func connect()(*mongo.Collection,error){
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_CS"))
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	return client.Database(os.Getenv("DB")).Collection(os.Getenv("COLLECTION")),nil
}
func CrearRegistro(registro ts.Login)(string,error){
	
	collection,errr :=connect()

	if errr != nil {
        return "Esta vacio",nil 
    }
	r,err := collection.InsertOne(ctx, registro)
    if err != nil {
        return "Esta vacio",nil 
    }
    return " loaded successfully with id: "+fmt.Sprint(r.InsertedID), nil
}

func Login(login ts.Login)(ts.Getlogin ){
	collection,err :=connect()
	cur, currErr := collection.Find(ctx, login)

    if currErr != nil { panic(currErr) }
    defer cur.Close(ctx)

    var results[] ts.Getlogin 
    var resultado ts.Getlogin 

    if err = cur.All(ctx, &results); err != nil {
          panic(err)
    }

	for _, result := range results {
		cur.Decode(&result)
		output, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			panic(err)
		}
	   
	   // fmt.Printf("%s\n", output)
	   if err := json.Unmarshal(output, &resultado); err != nil {
		    panic(err)
       }
	}
   return resultado

}














