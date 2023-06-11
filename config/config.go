package config

import (
	"log"
	"os"

	"github.com/ftsog/ecom/handlers"
	"github.com/ftsog/ecom/models"
	"github.com/ftsog/ecom/routers"
	"github.com/ftsog/ecom/utils"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

type dbConfig struct {
	databaseHost     string
	databasePort     string
	databaseUser     string
	databaseName     string
	databasePassword string
	sslmode          string
	redisHost        string
	redisPort        string
	redisPassword    string
	secretKey        []byte
}

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Panic(err)
	}
}

func initDb() *dbConfig {
	dbConfig := &dbConfig{
		databaseHost:     os.Getenv("DATABASE_HOST"),
		databasePort:     os.Getenv("DATABASE_PORT"),
		databaseUser:     os.Getenv("DATABASE_USER"),
		databaseName:     os.Getenv("DATABASE_NAME"),
		databasePassword: os.Getenv("DATABASE_PASSWORD"),
		//sslmode:          os.Getenv("SSLMODE"),
		redisHost:     os.Getenv("REDIS_HOST"),
		redisPort:     os.Getenv("REDIS_PORT"),
		redisPassword: os.Getenv("REDIS_PASSWORD"),
		secretKey:     []byte(os.Getenv("SECRET_KEY")),
	}

	return dbConfig

}

func NewModel() *models.Model {
	LoadEnv()
	dbConfig := initDb()
	dbConn, err := utils.DatabaseConnection(dbConfig.databaseHost, dbConfig.databasePort, dbConfig.databaseUser, dbConfig.databaseName, dbConfig.databasePassword, dbConfig.sslmode)
	if err != nil {
		log.Panic(err)
	}

	rdStore, err := utils.RediStore(10, "tcp", dbConfig.redisHost, dbConfig.redisPort, dbConfig.redisPassword, dbConfig.secretKey)
	if err != nil {
		log.Panic(err)
	}

	m := &models.Model{
		DB:        dbConn,
		RediStore: rdStore,
	}

	return m
}

func NewHandler() *handlers.Handler {
	m := NewModel()
	l := utils.NewLogger()
	h := &handlers.Handler{
		Db:     m,
		Logger: l,
	}

	return h
}

func NewRouter(r *chi.Mux) *routers.Router {
	handler := NewHandler()
	router := &routers.Router{
		Route:   r,
		Handler: handler,
	}

	return router
}
