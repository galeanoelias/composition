package main

import (
	"fmt"

	"github.com/galeanoelias/composition/pkg/customer"
	"github.com/galeanoelias/composition/pkg/invoice"
	"github.com/galeanoelias/composition/pkg/invoiceItem"
)

func main() {
	i := invoice.New(
		"Argentina",
		"Avellaneda",
		customer.New("Elias", "calle flasa 123", "24454141"),
		// []invoiceitem.Item{
		// 	invoiceitem.New(1, "Cuaderno", 20.55),
		// 	invoiceitem.New(2, "Tijera", 10.34),
		// 	invoiceitem.New(3, "Plasticola", 34.12),
		// },
		invoiceitem.NewItems(
			invoiceitem.New(1, "Cuaderno", 20.55),
			invoiceitem.New(2, "Tijera", 10.34),
			invoiceitem.New(3, "Plasticola", 34.12),
		),
	)
	i.SetTotal()

	fmt.Printf("%+v", i)
}