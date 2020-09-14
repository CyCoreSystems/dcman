package common

// Customer describes a customer which may be invoiced
type Customer struct {

	// ID is the unique identifier of the customer.
	ID string `db:"id"`

	// InvoiceID is the unique identifier of the customer in the invoicing system
	InvoiceID string `db:"invoice_id"`

	// Name is the unique human-readable name for the customer
	Name string `db:"name"`

	// Active indicates whether this customer is active
	Active bool `db:"active"`
}
