package repositories

import (
	"core"
	"gorm.io/gorm"
	"models"
)

type Books struct {
	core.Repository
}

func (r *Books) FindAll() (*gorm.DB, []models.Book) {
	var list []models.Book
	result := r.Db().Find(&list)
	return result, list
}