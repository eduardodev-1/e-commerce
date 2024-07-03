package models

type Product struct {
	ID          int64   `db:"id"`
	Name        string  `db:"name"`
	Description string  `db:"description"`
	ImgURL      string  `db:"img_url"`
	Price       float64 `db:"price"`
}
