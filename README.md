// CHIRPY //

Chirpy is a Go-based backend server that provides API endpoints for a chirping platform. This platform supports user creation, chirp posting, login and refresh authentication, webhooks, metrics, and more. It also serves a front-end application from the ./app directory.

// FEATURES //
+ PostgreSQL database integration using SQLC for queries
+ User authentication with JWT
+ Secure password hashing with argon2id
+ API endpoints for creating chirps, users, login, refreshing tokens, and revoking tokens
+ Webhook handling (example route /api/polka/webhooks)
+ Metrics and reset endpoints under /admin
+ Static file serving for the frontend under /app
+ Health check endpoint for readiness

// REQUIREMENTS //
+ Go 1.25+
+ PostgreSQL database
+ Environment variables defined for:
    - DB_URL - PostgreSQL connection string
    - PLATFORM - Platform name or identifier
    - SECRET - JWT secret key
    - POLKA_KEY - API key for Polka webhook verification (optional)

// PROJECT STRUCTURE //
+ main.go - The main entry point for the server, setting up routes and middleware
+ sqlc.yaml - Configuration for SQLC code generation from SQL schema and queries
+ internal/database - Generated database query code via SQLC
+ functions - API handlers, middleware, and business logic (authentication, metrics)
+ app - Frontend static files served by the HTTP server

// RUNNING //
1. Set your environment variables:
    export DB_URL="your_postgres_connection_string"
    export PLATFORM="your_platform"
    export SECRET="your_jwt_secret"
    export POLKA_KEY="your_polka_api_key"

2. Ensure you have a running PostgreSQL instance with the correct schema.

3. Build and Run the server:
    go mod tidy
    go run main.go

4. The server listens on port 8080 by default.

// API ENDPOINTS //
+ POST /api/chirps - Create a new chirp
+ POST /api/users - Create a new user
+ POST /api/login - User login to get JWT token
+ POST /api/refresh - Refresh JWT token
+ POST /api/revoke - Revoke JWT token
+ POST /api/polka/webhooks - Handle Polka webhook events

+ GET /api/healthz - Health and readiness check
+ GET /api/chirps - Retrieve all chirps
    - Optional Query Parameters:
        - author_id (string of UUID): Returns chirps from the specified author only.
        - sort (string: "asc" or "desc"): Sort order of chirps by creation time. Defaults to "asc".
    - Examples:
        - /api/chirps
        - /api/chirps?author_id=uuid-string
        - /api/chirps?sort=desc
        - /api/chirps?author_id=uuid-string&sort=desc
+ GET /api/chirps/{chirpID} - Get a specific chirp
+ PUT /api/users - Update a user
+ DELETE /api/chirps/{chirpID} - Delete a chirp

// METRICS & ADMIN //
+ GET /admin/metrics - View server metrics
+ POST /admin/reset - Reset server state (metrics, data, etc.)

// DEPENDENCIES // 
+ Key Go dependencies are managed through go.mod and include:
    - github.com/alexedwards/argon2id - For secure password hashing
    - github.com/golang-jwt/jwt/v5 - JWT handling for authentication
    - github.com/google/uuid - UUID generation
    - github.com/joho/godotenv - Environment variable loading from .env
    - github.com/lib/pq - PostgreSQL driver