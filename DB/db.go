package db2

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/AryaTabani/Dorivo/models"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		panic("Could not initialize database connection: " + err.Error())
	}
	err = DB.Ping()
	if err != nil {
		panic("Database connection is not available: " + err.Error())
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	createTables()
	createDefaultTenant()
}

func createTables() {
	var err error
	createTenantsTable := `
    CREATE TABLE IF NOT EXISTS tenants (
        id INT PRIMARY KEY AUTO_INCREMENT,
        name VARCHAR(255) NOT NULL UNIQUE,
        config JSON NOT NULL
    );`
	_, err = DB.Exec(createTenantsTable)
	if err != nil {
		panic("Failed to create tenants table: " + err.Error())
	}

	createUsersTable := `
    CREATE TABLE IF NOT EXISTS users (
        id INT PRIMARY KEY AUTO_INCREMENT,
        tenant_id VARCHAR(150) NOT NULL,
        role VARCHAR(50) NOT NULL DEFAULT 'CUSTOMER',
        full_name VARCHAR(255) NOT NULL,
        email VARCHAR(150) NOT NULL,
        mobile_number VARCHAR(255),
        password_hash VARCHAR(255) NOT NULL,
        date_of_birth TEXT,
        avatar_url VARCHAR(255),
        notification_preference JSON,
        FOREIGN KEY (tenant_id) REFERENCES tenants(name) ON DELETE CASCADE,
        UNIQUE (tenant_id, email)
    );`
	_, err = DB.Exec(createUsersTable)
	if err != nil {
		panic("Failed to create users table: " + err.Error())
	}
	createSuperAdminsTable := `
    CREATE TABLE IF NOT EXISTS super_admins (
    id INT PRIMARY KEY AUTO_INCREMENT,
    email VARCHAR(191) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL
);`
	_, err = DB.Exec(createSuperAdminsTable)
	if err != nil {
		panic("Failed to create super_admins table: " + err.Error())
	}

	createUserAddressesTable := `
    CREATE TABLE IF NOT EXISTS user_addresses (
        id INT PRIMARY KEY AUTO_INCREMENT,
        user_id INT NOT NULL,
        name VARCHAR(255) NOT NULL,
        address TEXT NOT NULL,
        FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
    );`
	_, err = DB.Exec(createUserAddressesTable)
	if err != nil {
		panic("Failed to create user_addresses table: " + err.Error())
	}

	createPasswordResetsTable := `
    CREATE TABLE IF NOT EXISTS password_reset_tokens (
        id INT PRIMARY KEY AUTO_INCREMENT,
        user_id INT NOT NULL,
        token_hash VARCHAR(255) NOT NULL UNIQUE,
        expires_at TIMESTAMP NOT NULL,
        FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
    );`
	_, err = DB.Exec(createPasswordResetsTable)
	if err != nil {
		panic("Failed to create password_reset_tokens table: " + err.Error())
	}

	createOrdersTable := `
    CREATE TABLE IF NOT EXISTS orders (
        id INT PRIMARY KEY AUTO_INCREMENT,
        user_id INT NOT NULL,
        tenant_id VARCHAR(191) NOT NULL,
        status VARCHAR(255) NOT NULL,
        total_price DECIMAL(10, 2) NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
        FOREIGN KEY (tenant_id) REFERENCES tenants(name) ON DELETE CASCADE
    );`
	_, err = DB.Exec(createOrdersTable)
	if err != nil {
		panic("Failed to create orders table: " + err.Error())
	}

	createOrderItemsTable := `
    CREATE TABLE IF NOT EXISTS order_items (
        id INT PRIMARY KEY AUTO_INCREMENT,
        order_id INT NOT NULL,
        item_name VARCHAR(255) NOT NULL,
        quantity INT NOT NULL,
        price DECIMAL(10, 2) NOT NULL,
        image_url VARCHAR(255),
        FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE
    );`
	_, err = DB.Exec(createOrderItemsTable)
	if err != nil {
		panic("Failed to create order_items table: " + err.Error())
	}

	createCancellationTable := `
    CREATE TABLE IF NOT EXISTS cancellations (
        id INT PRIMARY KEY AUTO_INCREMENT,
        order_id INT NOT NULL,
        user_id INT NOT NULL,
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
        id INT PRIMARY KEY AUTO_INCREMENT,
        order_id INT NOT NULL,
        user_id INT NOT NULL,
        rating INT NOT NULL CHECK(rating >= 1 AND rating <= 5),
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
        id INT PRIMARY KEY AUTO_INCREMENT,
        user_id INT NOT NULL,
        processor_token VARCHAR(255) NOT NULL UNIQUE,
        card_brand VARCHAR(255) NOT NULL,
        last_four VARCHAR(4) NOT NULL,
        expiry_month INT NOT NULL,
        expiry_year INT NOT NULL,
        is_default TINYINT(1) DEFAULT 0,
        FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
    );`
	_, err = DB.Exec(createPaymentMethodsTable)
	if err != nil {
		panic("Failed to create payment_methods table: " + err.Error())
	}

	createFaqsTable := `
    CREATE TABLE IF NOT EXISTS faqs (
        id INT PRIMARY KEY AUTO_INCREMENT,
        tenant_id VARCHAR(191) NOT NULL,
        category VARCHAR(255) NOT NULL,
        question TEXT NOT NULL,
        answer TEXT NOT NULL,
        FOREIGN KEY (tenant_id) REFERENCES tenants(name) ON DELETE CASCADE
    );`
	_, err = DB.Exec(createFaqsTable)
	if err != nil {
		panic("Failed to create faqs table: " + err.Error())
	}

	createNotificationsTable := `
    CREATE TABLE IF NOT EXISTS notifications (
        id INT PRIMARY KEY AUTO_INCREMENT,
        user_id INT NOT NULL,
        title VARCHAR(255) NOT NULL,
        type VARCHAR(255) NOT NULL,
        is_read TINYINT(1) DEFAULT 0,
        metadata JSON,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
    );`
	_, err = DB.Exec(createNotificationsTable)
	if err != nil {
		panic("Failed to create notifications table: " + err.Error())
	}

	createProductsTable := `
    CREATE TABLE IF NOT EXISTS products (
        id INT PRIMARY KEY AUTO_INCREMENT,
        tenant_id VARCHAR(191) NOT NULL,
        name VARCHAR(255) NOT NULL,
        description TEXT,
        price DECIMAL(10, 2) NOT NULL,
        rating DECIMAL(3, 2) DEFAULT 0,
        image_url VARCHAR(255),
        main_category VARCHAR(255) NOT NULL,
        discount_price DECIMAL(10, 2),
        is_featured TINYINT(1) DEFAULT 0,
        is_recommended TINYINT(1) DEFAULT 0,
        FOREIGN KEY (tenant_id) REFERENCES tenants(name) ON DELETE CASCADE
    );`
	_, err = DB.Exec(createProductsTable)
	if err != nil {
		panic("Failed to create products table: " + err.Error())
	}

	createTagsTable := `
    CREATE TABLE IF NOT EXISTS tags (
        id INT PRIMARY KEY AUTO_INCREMENT,
        tenant_id VARCHAR(150) NOT NULL,
        name VARCHAR(150) NOT NULL,
        main_category VARCHAR(255) NOT NULL,
        UNIQUE(tenant_id, name)
    );`
	_, err = DB.Exec(createTagsTable)
	if err != nil {
		panic("Failed to create tags table: " + err.Error())
	}

	createProductTagsTable := `
    CREATE TABLE IF NOT EXISTS product_tags (
        product_id INT NOT NULL,
        tag_id INT NOT NULL,
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
        id INT PRIMARY KEY AUTO_INCREMENT,
        product_id INT NOT NULL,
        name VARCHAR(255) NOT NULL,
        selection_type VARCHAR(255) NOT NULL,
        FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
    );`
	_, err = DB.Exec(createOptionGroupsTable)
	if err != nil {
		panic("Failed to create option_groups table: " + err.Error())
	}

	createOptionsTable := `
    CREATE TABLE IF NOT EXISTS options (
        id INT PRIMARY KEY AUTO_INCREMENT,
        option_group_id INT NOT NULL,
        name VARCHAR(255) NOT NULL,
        price_modifier DECIMAL(10, 2) NOT NULL DEFAULT 0,
        FOREIGN KEY (option_group_id) REFERENCES option_groups(id) ON DELETE CASCADE
    );`
	_, err = DB.Exec(createOptionsTable)
	if err != nil {
		panic("Failed to create options table: " + err.Error())
	}

	createCartsTable := `
    CREATE TABLE IF NOT EXISTS carts (
        id INT PRIMARY KEY AUTO_INCREMENT,
        user_id INT NOT NULL UNIQUE,
        FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
    );`
	_, err = DB.Exec(createCartsTable)
	if err != nil {
		panic("Failed to create carts table: " + err.Error())
	}

	createCartItemsTable := `
    CREATE TABLE IF NOT EXISTS cart_items (
        id INT PRIMARY KEY AUTO_INCREMENT,
        cart_id INT NOT NULL,
        product_id INT NOT NULL,
        quantity INT NOT NULL,
        FOREIGN KEY (cart_id) REFERENCES carts(id) ON DELETE CASCADE,
        FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
    );`
	_, err = DB.Exec(createCartItemsTable)
	if err != nil {
		panic("Failed to create cart_items table: " + err.Error())
	}

	createCartItemOptionsTable := `
    CREATE TABLE IF NOT EXISTS cart_item_options (
        cart_item_id INT NOT NULL,
        option_id INT NOT NULL,
        PRIMARY KEY (cart_item_id, option_id),
        FOREIGN KEY (cart_item_id) REFERENCES cart_items(id) ON DELETE CASCADE,
        FOREIGN KEY (option_id) REFERENCES options(id) ON DELETE CASCADE
    );`
	_, err = DB.Exec(createCartItemOptionsTable)
	if err != nil {
		panic("Failed to create cart_item_options table: " + err.Error())
	}

	createUserFavoritesTable := `
    CREATE TABLE IF NOT EXISTS user_favorites (
        user_id INT NOT NULL,
        product_id INT NOT NULL,
        PRIMARY KEY (user_id, product_id),
        FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
        FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
    );`
	_, err = DB.Exec(createUserFavoritesTable)
	if err != nil {
		panic("Failed to create user_favorites table: " + err.Error())
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
			Features:     models.RawJSONObject{"featureX": true, "featureY": false},
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
