# Dorivo - Multi-Tenant Ordering System Backend

![Go Version](https://img.shields.io/badge/Go-1.21+-blue.svg)
![Docker](https://img.shields.io/badge/Docker-Ready-blue?logo=docker)
![Swagger](https://img.shields.io/badge/API_Docs-Swagger-orange?logo=swagger)
![License](https://img.shields.io/badge/License-MIT-green.svg)

**Dorivo** is a comprehensive, scalable, and secure backend service built with Go for multi-tenant ordering applications, such as food delivery platforms or e-commerce stores. It is designed with a clean 3-tier architecture, is fully containerized with Docker, and includes auto-generated API documentation with Swagger.

---

## ✨ Key Features

This backend provides a complete, production-ready set of features:

* **Multi-Level, Multi-Tenant Architecture**:
    * **Multi-Tenancy**: A core architectural feature allowing multiple businesses (tenants) to use the same application instance with completely isolated data.
    * **Role-Based Access Control (RBAC)**: A robust permission system with three distinct roles: `CUSTOMER`, `ADMIN` (for tenants), and `SUPER_ADMIN` (for the platform owner).

* **Admin Panel (Tenant-Level)**:
    * Full CRUD (Create, Read, Update, Delete) management for products.
    * Dashboard for viewing all tenant-specific orders and updating their status.
    * Ability to update their own tenant's configuration (name, theme, etc.).
    * Read-only access to their customer list.
    * Analytics dashboard with key metrics like total revenue and daily orders.

* **Super Admin Panel (Platform-Level)**:
    * Secure, separate login for the platform owner.
    * Full CRUD (Create, Read, Update, Delete) management for all tenants on the platform.

* **Full Product & Menu Management**:
    * Dynamic product searching and filtering.
    * Support for product customizations with option groups and add-ons.
    * Curated product lists for **Best Sellers**, **Promotions**, and **Chef's Recommendations**.

* **Complete Shopping Cart & Checkout System**:
    * Persistent shopping cart for each user.
    * Support for promo codes and discounts.
    * Transactional order creation to ensure data integrity.

* **Comprehensive User Profile & Order Management**:
    * Full control over user profiles, addresses, and payment methods.
    * Detailed order history and ability to leave reviews.
    * Secure password change and account deletion.

* **Automated API Documentation**:
    * Live, interactive API documentation is automatically generated using **Swagger**, making it easy for frontend developers to understand and test the API.

---

## 🏗️ Architecture

The project follows a classic **3-Tier Architecture** to ensure a clean separation of concerns, making the codebase modular, scalable, and easy to maintain.

* **Controllers (Presentation Layer)**: Located in the `/controllers` directory. Responsible for handling HTTP requests, parsing input, and formatting JSON responses.
* **Services (Business Logic Layer)**: Located in the `/services` directory. Contains the core application logic, enforces business rules, and orchestrates data flow.
* **Repositories (Data Access Layer)**: Located in the `/repository` directory. This layer is responsible for all communication with the database via SQL queries.

---

## 🚀 Tech Stack

* **Language**: **Go**
* **Web Framework**: **Gin** (for routing and HTTP handling)
* **Database**: **MySQL**
* **Authentication**: **JWT (JSON Web Tokens)**
* **Password Hashing**: **Bcrypt**
* **Containerization**: **Docker & Docker Compose**
* **API Documentation**: **Swagger (swaggo)**

---

## 🏁 Getting Started

Follow these instructions to get the project and its database running on your local machine.

### Prerequisites

* **Git** for cloning the repository.
* **Docker Desktop** installed and running on your machine.

### 🐳 Running with Docker (Recommended)

This is the simplest and most reliable way to run the entire application stack.

1.  **Clone the repository:**
    ```sh
    git clone https://github.com/AryaTabani/Dorivo.git
    cd Dorivo
    ```

2.  **Set up environment variables:**
    Create a **`.env`** file in the root of the project. This file stores your JWT secret and database credentials.
    ```env
    # JWT Secret Key
    JWT_SECRET_KEY="your-super-secret-key-that-is-long-and-secure"

    # MySQL Connection Details for Docker Compose
    DB_HOST=db
    DB_PORT=3306
    DB_NAME=dorivo_db
    DB_USER=dorivo_user
    DB_PASSWORD=strongpassword123
    DB_ROOT_PASSWORD=superstrongrootpassword
    ```

3.  **Build and run the containers:**
    This single command will start both the Go application and a persistent MySQL database container.
    ```sh
    docker-compose up --build
    ```

* The server will be available at `http://localhost:8080`.
* The interactive API documentation will be available at `http://localhost:8080/swagger/index.html`.

---

## 📄 API Endpoints

A complete, interactive list of all API endpoints is available via the auto-generated Swagger documentation. After running the application, navigate to:

**[http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)**



The documentation provides detailed information on every endpoint, including public routes, authenticated user routes, and the protected routes for the **Admin** and **Super Admin** panels.

---

## 💡 Future Improvements

* **Add Unit & Integration Tests**: Implement a comprehensive test suite to ensure code quality and reliability.
* **Real Payment Gateway**: Integrate with a real payment processor like Stripe.
* **Caching**: Implement a caching layer with Redis to improve performance for frequently accessed data.
* **Deployment to Cloud**: Prepare and document the process for deploying to a cloud provider like Vercel, AWS, or Google Cloud.
