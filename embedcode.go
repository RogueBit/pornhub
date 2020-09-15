package pornhub

// EmbedCode structure
type EmbedCode struct {
	Embed Code `json:"embed,omitempty"`
}

// Code structure
type Code struct {
	Code string `json:"code,omitempty"`
}
