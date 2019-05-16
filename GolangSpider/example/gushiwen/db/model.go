package db

import "github.com/jinzhu/gorm"

//诗文对象
type Poem struct {
	ID	string	`gorm:"primary_key"`
	Title string	`gorm:"type:varchar(64);"`
	Author *Author	`gorm:"foreignkey:AuthorRefer"`
	Dynasty	string
	Content  string `gorm:"type:varchar(2000);"`
	Translate string `gorm:"type:varchar(2000);"`
	Like	string `gorm:"type:varchar(2000);"`
	Background	string `gorm:"type:varchar(2000);"`
	Url		string
}

func NewPoem(title string, author *Author, url string) *Poem {
	return &Poem{Title: title, Author: author, Url: url}
}


//作者
type Author struct {
	gorm.Model
	Name string
	Profile string	`gorm:"size(2000)"`
	//Poems []*Poem
	Picture	string
}

func NewAuthor2(name string, profile string, picture string) *Author {
	return &Author{Name: name, Profile: profile, Picture: picture}
}

func NewAuthor(name string) *Author {
	return &Author{Name: name}
}
