package model

type Tour struct {
	ID       int        `json:"id"`
	Collects []*Collect `json:"collects"`
}
