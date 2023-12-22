package models

type Joke struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}