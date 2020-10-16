package model

type Collect struct {
	ID    int     `json:"id"`
	Title string  `json:"title"`
	Songs []*Song `json:"songs,omitempty"`
}
