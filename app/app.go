package app

import (
	"fmt"
	"log"

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

	// db, err
	_, err := gorm.Open(config.DB.Dialect, dbURI)
	if err != nil {
		log.Fatal("Não foi possível se conectar ao banco de dados!")
	}
}
