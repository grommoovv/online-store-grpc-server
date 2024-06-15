package domain

type Product struct {
	ID          int64   `db:"id"`
	Title       string  `db:"title"`
	Description string  `db:"description"`
	Price       float32 `db:"price"`
	ImageURL    string  `db:"image_url"`
	Category    string  `db:"category"`
}
