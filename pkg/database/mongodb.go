package database

import (
	"context"
	"fmt"
	"github.com/hiiamtrong/golang-gin-mongodb-rest-api-starter/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Mongodb struct {
	Database *mongo.Database
	*mongo.Client
}

func NewMongodb(config *config.Config) *Mongodb {
	mongoCfg := config.Mongodb
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/admin?authSource=admin", mongoCfg.Username, mongoCfg.Password, mongoCfg.Host, mongoCfg.Port)
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Panicln("database.NewMongodb:Connect error: ", err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Panicln("database.NewMongodb:Ping error: ", err)
	}

	database := client.Database(mongoCfg.Database)

	return &Mongodb{
		Database: database,
		Client:   client,
	}
}

func (m *Mongodb) Shutdown(ctx context.Context) error {
	return m.Disconnect(ctx)
}
