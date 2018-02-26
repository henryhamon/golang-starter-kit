package helper

import (
	"math"

	"github.com/jinzhu/gorm"
)

func (parameter *Parameter) Paginate(db *gorm.DB, dataSource interface{}) (*gorm.DB, error) {
	var count int
	done := make(chan bool, 1)

	go CountRecords(db, dataSource, done, &count)
	db = db.Limit(parameter.Limit).Offset(parameter.Limit * (parameter.Page - 1))
	<-done

	parameter.TotalRecords = count
	parameter.TotalPages = GetTotalPages(parameter.Limit, count)

	return db, nil
}

func CountRecords(db *gorm.DB, countDataSource interface{}, done chan bool, count *int) {
	db.Model(countDataSource).Count(count)
	done <- true
}

func GetTotalPages(perPage int, totalRecords int) int64 {
	totalPages := float64(totalRecords) / float64(perPage)
	return int64(math.Ceil(totalPages))
}
