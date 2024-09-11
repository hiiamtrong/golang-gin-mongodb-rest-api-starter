package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hiiamtrong/golang-gin-mongodb-rest-api-starter/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongodb struct {
	Database *mongo.Database
	*mongo.Client
}

func NewMongodb(config *config.Config) *Mongodb {
	mongoCfg := config.Mongodb
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/admin?authSource=admin", mongoCfg.Username, mongoCfg.Password, mongoCfg.Host, mongoCfg.Port)
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Panicln("database.NewMongodb:Connect error: ", err)
	}

	err = client.Ping(ctx, nil)
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
