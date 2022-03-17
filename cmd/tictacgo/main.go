package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	index "github.com/tictacgo/config/environment"

	//pgDatabase "github.com/tictacgo/config/pgDatabase"
	//redisConnection "github.com/tictacgo/config/redis"
	//gameStateStorage "github.com/tictacgo/internal/dao/playgroundCacheStorage"
	//gameInfoStorage "github.com/tictacgo/internal/dao/playgroundDatabaseStorage"
	playgroundManager "github.com/tictacgo/internal/manager/playgroundManager"
)

func main() {
	router := mux.NewRouter()
	switch mode := index.EnvironmentConfig.ENV; mode {
	case "Development":
		//TODO
		//db, err := pgDatabase.SetDBConnection()
		//redis := redisConnection.NewRedisConfig(index.EnvironmentConfig.REDIS_CLIENT)
		// if err != nil {
		// 	errors.Wrap(err, "Database Connection is Invalid")
		// 	panic(err)
		// }
		//go gameInfoStorage.NewdbStorage(db)
		//go gameStateStorage.NewcacheStorage(redis)
		playgroundManager.InitilizeBoard()
		//nicely close the database connection
		//defer db.Close()
	default:
		log.Printf("Invalid Input")
	}
	http.ListenAndServe(index.EnvironmentConfig.ROUTER_PORT, router)
}
