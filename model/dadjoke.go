package model

type DadJokes struct {
	ID        uint    `gorm: "primary key;autoIncrement" json:"id"`
	Author    *string `json:"author"`
	Title     *string `json:"title"`
	Publisher *string `json:"publisher"`
}
type DadJoke struct {
	Author    string `json:"author"`
	Title     string `json:"title"`
	Publisher string `json:"publisher"`
}
