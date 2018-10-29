package structs

// Books - Model for the uses table
type Book struct {
	Title       string `json:"title" orm:"size(64)"`
	Author      string `json:"author" orm:"size(128)"`
	ISBN        string `json:"isbn" orm:"size(20)"`
	Description string `json:"description,omitempty" orm:"size(128)"`
	//define the book
}
