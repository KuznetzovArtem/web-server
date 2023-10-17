package repository

import (
	"gorm.io/gorm"
	"web-server/internal/dto"
)

type RequestWriterRepo struct {
	db *gorm.DB
}

func NewRequestWriterRepo(db *gorm.DB) *RequestWriterRepo {
	return &RequestWriterRepo{
		db: db,
	}
}

func (r *RequestWriterRepo) Add(data dto.Request) {
	r.db.Model(&dto.Request{}).Create(data)
}

func (r *RequestWriterRepo) GetAll() []dto.Request {
	var results []dto.Request
	r.db.Model(&dto.Request{}).Find(&results)

	return results
}

func (r *RequestWriterRepo) DeleteAll() {
	r.db.Where("1 = 1").Delete(&dto.Request{})
}
