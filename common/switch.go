package common

// Switch describes a network switch (or router) in the data center
type Switch struct {

	// ID is the unique equipment identifier for this switch
	ID string `db:"id"`

	// Ports is the list of ports on this switch
	Ports []*SwitchPort
}

// SwitchPort describes a port on a switch
type SwitchPort struct {

	// ID is the unique identifier of this switch port
	ID string `db:"id"`

	// SwitchID is the unique identifier of the switch equipment to which this port belongs.
	SwitchID string `db:"switch_id"`

	// Name is the switch-unique identifier of this port on the switch
	Name string `db:"name"`

	// Index is the numeric index of this port on the switch
	Index int `db:"index"`

	// ConnectionID is the unique identifier of the connection attached to this port, if there is one
	ConnectionID string `db:"connection_id"`

	// MaxSpeed is the highest operational speed of this port, in bits per second
	MaxSpeed int64 `db:"max_speed"`
}
