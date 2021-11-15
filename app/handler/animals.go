package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ArtuoS/animals-api/app/model"
	"github.com/jinzhu/gorm"
)

func GetAllAnimals(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	animals := []model.Animal{}
	db.Find(&animals)
	respondJSON(w, http.StatusOK, animals)
}

func CreateAnimal(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	animal := []model.Animal{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&animal); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&animal).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, animal)
}
