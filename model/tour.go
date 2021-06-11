package model

type Tour struct {
	ID       int        `json:"id"`
	Title    string     `json:"title"`
	Collects []*Collect `json:"collects"`
}
