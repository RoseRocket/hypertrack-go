package hypertrack

import "time"

type Place struct {
	Id         string    `json:"id"`          // Unique identifier for the object
	Name       string    `json:"name"`        // The name of the entity, eg, hub, customer location
	Location   GeoJSON   `json:"location"`    // Location in GeoJSON format
	Address    string    `json:"address"`     // Street address of the object
	Landmark   string    `json:"landmark"`    // Landmark near the object
	ZipCode    string    `json:"zip_code"`    // Zip or postal code of the object
	City       string    `json:"city"`        // City of the object
	State      string    `json:"state"`       // State of the object
	Country    string    `json:"country"`     // Country of the object
	CreatedAt  time.Time `json:"created_at"`  // Timestamp of when customer was created
	ModifiedAt time.Time `json:"modified_at"` // Timestamp of when customer was modified
}
