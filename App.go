package spark

import (
	DIPackage "github.com/AlladinDev/Spark/Internals/Di"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	DI *DIPackage.DI

	MongoClient *mongo.Client
	Redis       *redis.Client
}

func NewApp() *App {
	return &App{
		DI: NewApp().DI,
		// Router: NewApp().Router,
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
	//return http.ListenAndServe(addr, a.Router.Mux)
	return nil
}
