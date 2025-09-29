package models

type Product struct {
	ID           int64         `json:"id"`
	TenantID     string        `json:"-"`
	Name         string        `json:"name"`
	Description  string        `json:"description"`
	Price        float64       `json:"price"`
	Rating       float64       `json:"rating"`
	ImageURL     string        `json:"image_url"`
	MainCategory string        `json:"main_category"`
	Tags         []string      `json:"tags,omitempty"`
	OptionGroups []OptionGroup `json:"option_groups,omitempty"`
}

type Tag struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	MainCategory string `json:"main_category"`
}
type Option struct {
	ID            int64   `json:"id"`
	Name          string  `json:"name"`
	PriceModifier float64 `json:"price_modifier"`
}
type OptionGroup struct {
	ID            int64    `json:"id"`
	Name          string   `json:"name"`
	SelectionType string   `json:"selection_type"`
	Options       []Option `json:"options"`
}
