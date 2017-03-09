package Ticket

import (
	. "Gyl/Module/System"
)

type ProductPurchaseDetail struct {
	ProductID ProductID
	Amount    int64
}

func (u *ProductPurchaseDetail) NewProductPurchaseDetail() ProductPurchaseDetail {
	//TODO:
	return ProductPurchaseDetail{}
}
