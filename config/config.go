package config

import (
	"fmt"
	"log"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/ftsog/ecom/handlers"
	"github.com/ftsog/ecom/models"
	"github.com/ftsog/ecom/routers"
	"github.com/ftsog/ecom/utils"
)

type dbConfig struct {
	host      string
	port      string
	user      string
	dbName    string
	password  string
	secretKey string
	sslmode   string
}

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error Loading Environment file", err)
		return
	}
}

func initDb() *dbConfig {
	dbConfig := &dbConfig{
		host:     os.Getenv("HOST"),
		port:     os.Getenv("PORT"),
		user:     os.Getenv("DBUSER"),
		dbName:   os.Getenv("DBNAME"),
		password: os.Getenv("PASSWORD"),
		sslmode:  os.Getenv("SSLMODE"),
	}

	return dbConfig

}

func NewModel() *models.Model {
	LoadEnv()
	dbConfig := initDb()
	dbConn, err := utils.DatabaseConnection(dbConfig.host, dbConfig.port, dbConfig.user, dbConfig.dbName, dbConfig.password, dbConfig.sslmode)
	if err != nil {
		log.Fatal(err)
	}

	rdStore, err := utils.RediStore(10, "tcp", dbConfig.host, os.Getenv("REDIS_PORT"), dbConfig.password, []byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		log.Fatal(err)
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
