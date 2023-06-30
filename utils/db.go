package utils

import (
    "context"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type Collections struct {
    Stories *mongo.Collection
	Dictionary *mongo.Collection
}

func (c *Collections) ConnectDb()  {
	client, err := mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI("mongodb+srv://doadmin:Z3d87ni4E91g05aX@logiciel-applab-dab57134.mongo.ondigitalocean.com/admin?tls=true&authSource=admin&replicaSet=logiciel-applab&tlsCAFile=ca-certificate.crt"))
	if err != nil {
		panic(err)
	}
	db :=  client.Database("LearnReadApp")
	c.Stories = db.Collection("Stories")
	c.Dictionary =  db.Collection("Dictionary")
}

var DbCollections = Collections{}
