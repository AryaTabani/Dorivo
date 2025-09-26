package models

import "encoding/json"

type Plan string
type Theme string

const (
	PlanBase Plan = "BASE"
	PlanPro  Plan = "PRO"
	PlanVip  Plan = "VIP"

	ThemeDark  Theme = "dark"
	ThemeLight Theme = "light"
)

type ThemeColors struct {
	Primary    string `json:"primary"`
	Primary2   string `json:"primary2"`
	Secondary  string `json:"secondary"`
	Secondary2 string `json:"secondary2"`
}

type TenantConfig struct {
	Name         string          `json:"name"`
	Logo         string          `json:"logo,omitempty"`
	Plan         Plan            `json:"plan"`
	MultiTheme   bool            `json:"multiTheme"`
	DefaultTheme Theme           `json:"defaultTheme"`
	ThemeColors  ThemeColors     `json:"themeColors"`
	Features     json.RawMessage `json:"features"`
	ContactInfo  ContactInfo     `json:"contactInfo"`
}

type Tenant struct {
	ID     int64
	Name   string
	Config TenantConfig
}

type ContactInfo struct {
	CustomerService string `json:"customerService"`
	Website         string `json:"website"`
	Whatsapp        string `json:"whatsapp"`
	Facebook        string `json:"facebook"`
	Instagram       string `json:"instagram"`
}
