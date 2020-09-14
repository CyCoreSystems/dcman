package common

// Equipment describes a piece of equipment, be it a server, a switch, a router, etc.
type Equipment struct {

	// ID is the unique identifier of the piece of equipment.
	ID string `db:"id"`

	// CustomerID is the unique identifier of the customer who owns this equipment
	CustomerID string `db:"customer_id"`

	// Name is the human-readable name for this piece of equipment
	Name string `db:"name"`

	// Active indicates whether this piece of equipment is active
	Active bool `db:"active"`

	// RackID indicates the unique identifier of the rack in which this equipment is housed.  If it is not installed, this will be empty.
	RackID string `db:"rack_id"`

	// BottomRU indicates the bottom rack unit number occupied by this piece of equipment.
	BottomRU int `db:"bottom_ru"`

	// Height is the number of rack units this unit occupies.
	Height int `db:"height"`
}
