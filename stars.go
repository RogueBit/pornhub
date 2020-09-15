package pornhub

// StarsList structure
type StarsList struct {
	Stars []struct {
		Star struct {
			StarName string `json:"star_name"`
		} `json:"star"`
	} `json:"stars"`
}
