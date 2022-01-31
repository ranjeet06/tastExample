package models

type Book struct {
	ID     string `xml:"id" json:"id"`
	Title  string `xml:"title" json:"title"`
	Author string `xml:"author" json:"author"`
}
