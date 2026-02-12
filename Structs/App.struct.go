package structs

import (
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	DI          *DI
	Router      *Router
	MongoClient *mongo.Client
	Redis       *redis.Client
}
