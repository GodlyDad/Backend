package helpers

import (
	"github.com/GodlyDad/Backend/pkg/graph/model"
	"gorm.io/gorm"
)

func PopulateTranslations(db *gorm.DB) error {

	translations := []*model.Translation{
		{
			ID:           "1",
			Abbreviation: "KJV",
			Version:      "King James Version",
		},
		{
			ID:           "2",
			Abbreviation: "NLT",
			Version:      "New Living Translation",
		},
	}

	return db.FirstOrCreate(&translations, translations).Error
}
