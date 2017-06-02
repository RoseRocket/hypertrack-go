package hypertrack

import "time"

type Location struct {
	ActivityConfidence int        `json:"activity_confidence"`
	Altitude           float64    `json:"altitude"`
	Activity           string     `json:"activity"`
	GeoJSON            GeoJSON    `json:"geojson"`
	RecordedAt         *time.Time `json:"recorded_at"`
	Speed              float64    `json:"speed"`
	Accuracy           float64    `json:"accuracy"`
}
