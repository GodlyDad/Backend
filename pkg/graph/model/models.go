package model

type Book struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Testament string    `json:"testament"`
	Chapters  []Chapter `gorm:"foreignKey:BookID"`
}

type Chapter struct {
	ID        string `json:"id"`
	BookID    string `json:"bookID"`
	ChapterID int    `json:"chapterID"`
	VerseID   int    `json:"verseID"`
	Verse     string `json:"verse"`
}
