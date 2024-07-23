package models

type Product struct {
	ID          int64   `db:"id" json:"id,omitempty"`
	Name        string  `db:"name" json:"name,omitempty"`
	Description string  `db:"description" json:"description,omitempty"`
	ImgURL      string  `db:"img_url" json:"img_url,omitempty"`
	Price       float64 `db:"price" json:"price,omitempty"`
	SellerId    int64   `db:"seller" json:"seller_id,omitempty"`
	Quantity    int     `db:"quantity" json:"quantity,omitempty"`
}
