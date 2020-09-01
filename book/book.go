package book

// Book a book (id, title, author, isbn)
type Book struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"ISBN"`
}
