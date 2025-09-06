# Go Multi-Tenant REST API ✨

A robust backend boilerplate built in Go, designed for multi-tenant applications. This project features a clean, layered architecture, secure JWT-based authentication for users, and a dynamic configuration system where each tenant can have a unique appearance and feature set.

![Project Architecture Diagram](https://storage.googleapis.com/gweb-cloud-storage-images/generative-ai/rest/layered-architecture-diagram.png)

---

## Features

* **Multi-Tenant Architecture:** Securely isolates user data and configurations based on a tenant identifier (passed via URL path).
* **JWT Authentication:** Secure, token-based authentication flow for user registration and login.
* **Tenant-Specific Configuration:** An endpoint to fetch tenant-specific UI configurations, including themes, logos, and feature flags.
* **Layered Architecture:** Organized into `controllers`, `services`, and `repository` layers for clean separation of concerns and maintainability.
* **Configuration for Serverless Deployment:** Includes the necessary setup (`vercel.json`) to deploy the application on Vercel.
* **Live Reloading:** Pre-configured for hot reloading during development with [Air](https://github.com/cosmtrek/air).

---

## Tech Stack

* **Language:** [Go](https://golang.org/)
* **Framework:** [Gin](https://gin-gonic.com/)
* **Database:** [SQLite](https://www.sqlite.org/) (using `mattn/go-sqlite3`)
* **Authentication:** `golang-jwt/jwt` for JSON Web Tokens
* **Password Hashing:** `golang.org/x/crypto/bcrypt`

---



## Project Structure

The project follows a standard layered architecture to keep code organized and maintainable.
