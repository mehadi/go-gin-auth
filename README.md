# Go Gin Authentication API

A secure authentication API built with Go, Gin, and PostgreSQL. This project implements user registration, login, JWT-based authentication, and CRUD operations for posts.

## Project Structure

```
go-gin-auth-api-starter-kit/
├── cmd/
│   ├── server/
│   │   └── main.go          # Application entry point
│   └── seeder/
│       └── main.go          # Database seeder command
├── config/
│   └── config.go            # Database configuration
├── controllers/
│   ├── auth_controller.go   # Authentication handlers
│   ├── post_controller.go   # Post management handlers
│   └── user_controller.go   # User management handlers
├── middleware/
│   └── auth_middleware.go   # JWT authentication middleware
├── models/
│   ├── user.go              # User data model
│   └── post.go              # Post data model
├── repositories/
│   ├── user_repository.go   # User database operations
│   └── post_repository.go   # Post database operations
├── services/
│   ├── auth_service.go      # Authentication business logic
│   └── post_service.go      # Post business logic
├── utils/
│   ├── hash.go              # Password hashing
│   └── token.go             # JWT token handling
└── docker-compose.yml       # Docker configuration
```

## Features

- User Registration and Login
- JWT-based Authentication
- Protected Routes with Middleware
- CRUD Operations for Posts
- Password Hashing
- Database Seeding
- Docker Support

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
     - GET `/api/v1/posts` - List all posts (protected)
     - POST `/api/v1/posts` - Create new post (protected)
     - GET `/api/v1/posts/:id` - Get post by ID (protected)
     - PUT `/api/v1/posts/:id` - Update post (protected)
     - DELETE `/api/v1/posts/:id` - Delete post (protected)

3. **Middleware** (`middleware/auth_middleware.go`)
   - Validates JWT tokens
   - Extracts user claims
   - Sets user context
   - Handles unauthorized access

4. **Controllers** (`controllers/`)
   - `auth_controller.go`: Handles registration and login
   - `user_controller.go`: Handles user listing and management
   - `post_controller.go`: Handles post CRUD operations

5. **Services** (`services/`)
   - `auth_service.go`: Authentication business logic
   - `post_service.go`: Post business logic
   - Handles password hashing
   - Manages JWT token generation
   - Calls repository layer

6. **Repositories** (`repositories/`)
   - `user_repository.go`: User database operations
   - `post_repository.go`: Post database operations
   - Performs CRUD operations
   - Handles data persistence

7. **Models** (`models/`)
   - `user.go`: User data model
   - `post.go`: Post data model
   - Defines data structures
   - Maps to database tables
   - Contains validation rules

8. **Utils**
   - `hash.go`: Securely hashes passwords
   - `token.go`: Generates and validates JWT tokens

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

## API Endpoints

### Authentication

#### Register
```bash
curl -X POST http://localhost:8080/api/v1/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "email": "test@example.com",
    "password": "password123"
  }'
```

**Request Body:**
```json
{
  "username": "testuser",
  "email": "test@example.com",
  "password": "password123"
}
```

#### Login
```bash
curl -X POST http://localhost:8080/api/v1/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "password": "password123"
  }'
```

**Request Body:**
```json
{
  "email": "test@example.com",
  "password": "password123"
}
```

### Posts (All endpoints require authentication)

#### List Posts
```bash
curl -X GET http://localhost:8080/api/v1/posts \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

**Response:**
```json
{
  "posts": [
    {
      "id": 1,
      "title": "First Post",
      "content": "Content of first post",
      "created_at": "2024-01-01 12:00:00"
    }
  ]
}
```

#### Create Post
```bash
curl -X POST http://localhost:8080/api/v1/posts \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "My First Post",
    "content": "This is the content of my first post"
  }'
```

**Request Body:**
```json
{
  "title": "My First Post",
  "content": "This is the content of my first post"
}
```

#### Get Post by ID
```bash
curl -X GET http://localhost:8080/api/v1/posts/1 \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

**Response:**
```json
{
  "post": {
    "id": 1,
    "title": "First Post",
    "content": "Content of first post",
    "created_at": "2024-01-01 12:00:00"
  }
}
```

#### Update Post
```bash
curl -X PUT http://localhost:8080/api/v1/posts/1 \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Updated Title",
    "content": "Updated content"
  }'
```

**Request Body:**
```json
{
  "title": "Updated Title",
  "content": "Updated content"
}
```

#### Delete Post
```bash
curl -X DELETE http://localhost:8080/api/v1/posts/1 \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### Users (Protected)

#### List Users
```bash
curl -X GET http://localhost:8080/api/v1/users \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

**Response:**
```json
{
  "users": [
    {
      "id": 1,
      "username": "admin",
      "email": "admin@example.com",
      "created_at": "2024-01-01 12:00:00"
    }
  ]
}
```

#### Dashboard
```bash
curl -X GET http://localhost:8080/api/v1/dashboard \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

**Response:**
```json
{
  "message": "Welcome to your dashboard!",
  "username": "your_username"
}
```

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

5. **Seed the Database (Optional)**
   ```bash
   # Normal seeding (only if no users exist)
   go run cmd/seeder/main.go

   # Force reseeding (deletes all users and creates new ones)
   go run cmd/seeder/main.go -force
   ```

## Default Users

The seeder creates these default users:
- Username: `admin`, Email: `admin@example.com`, Password: `admin123`
- Username: `user1`, Email: `user1@example.com`, Password: `user123`
- Username: `user2`, Email: `user2@example.com`, Password: `user123`

## Security Features

- Password hashing using bcrypt
- JWT token authentication with middleware
- Protected routes with middleware
- Secure database configuration
- Environment variable management
- Input validation
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