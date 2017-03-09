package Goods

import (
	. "Gyl/Module/Roll"
	. "Gyl/Module/System"
)

type Product struct {
	ProductID         ProductID
	SupplyerID        UserID
	ProductName       string
	ProductType       string
	ProductUnitWeight int64
	ProductUnitPrice  int64
	PorductDiscount   int64
	ProductWarehouse  []Warehouse
}

func (p *Product) NewProduct() Product {
	//TODO:
	return Product{}
}

func (p *Product) AddWarehouse(warehouse Warehouse) error {
	//TODO:
	return nil
}
