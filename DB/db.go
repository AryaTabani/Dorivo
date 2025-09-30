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
	DB, err = sql.Open("sqlite3", "/tmp/api.db")
	//DB, err = sql.Open("sqlite3", "api.db")

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
    avatar_url TEXT, 
	notification_preference TEXT,
    FOREIGN KEY (tenant_id) REFERENCES tenants(name) ON DELETE CASCADE,
    UNIQUE (tenant_id, email)
);`

	_, err = DB.Exec(createUsersTable)
	if err != nil {
		panic("Failed to create users table: " + err.Error())
	}
	createUserAddressesTable := `
    CREATE TABLE IF NOT EXISTS user_addresses (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    name TEXT NOT NULL,
    address TEXT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);`
	_, err = DB.Exec(createUserAddressesTable)
	if err != nil {
		panic("Failed to create user_addresses table: " + err.Error())
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
		panic("Failed to create password_reset_tokens table: " + err.Error())
	}

	createOrdersTabale := `
	CREATE TABLE IF NOT EXISTS orders (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	user_id INTEGER NOT NULL,
	tenant_id TEXT NOT NULL,
	status TEXT NOT NULL,
	total_price REAL NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY (user_id) REFERENCES users(id),
	FOREIGN KEY (tenant_id) REFERENCES tenants(name)
	);`
	_, err = DB.Exec(createOrdersTabale)
	if err != nil {
		panic("Failed to create orders table: " + err.Error())
	}
	createOrderItemsTable := `
	CREATE TABLE IF NOT EXISTS order_items(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	order_id INTEGER NOT NULL,
	item_name TEXT NOT NULL,
	quantity INTEGER NOT NULL,
	price REAL NOT NULL,
	image_url TEXT,
	FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE
	);`

	_, err = DB.Exec(createOrderItemsTable)
	if err != nil {
		panic("Failed to create password_reset_tokens table: " + err.Error())
	}
	createCancellationTable := `
	CREATE TABLE IF NOT EXISTS cancellations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		order_id INTEGER NOT NULL,
		user_id INTEGER NOT NULL,
		reason TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE,
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
    );`
	_, err = DB.Exec(createCancellationTable)
	if err != nil {
		panic("Failed to create cancellations table: " + err.Error())
	}
	createReviewsTable := `
	CREATE TABLE IF NOT EXISTS reviews (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		order_id INTEGER NOT NULL,
		user_id INTEGER NOT NULL,
		rating INTEGER NOT NULL CHECK(rating >= 1 AND rating <= 5),
		comment TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE,
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
    );`
	_, err = DB.Exec(createReviewsTable)
	if err != nil {
		panic("Failed to create reviews table: " + err.Error())
	}
	createPaymentMethodsTable := `
    CREATE TABLE IF NOT EXISTS payment_methods (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    processor_token TEXT NOT NULL UNIQUE,
    card_brand TEXT NOT NULL,
    last_four TEXT NOT NULL,
    expiry_month INTEGER NOT NULL,
    expiry_year INTEGER NOT NULL,
    is_default BOOLEAN DEFAULT FALSE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);`
	_, err = DB.Exec(createPaymentMethodsTable)
	if err != nil {
		panic("Failed to create paymentsMethod table: " + err.Error())
	}
	createfraqsTable := `
	CREATE TABLE IF NOT EXISTS faqs (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		tenant_id TEXT NOT NULL,
		category TEXT NOT NULL,
		question TEXT NOT NULL,
		answer TEXT NOT NULL,
		FOREIGN KEY (tenant_id) REFERENCES tenants(name) ON DELETE CASCADE
	);`
	_, err = DB.Exec(createfraqsTable)
	if err != nil {
		panic("Failed to create faqs table: " + err.Error())
	}

	createNotificationsTable := `
	CREATE TABLE IF NOT EXISTS notifications (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		title TEXT NOT NULL,
		type TEXT NOT NULL,
		is_read BOOLEAN DEFAULT FALSE,
		metadata TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
	);`
	_, err = DB.Exec(createNotificationsTable)
	if err != nil {
		panic("Failed to create notifications table: " + err.Error())
	}
	createProductsTable := `
	CREATE TABLE IF NOT EXISTS products (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		tenant_id TEXT NOT NULL,
		name TEXT NOT NULL,
		description TEXT,
		price REAL NOT NULL,
		rating REAL DEFAULT 0,
		image_url TEXT,
		main_category TEXT NOT NULL,
        discount_price REAL, 
        is_featured BOOLEAN DEFAULT FALSE, 
		is_recommended BOOLEAN DEFAULT FALSE,
		FOREIGN KEY (tenant_id) REFERENCES tenants(name) ON DELETE CASCADE
	);`
	_, err = DB.Exec(createProductsTable)
	if err != nil {
		panic("Failed to create products table: " + err.Error())
	}

	createTagsTable := `
	CREATE TABLE IF NOT EXISTS tags (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		tenant_id TEXT NOT NULL,
		name TEXT NOT NULL,
		main_category TEXT NOT NULL,
		UNIQUE (tenant_id, name),
		FOREIGN KEY (tenant_id) REFERENCES tenants(name) ON DELETE CASCADE
	);`
	_, err = DB.Exec(createTagsTable)
	if err != nil {
		panic("Failed to create tags table: " + err.Error())
	}
	createProductTagsTable := `
	CREATE TABLE IF NOT EXISTS product_tags (
		product_id INTEGER NOT NULL,
		tag_id INTEGER NOT NULL,
		PRIMARY KEY (product_id, tag_id),
		FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE,
		FOREIGN KEY (tag_id) REFERENCES tags(id) ON DELETE CASCADE	
	);`
	_, err = DB.Exec(createProductTagsTable)
	if err != nil {
		panic("Failed to create product_tags table: " + err.Error())
	}
	createOptionGroupsTable := `
	CREATE TABLE IF NOT EXISTS option_groups (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		product_id INTEGER NOT NULL,
		name TEXT NOT NULL,
		selection_type TEXT NOT NULL,
		FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
	);`
	_, err = DB.Exec(createOptionGroupsTable)
	if err != nil {
		panic("Failed to create option_groups table: " + err.Error())
	}
	createOptionsTable := `
	CREATE TABLE IF NOT EXISTS options (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		option_group_id INTEGER NOT NULL,
		name TEXT NOT NULL,
		price_modifier REAL NOT NULL DEFAULT 0,
		FOREIGN KEY (option_group_id) REFERENCES option_groups(id) ON DELETE CASCADE
	);`
	_, err = DB.Exec(createOptionsTable)

	if err != nil {
		panic("Failed to create options table: " + err.Error())
	}
	createCartsTable := `
	CREATE TABLE IF NOT EXISTS carts (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
	);`
	_, err = DB.Exec(createCartsTable)
	if err != nil {
		panic("Failed to create carts table: " + err.Error())
	}
	createCartItemsTable := `
	CREATE TABLE IF NOT EXISTS cart_items (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		cart_id INTEGER NOT NULL,
		product_id INTEGER NOT NULL,
		quantity INTEGER NOT NULL,
		FOREIGN KEY (cart_id) REFERENCES carts(id) ON DELETE CASCADE,
		FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
	);`
	_, err = DB.Exec(createCartItemsTable)
	if err != nil {
		panic("Failed to create cart_items table: " + err.Error())
	}
	createCartItemOptionsTable := `
	CREATE TABLE IF NOT EXISTS cart_item_options (
		cart_item_id INTEGER NOT NULL,
		option_id INTEGER NOT NULL,
		PRIMARY KEY (cart_item_id, option_id),
		FOREIGN KEY (cart_item_id) REFERENCES cart_items(id) ON DELETE CASCADE,
		FOREIGN KEY (option_id) REFERENCES options(id) ON DELETE CASCADE
	);`
	_, err = DB.Exec(createCartItemOptionsTable)
	if err != nil {
		panic("Failed to create cart_item_options table: " + err.Error())
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
			ContactInfo: models.ContactInfo{
				CustomerService: "support@example.com",
				Website:         "https://www.example.com",
				Whatsapp:        "+1234567890",
				Facebook:        "https://facebook.com/example",
				Instagram:       "https://instagram.com/example",
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
