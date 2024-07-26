package models

type Post struct {
	ID      int    `json:"id" gorm:"primary_key;auto_increment"`
	Title   string `json:"title" gorm:"type:varchar(255);not null"`
	Content string `json:"content" gorm:"type:text;not null"`
	Tags    []Tag  `gorm:"many2many:post_tags;"`
}

type Tag struct {
	ID    int     `json:"id" gorm:"primary_key;auto_increment"`
	Label string  `json:"label" gorm:"type:varchar(255);not null"`
	Posts []*Post `gorm:"many2many:post_tags;"`
}

type PostTag struct {
	PostID int `gorm:"type:integer;not null"`
	TagID  int `gorm:"type:integer;not null"`
}
