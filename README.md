# Go Gin Authentication API

A secure authentication API built with Go, Gin, and PostgreSQL. This project implements user registration, login, JWT-based authentication, and user management.

## Project Structure

```
go-gin-auth/
├── cmd/
│   ├── server/
│   │   └── main.go          # Application entry point
│   └── seeder/
│       └── main.go          # Database seeder command
├── config/
│   └── config.go            # Database configuration
├── controllers/
│   ├── auth_controller.go   # Authentication handlers
│   └── user_controller.go   # User management handlers
├── middleware/
│   └── auth_middleware.go   # JWT authentication middleware
├── models/
│   └── user.go              # User data model
├── pkg/
│   └── seeder/
│       └── seeder.go        # Database seeding logic
├── repositories/
│   └── user_repository.go   # Database operations
├── routes/
│   └── routes.go            # API route definitions
├── services/
│   └── auth_service.go      # Business logic
├── utils/
│   ├── hash.go              # Password hashing
│   └── token.go             # JWT token handling
└── docker-compose.yml       # Docker configuration
```

## Code Flow

1. **Entry Point** (`cmd/server/main.go`)
   - Starts the server
   - Loads environment variables
   - Sets up database connection
   - Configures routes

2. **Routes** (`routes/routes.go`)
   - Defines versioned API endpoints (v1):
     - POST `/api/v1/register` - Create new user
     - POST `/api/v1/login` - Authenticate user
     - GET `/api/v1/dashboard` - Protected dashboard
     - GET `/api/v1/users` - List all users (protected)

3. **Middleware** (`middleware/auth_middleware.go`)
   - Validates JWT tokens
   - Extracts user claims
   - Sets user context
   - Handles unauthorized access

4. **Controllers** (`controllers/`)
   - `auth_controller.go`: Handles registration and login
   - `user_controller.go`: Handles user listing and management

5. **Services** (`services/auth_service.go`)
   - Contains business logic
   - Handles password hashing
   - Manages JWT token generation
   - Calls repository layer

6. **Repositories** (`repositories/user_repository.go`)
   - Interacts with database
   - Performs CRUD operations
   - Handles data persistence

7. **Models** (`models/user.go`)
   - Defines data structures
   - Maps to database tables
   - Contains validation rules

8. **Utils**
   - `hash.go`: Securely hashes passwords
   - `token.go`: Generates and validates JWT tokens

9. **Seeder** (`pkg/seeder/seeder.go`)
   - Creates initial users
   - Handles database seeding
   - Supports force reseeding

## Authentication Flow

1. **Registration**
   ```
   Client -> POST /api/v1/register -> Controller -> Service -> Repository -> Database
   ```

   - Client sends user data (username, email, password)
   - Password is hashed
   - User is saved to database
   - Success response is returned

2. **Login**
   ```
   Client -> POST /api/v1/login -> Controller -> Service -> Repository -> Database
   ```

   - Client sends credentials (email, password)
   - Password is verified
   - JWT token is generated
   - Token is returned to client

3. **Protected Routes**
   ```
   Client -> GET /api/v1/dashboard -> Middleware -> Controller -> Service -> Response
   ```

   - Client includes JWT token in Authorization header
   - Middleware validates token
   - If valid, request proceeds to controller
   - If invalid, 401 Unauthorized is returned

## Database Seeding

The project includes a seeder to create initial users. To use it:

1. **Normal Seeding** (only if no users exist):
   ```bash
   go run cmd/seeder/main.go
   ```

2. **Force Reseeding** (deletes all users and creates new ones):
   ```bash
   go run cmd/seeder/main.go -force
   ```

The seeder creates these default users:
- Username: `admin`, Email: `admin@example.com`, Password: `admin123`
- Username: `user1`, Email: `user1@example.com`, Password: `user123`
- Username: `user2`, Email: `user2@example.com`, Password: `user123`

## Setup Instructions

1. **Prerequisites**
   - Go 1.16 or higher
   - Docker and Docker Compose
   - PostgreSQL

2. **Environment Variables**
   Create a `.env` file with:
   ```
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=your_password
   DB_NAME=auth_db
   JWT_SECRET=your_jwt_secret
   ```

3. **Database Setup**
   ```bash
   docker-compose up -d
   ```

4. **Run the Application**
   ```bash
   go run cmd/server/main.go
   ```

## API Endpoints

### Register
- **URL**: `/api/v1/register`
- **Method**: `POST`
- **Body**:
  ```json
  {
    "username": "testuser",
    "email": "test@example.com",
    "password": "password123"
  }
  ```

### Login
- **URL**: `/api/v1/login`
- **Method**: `POST`
- **Body**:
  ```json
  {
    "email": "test@example.com",
    "password": "password123"
  }
  ```

### Dashboard
- **URL**: `/api/v1/dashboard`
- **Method**: `GET`
- **Headers**:
  ```
  Authorization: Bearer <your_jwt_token>
  ```
- **Response**:
  ```json
  {
    "message": "Welcome to your dashboard!",
    "username": "your_username"
  }
  ```

### List Users
- **URL**: `/api/v1/users`
- **Method**: `GET`
- **Headers**:
  ```
  Authorization: Bearer <your_jwt_token>
  ```
- **Response**:
  ```json
  {
    "users": [
      {
        "id": 1,
        "username": "admin",
        "email": "admin@example.com",
        "created_at": "2024-01-01 12:00:00"
      },
      {
        "id": 2,
        "username": "user1",
        "email": "user1@example.com",
        "created_at": "2024-01-01 12:00:00"
      }
    ]
  }
  ```

## Security Features

- Password hashing using bcrypt
- JWT token authentication with middleware
- Secure database configuration
- Environment variable management
- Input validation
- Protected routes with middleware
- Password hash exclusion from responses

## Dependencies

- Gin: Web framework
- GORM: ORM for database operations
- JWT: Token generation and validation
- Bcrypt: Password hashing
- PostgreSQL: Database
- Docker: Containerization

## Best Practices

- Clean Architecture (Controllers -> Services -> Repositories)
- Secure password handling
- Environment variable configuration
- Error handling
- Code documentation
- Database migrations
- API versioning
- Middleware for authentication
- Database seeding for development 