package hypertrack

type GeoJSON struct {
	Type        GeoJSONType `json:"type"`
	Coordinates Coordinate  `json:"coordinates"`
}

type GeoJSONType string

const (
	GeoJSONTypePoint = AvailabilityStatus("Point")
)

type Coordinate []float64
