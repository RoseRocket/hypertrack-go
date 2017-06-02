package hypertrack

import (
	"time"
)

/*
Sample:
{
  "id": "09ff4544-b947-4963-ae06-54f72558dd1d",
  "group_id": null,
  "lookup_id": null,
  "name": "Alex1",
  "phone": null,
  "photo": "https://hypertrack-api-v2-prod.s3.amazonaws.com/default_drivers/v2/a1.png",
  "availability_status": "offline",
  "vehicle_type": "unknown",
  "pending_actions": [],
  "last_location": {},
  "last_online_at": null,
  "last_heartbeat_at": null,
  "last_battery": null,
  "location_status": "location_available",
  "display": {
    "activity_text": "",
    "status_text": "Logged off",
    "sub_status_text": "",
    "has_new_sdk": true,
    "is_warning": false
  },
  "created_at": "2017-06-01T18:32:11.506337Z",
  "modified_at": "2017-06-01T18:32:11.506366Z",
  "is_connected": false
}
*/
type User struct {
	Id                 string             `json:"id"`        // Unique identifier for the object
	Name               string             `json:"name"`      // Name of the user
	Phone              string             `json:"phone"`     // E164 formatted phone number of user
	GroupId            *string            `json:"group_id"`  // The id of the fleet to which this user belongs
	LookupId           *string            `json:"lookup_id"` // A unique id that you can add to the user to identify
	Photo              string             `json:"photo"`
	AvailabilityStatus AvailabilityStatus `json:"availability_status"` // The availability status of the user, possible values are online, offline and active.
	VehicleType        string             `json:"vehicle_type"`
	PendingActions     []string           `json:"pending_actions"` // List of action ids that have been assigned to the user and are not complete
	// LastBattery ? `json:"last_battery"`
	LocationStatus  string     `json:"location_status"`
	Display         Display    `json:"display"`
	LastLocation    *Location  `json:"last_location"`     // The last known location of the user with the speed, bearing, location accuracy and recorded time fields
	LastHeartbeatAt *time.Time `json:"last_heartbeat_at"` // Date time for last communication with the user's device
	LastOnlineAt    *time.Time `json:"last_online_at"`    // Date time when the last tracking session was started for the user
	CreatedAt       time.Time  `json:"created_at"`        // Timestamp of when user was created
	ModifiedAt      time.Time  `json:"modified_at"`       // Timestamp of when user was last modified
	IsConnected     bool       `json:"is_connected"`
}

type AvailabilityStatus string

const (
	AvailabilityStatusOnline  = AvailabilityStatus("online")
	AvailabilityStatusOffline = AvailabilityStatus("offline")
	AvailabilityStatusActive  = AvailabilityStatus("active")
)

type Display struct {
	ActivityText  string `json:"activity_text"`
	StatusText    string `json:"status_text"`
	SubStatusText string `json:"sub_status_text"`
	HasNewSdk     bool   `json:"has_new_sdk"`
	IsWarning     bool   `json:"is_warning"`
}
