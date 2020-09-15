package pornhub

// CategoriesList structure
type CategoriesList struct {
	Categories []CategoryList `json:"categories,omitempty"`
}

// CategoryList structure
type CategoryList struct {
	ID       string `json:"video_id,omitempty"`
	Category string `json:"category,omitempty"`
}
