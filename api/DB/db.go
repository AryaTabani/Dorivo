package db

import (
	"database/sql"
	"encoding/json"
	"log"

	"github.com/AryaTabani/Dorivo/models"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to database")
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	createTables()
	createDefaultTenant()
}

func createTables() {
	createTenantsTable := `
    CREATE TABLE IF NOT EXISTS tenants (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL UNIQUE, 
        config TEXT NOT NULL
    );`

	_, err := DB.Exec(createTenantsTable)
	if err != nil {
		panic("Failed to create tenants table: " + err.Error())
	}

	createUsersTable := `
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    tenant_id TEXT NOT NULL, 
    full_name TEXT NOT NULL,
    email TEXT NOT NULL,
    mobile_number TEXT,
    password_hash TEXT NOT NULL,
    date_of_birth TEXT,
    FOREIGN KEY (tenant_id) REFERENCES tenants(name) ON DELETE CASCADE,,
    UNIQUE (tenant_id, email)
);`

	_, err = DB.Exec(createUsersTable)
	if err != nil {
		panic("Failed to create users table")
	}

	createPasswordResetsTable := `
	CREATE TABLE IF NOT EXISTS password_reset_tokens (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		token_hash TEXT NOT NULL UNIQUE,
		expires_at TIMESTAMP NOT NULL,
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
	);`

	_, err = DB.Exec(createPasswordResetsTable)
	if err != nil {
		panic("Failed to create password_reset_tokens table")
	}
	
}
func createDefaultTenant() {
	defaultTenant := "localhost:3000"
	var count int
	err := DB.QueryRow("SELECT COUNT(*) FROM tenants WHERE name = ?", defaultTenant).Scan(&count)
	if err != nil {
		panic("Failed to check for default tenant: " + err.Error())
	}

	if count == 0 {
		defaultConfig := models.TenantConfig{
			Name:         "دوریوو",
			Logo:         "/assets/logos/PWALogo-192x192.png",
			Plan:         models.PlanPro,
			MultiTheme:   true,
			DefaultTheme: models.ThemeLight,
			Features:     json.RawMessage(`{}`),
			ThemeColors: models.ThemeColors{
				Primary:    "163 177 138",
				Primary2:   "230 240 230",
				Secondary:  "45 106 79",
				Secondary2: "205 234 192",
			},
		}

		configJSON, err := json.Marshal(defaultConfig)
		if err != nil {
			panic("Failed to marshal default tenant config: " + err.Error())
		}

		query := `INSERT INTO tenants (name, config) VALUES (?, ?)`
		_, err = DB.Exec(query, defaultTenant, string(configJSON))
		if err != nil {
			panic("Failed to create default tenant: " + err.Error())
		}
		log.Println(" Default tenant created successfully!")
	}
}
