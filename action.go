package hypertrack

import "time"

type Action struct {
	Id             string          `json:"id"`              // Unique identifier for the object
	User           *User           `json:"user"`            // The user object to which the action was assigned
	Type           ActionType      `json:"type"`            // The type of action, can be pickup, delivery, dropoff, visit, stopover or task
	LookupId       *string         `json:"lookup_id"`       // An identifier for the action that can based on your internal ids
	AssignedAt     *time.Time      `json:"assigned_at"`     // The time at which the action was assigned to the user
	ScheduledAt    *time.Time      `json:"scheduled_at"`    // The scheduled time by when the action should be completed
	ExpectedPlace  Place           `json:"expected_place"`  // The place object where the action is to be completed
	Status         ActionStatus    `json:"status"`          // The action status. Can be created, assigned, started, completed, canceled, suspended
	SubStatus      ActionSubStatus `json:"sub_status"`      // Granular action status. Can be not_started, in_queue (for assigned status) leaving_now, on_the_way, arriving, arrived (for started status)
	InitialEta     *time.Time      `json:"initial_eta"`     // The initial eta calculated for the action
	Eta            *time.Time      `json:"eta"`             // The current eta for the action, calculated in real-time
	CompletedPlace Place           `json:"completed_place"` // The place where the action was completed
	CompletedAt    *time.Time      `json:"completed_at"`    // The time at which the action was completed
	TrackingUrl    string          `json:"tracking_url"`    // The URL at which this action can be tracked
	CreatedAt      time.Time       `json:"created_at"`      // Timestamp of when action was created
	ModifiedAt     time.Time       `json:"modified_at"`     // Timestamp of when action was modified
	Distance       float64         `json:"distance"`        // Distance travelled by the user
}

type ActionType string

const (
	ActionTypePickup   = ActionType("pickup")
	ActionTypeDelivery = ActionType("delivery")
	ActionTypeDropoff  = ActionType("dropoff")
	ActionTypeVisit    = ActionType("visit")
	ActionTypeStopover = ActionType("stopover")
	ActionTypeTask     = ActionType("task")
)

type ActionStatus string

const (
	ActionStatusCreated   = ActionStatus("created")
	ActionStatusAssigned  = ActionStatus("assigned")
	ActionStatusStarted   = ActionStatus("started")
	ActionStatusCompleted = ActionStatus("completed")
	ActionStatusCanceled  = ActionStatus("canceled")
	ActionStatusSuspended = ActionStatus("suspended")
)

type ActionSubStatus string

const (
	ActionSubStatusNotStarted = ActionSubStatus("not_started")
	ActionSubStatusInQueue    = ActionSubStatus("in_queue") // (for assigned status)
	ActionSubStatusLeavingNow = ActionSubStatus("leaving_now")
	ActionSubStatusOnTheWay   = ActionSubStatus("on_the_way")
	ActionSubStatusArriving   = ActionSubStatus("arriving")
	ActionSubStatusArrived    = ActionSubStatus("arrived") // (for started status)
)
