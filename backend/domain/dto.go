package domain

// BaseResponse is the standard response format
type BaseResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Object  interface{} `json:"object,omitempty"`
	Errors  []string    `json:"errors,omitempty"`
}

// PaginatedResponse is for paginated lists
type PaginatedResponse struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	Object     interface{} `json:"object,omitempty"`
	PageNumber int         `json:"pageNumber"`
	PageSize   int         `json:"pageSize"`
	TotalSize  int64       `json:"totalSize"`
	Errors     []string    `json:"errors,omitempty"`
}

// AuthResponse for authentication endpoints
type AuthResponse struct {
	Success bool     `json:"success"`
	Message string   `json:"message"`
	Token   string   `json:"token,omitempty"`
	Errors  []string `json:"errors,omitempty"`
}

// Request DTOs
type SignupRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CreateMovieRequest struct {
	Title       string   `json:"title" binding:"required"`
	Description string   `json:"description" binding:"required"`
	Poster      string   `json:"poster" binding:"required"`
	Trailer     string   `json:"trailer" binding:"required"`
	Actors      []string `json:"actors" binding:"required"`
	Genres      []string `json:"genres" binding:"required"`
	UserID      string   `json:"-"`
}

type UpdateMovieRequest struct {
	Title       string   `json:"title" binding:"required"`
	Description string   `json:"description" binding:"required"`
	Poster      string   `json:"poster" binding:"required"`
	Trailer     string   `json:"trailer" binding:"required"`
	Actors      []string `json:"actors" binding:"required"`
	Genres      []string `json:"genres" binding:"required"`
}