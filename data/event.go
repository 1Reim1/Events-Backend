package data

type Event struct {
	Id               int      `json:"id"`
	Title            string   `json:"title"`
	ShortDescription string   `json:"shortDescription"`
	Description      string   `json:"description"`
	EventDate        string   `json:"eventDate"`
	Latitude         float64  `json:"latitude"`
	Longitude        float64  `json:"longitude"`
	Images           []string `json:"images"`
	Preview          string   `json:"preview"`
}
