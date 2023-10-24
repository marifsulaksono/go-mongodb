package config

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(conf DBConfig) (*mongo.Database, error) {
	db, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(conf.DatabaseURL))
	if err != nil {
		log.Fatalf("Error connection : %v" + err.Error())
	}

	return db.Database(conf.DatabaseName), err
}
