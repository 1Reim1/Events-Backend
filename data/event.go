package data

type Event struct {
	Id               int
	Title            string
	ShortDescription string
	Description      string
	EventDate        string
	Latitude         float64
	Longitude        float64
	Images           []string
	Preview          string
}
