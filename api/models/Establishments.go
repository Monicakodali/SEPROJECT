package models

//establishment structure
type Establishment struct {
	ID           string  `json:"id"`
	TYPE         string  `json:"type"`
	NAME         string  `json:"name"`
	BUILDING     string  `json:"building"`
	ROOM         string  `json:"room"`
	URL          string  `json:"url"`
	X_COORDINATE float64 `json:"x"`
	Y_COORDINATE float64 `json:"y"`
}
