package models

type Product struct {
	ID           int64    `json:"id"`
	TenantID     string   `json:"-"`
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	Price        float64  `json:"price"`
	Rating       float64  `json:"rating"`
	ImageURL     string   `json:"image_url"`
	MainCategory string   `json:"main_category"`
	Tags         []string `json:"tags,omitempty"`
}

type Tag struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	MainCategory string `json:"main_category"`
}
