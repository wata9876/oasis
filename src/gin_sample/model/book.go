package model

//Book 書籍情報
type Book struct {
	ID      int    `xorm:"pk autoincr int" form:"id" json:"id"`
	Title   string `json:"title" 　　　　　　　xorm:"'title'"`
	Author  string `json:"author" 　　　　　　　xorm:"'author'"`
	Content string `json:"content" 　　　　　　　xorm:"'content'"`
}
