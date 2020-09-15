package pornhub

// Video structure
type Video struct {
	ID           string  `json:"video_id,omitempty"`
	Duration     string  `json:"duration,omitempty"`
	Views        float64 `json:"views,omitempty"`
	Rating       string  `json:"rating,omitempty"`
	Ratings      float64 `json:"ratings,omitempty"`
	Title        string  `json:"title,omitempty"`
	URL          string  `json:"url,omitempty"`
	DefaultThumb string  `json:"default_thumb,omitempty"`
	Thumb        string  `json:"thumb,omitempty"`
	PublishDate  string  `json:"publish_date,omitempty"`
	Thumbs       []Thumb
	Tags         []Tag
	Pornstars    []Pornstar
	Categories   []Category
	Segment      string `json:"segment,omitempty"`
}

// Thumb structure
type Thumb struct {
	Size   string `json:"size,omitempty"`
	Width  string `json:"width,omitempty"`
	Height string `json:"height,omitempty"`
	Src    string `json:"src,omitempty"`
}

// Tag structure
type Tag struct {
	Name string `json:"tag_name,omitempty"`
}

// Pornstar structure
type Pornstar struct {
	Name string `json:"pornstar_name,omitempty"`
}

// Category structure
type Category struct {
	Name string `json:"category,omitempty"`
}
