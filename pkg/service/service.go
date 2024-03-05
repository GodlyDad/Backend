package service

import "gorm.io/gorm"

type IBibleService interface {
}

type BibleService struct {
	GormDB *gorm.DB
}

func NewBibleService(db *gorm.DB) *BibleService {
	return &BibleService{
		GormDB: db,
	}
}
