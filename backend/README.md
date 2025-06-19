# Movie Collection API

A RESTful API for managing personal movie collections, built with Go, MongoDB, and JWT authentication.

## Features

- User authentication (Signup/Login with JWT)
- CRUD operations for movies
- Paginated movie listings
- Search functionality
- Secure password storage (bcrypt)

## Technologies

- **Backend**: Go (Gin framework)
- **Database**: MongoDB
- **Authentication**: JWT

## API Endpoints

### Authentication
| Method | Endpoint          | Description                     |
|--------|-------------------|---------------------------------|
| POST   | `/api/v1/auth/users/signup` | Register a new user           |
| POST   | `/api/v1/auth/users/login`  | Login and get JWT token      |

### Movies
| Method | Endpoint                   | Description                     |
|--------|----------------------------|---------------------------------|
| GET    | `/api/v1/movies`           | Get paginated list of movies    |
| GET    | `/api/v1/movies/search`    | Search movies by title          |
| GET    | `/api/v1/movies/:id`       | Get movie details               |
| POST   | `/api/v1/movies`           | Create a new movie (Auth)       |
| PUT    | `/api/v1/movies/:id`       | Update a movie (Auth)           |
| DELETE | `/api/v1/movies/:id`       | Delete a movie (Auth)           |

## Installation

### Prerequisites
- Go 1.16+
- MongoDB

### Steps
1. Clone the repository:
   ```bash
   git clone https://github.com/AfomiaTadesse/Afomia_M.git
   cd Afomia_M/backend