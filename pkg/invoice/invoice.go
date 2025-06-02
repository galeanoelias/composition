package invoice

import (
	"github.com/galeanoelias/composition/pkg/customer"
	"github.com/galeanoelias/composition/pkg/invoiceItem"
)

// Invoice is the struc of a invoice
type Invoice struct {
	country string
	city    string
	total   float64
	client  customer.Customer
	items []invoiceitem.Item
}

// New return a new Invoice
func New(country, city string, client customer.Customer, items []invoiceitem.Item) Invoice {
	return Invoice{
		country: country,
		city: city,
		client: client,
		items: items,
	}
}

// SetTotal is the setter of Invoice.total
func (i *Invoice) SetTotal() {
	for _, item := range i.items {
		i.total += item.Value()
	}
}