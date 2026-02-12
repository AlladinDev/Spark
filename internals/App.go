package internals

// import (
// 	"net/http"

// 	"github.com/redis/go-redis"
// 	"go.mongodb.org/mongo-driver/mongo"
// )

// func NewApp() *App {
// 	return &App{
// 		DI:     NewDI(),
// 		Router: NewRouter(),
// 	}
// }

// func (a *App) WithMongoDB(mongoClient *mongo.Client) *App {
// 	a.MongoClient = mongoClient
// 	return a
// }

// func (a *App) WithRedis(redisClient *redis.Client) *App {
// 	a.Redis = redisClient
// 	return a
// }

// func (a *App) Run(addr string) error {
// 	return http.ListenAndServe(addr, a.Router.Mux)
// }
