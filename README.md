# Dorivo - Multi-Tenant Ordering System Backend

![Go Version](https://img.shields.io/badge/Go-1.18+-blue.svg)
![License](https://img.shields.io/badge/License-MIT-green.svg)

**Dorivo** is a comprehensive, scalable, and secure backend service built with Go for multi-tenant ordering applications, such as food delivery platforms, e-commerce stores, or restaurant management systems. It is designed with a clean 3-tier architecture to be maintainable and easy to extend.

---

## ✨ Key Features

This backend provides a complete set of features required for a modern ordering application:

* **Multi-Tenancy**: A core architectural feature allowing multiple businesses (tenants) to use the same application instance with completely isolated configurations, products, and orders.
* **Secure Authentication**: Standard JWT-based authentication for user registration, login, and secure access to protected routes. Passwords are securely hashed using bcrypt.
* **Full Product & Menu Management**:
    * Dynamic product searching and filtering based on categories, tags, price, and ratings.
    * Support for product customizations with option groups and add-ons (e.g., "Toppings," "Select a size").
    * Curated product lists for **Best Sellers**, **Promotions**, and **Chef's Recommendations**.
* **Complete Shopping Cart System**:
    * Persistent shopping cart for each user.
    * Accurate price calculation, including product options and modifiers.
* **Order Management**: View order history with different statuses (Active, Completed, Cancelled), cancel orders, and leave reviews for completed orders.
* **Comprehensive User Profile Management**:
    * Update user profile information (name, avatar, etc.).
    * Manage multiple delivery addresses and payment methods.
    * Granular notification settings.
    * Secure password change and account deletion.
* **Favorites System**: Allows users to save and view their favorite products.
* **Notification Center**: A system for storing and retrieving user-specific notifications triggered by application events (e.g., order status changes).
* **Tenant-Specific Content**: Each tenant can configure their own contact information and Help/FAQ sections.

---

## 🏗️ Architecture

The project follows a classic **3-Tier Architecture** to ensure a clean separation of concerns, making the codebase modular, scalable, and easy to maintain.



* **Controllers (Presentation Layer)**: Located in the `/controllers` directory. Responsible for handling HTTP requests, parsing input (payloads, URL parameters), and formatting JSON responses.
* **Services (Business Logic Layer)**: Located in the `/services` directory. Contains the core application logic, enforces business rules, performs calculations, and orchestrates data flow between controllers and repositories.
* **Repositories (Data Access Layer)**: Located in the `/repository` directory. This layer is responsible for all communication with the database. It abstracts the SQL queries and database logic from the rest of the application.

---

## 🚀 Tech Stack

* **Language**: **Go**
* **Web Framework**: **Gin** (for routing and HTTP handling)
* **Database**: **SQLite** (for simplicity and ease of setup)
* **Authentication**: **JWT (JSON Web Tokens)**
* **Password Hashing**: **Bcrypt**

---

## データベーススキーマ

The database is designed with a relational structure to ensure data integrity.



**Core Tables:**
* `tenants`: Stores unique configurations for each business.
* `users`: Manages user accounts, profiles, and settings.
* `products`, `tags`, `option_groups`, `options`: A comprehensive schema for the product catalog and its customizations.
* `carts`, `cart_items`, `cart_item_options`: Manages the state of user shopping carts.
* `orders`, `order_items`: Stores historical order data.
* `user_favorites`, `notifications`, `reviews`, `cancellations`, and more.

---

## 🏁 Getting Started

Follow these instructions to get the project running on your local machine.

### Prerequisites

* Go (version 1.18 or higher)
* Git

### Installation & Setup

1.  **Clone the repository:**
    ```sh
    git clone https://github.com/AryaTabani/Dorivo.git
    cd Dorivo
    ```

2.  **Install dependencies:**
    ```sh
    go mod tidy
    ```

3.  **Set up environment variables:**
    Create a `.env` file in the root of the project. The application requires a secret key for signing JWTs.
    ```env
    JWT_SECRET_KEY="your-super-secret-key-that-is-long-and-secure"
    ```
    *Note: The application will not run without this key.*

4.  **Run the application:**
    ```sh
    go run main.go
    ```
    The server will start on `http://localhost:8080`.

---

##  API Endpoints

All API responses are wrapped in a standard JSON object:
`{ "success": true/false, "data": ..., "message": ..., "error": ... }`

### Public Routes

| Method | Endpoint                                         | Description                                    |
| :----- | :----------------------------------------------- | :--------------------------------------------- |
| `GET`  | `/tenant/:tenantId`                              | Get the configuration for a specific tenant.   |
| `POST` | `/:tenantId/register`                            | Register a new user for a tenant.              |
| `POST` | `/:tenantId/login`                               | Log in a user and get a JWT.                   |
| `GET`  | `/:tenantId/faqs`                                | Get the FAQs for a tenant.                     |
| `GET`  | `/:tenantId/products`                            | Search and filter products.                    |
| `GET`  | `/:tenantId/tags`                                | Get all available filter tags for a tenant.    |
| `GET`  | `/:tenantId/products/:productId`                 | Get detailed information about a single product. |
| `GET`  | `/:tenantId/products/bestsellers`                | Get a list of best-selling products.           |
| `GET`  | `/:tenantId/products/featured`                   | Get the main featured/promotional product.     |
| `GET`  | `/:tenantId/products/recommended`                | Get a list of recommended products.            |

### Authenticated Routes

*(Requires `Authorization: Bearer <JWT>` header)*

| Method   | Endpoint                               | Description                                        |
| :------- | :------------------------------------- | :------------------------------------------------- |
| **User & Profile** |
| `GET`    | `/profile`                             | Get the current user's profile.                    |
| `PUT`    | `/profile`                             | Update the current user's profile.                 |
| `GET`    | `/addresses`                           | Get the user's saved addresses.                    |
| `POST`   | `/addresses`                           | Add a new address.                                 |
| `DELETE` | `/addresses/:addressId`                | Delete an address.                                 |
| `GET`    | `/payment-methods`                     | Get the user's saved payment methods.              |
| `POST`   | `/payment-methods`                     | Add a new payment method.                          |
| `DELETE` | `/payment-methods/:methodId`           | Delete a payment method.                           |
| `PUT`    | `/profile/change-password`             | Change the user's password.                        |
| `DELETE` | `/profile`                             | Delete the user's account.                         |
| **Favorites** |
| `GET`    | `/favorites`                           | Get all of the user's favorite products.           |
| `POST`   | `/products/:productId/favorite`        | Add a product to favorites.                        |
| `DELETE` | `/products/:productId/favorite`        | Remove a product from favorites.                   |
| **Cart** |
| `GET`    | `/cart`                                | Get the contents of the user's shopping cart.      |
| `POST`   | `/cart/items`                          | Add an item to the cart.                           |
| `PUT`    | `/cart/items/:itemId`                  | Update an item's quantity in the cart.             |
| `DELETE` | `/cart/items/:itemId`                  | Remove an item from the cart.                      |
| **Orders** |
| `GET`    | `/orders`                              | Get the user's order history.                      |
| `POST`   | `/orders/:orderId/cancel`              | Cancel an active order.                            |
| `POST`   | `/orders/:orderId/review`              | Leave a review for a completed order.              |
| **Notifications & Settings** |
| `GET`    | `/notifications`                       | Get the user's notifications.                      |
| `PUT`    | `/notifications/read`                  | Mark notifications as read.                        |
| `GET`    | `/profile/notification-settings`       | Get the user's notification preferences.           |
| `PUT`    | `/profile/notification-settings`       | Update the user's notification preferences.        |

---

## 💡 Future Improvements

* **Switch to PostgreSQL**: Migrate from SQLite to a more robust, production-ready database like PostgreSQL.
* **Add Unit & Integration Tests**: Implement a comprehensive test suite to ensure code quality and reliability.
* **Real Payment Gateway**: Integrate with a real payment processor like Stripe.
* **Containerization**: Add a `Dockerfile` to containerize the application for easy deployment.
* **Caching**: Implement a caching layer with Redis to improve performance for frequently accessed data.

---

