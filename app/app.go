package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ArtuoS/animals-api/app/handler"
	"github.com/ArtuoS/animals-api/app/model"
	"github.com/ArtuoS/animals-api/config"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (a *App) Initialize(config *config.Config) {
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True",
		config.DB.Username,
		config.DB.Password,
		config.DB.Host,
		config.DB.Port,
		config.DB.Name,
		config.DB.Charset)

	db, err := gorm.Open(config.DB.Dialect, dbURI)
	if err != nil {
		log.Fatal("Não foi possível se conectar ao banco de dados!")
	}

	a.DB = model.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}

func (a *App) setRouters() {
	a.Get("/animals", a.handleRequest(handler.GetAllAnimals))
}

// Get wraps the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

type RequestHandlerFunction func(db *gorm.DB, w http.ResponseWriter, r *http.Request)

func (a *App) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(a.DB, w, r)
	}
}
