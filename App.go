// Package spark provides the main application structure and setup for the Spark framework.
package spark

import (
	"net/http"

	DIPackage "github.com/AlladinDev/Spark/internals/Di"
	RouterPkg "github.com/AlladinDev/Spark/internals/Router"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	DI          *DIPackage.Package
	Router      *http.ServeMux
	MongoClient *mongo.Client
	Redis       *redis.Client
}

func NewApp() *App {
	return &App{
		DI:     DIPackage.NewDI(),
		Router: RouterPkg.NewRouter(),
	}
}

func (a *App) WithMongoDB(mongoClient *mongo.Client) *App {
	a.MongoClient = mongoClient
	return a
}

func (a *App) WithRedis(redisClient *redis.Client) *App {
	a.Redis = redisClient
	return a
}

func (a *App) Run(addr string) error {
	return http.ListenAndServe(addr, a.Router)

}
