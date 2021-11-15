package model

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Animal struct {
	gorm.Model
	Name           string    `json:"name"`
	EatingBehavior string    `json:"eatingBehavior"`
	Extincted      bool      `json:"extincted"`
	ExtinctDate    time.Time `gorm:"default:null" json:"extinctDate"`
}

func (a *Animal) Extinct() {
	a.Extincted = true
	a.ExtinctDate = time.Now()
}

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Animal{})
	return db
}
