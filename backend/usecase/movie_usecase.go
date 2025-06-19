package usecase

import (
	"context"

	"github.com/AfomiaTadesse/Afomia_M/backend/domain"
	"github.com/AfomiaTadesse/Afomia_M/backend/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
type MovieUsecase interface {
	CreateMovie(req *domain.CreateMovieRequest) (*domain.BaseResponse, error)
	GetMovies(page, size int) (*domain.PaginatedResponse, error)  // This was missing
	GetMovieByID(id string) (*domain.BaseResponse, error)
	SearchMovies(title string, page, size int) (*domain.PaginatedResponse, error)
	UpdateMovie(id, userID string, req *domain.UpdateMovieRequest) (*domain.BaseResponse, error)
	DeleteMovie(id, userID string) (*domain.BaseResponse, error)
}

type movieUsecase struct {
	movieRepo repository.MovieRepository
}

func NewMovieUsecase(movieRepo repository.MovieRepository) MovieUsecase {
	return &movieUsecase{movieRepo: movieRepo}
}

func (uc *movieUsecase) CreateMovie(req *domain.CreateMovieRequest) (*domain.BaseResponse, error) {
	// Convert userID to ObjectID
	userID, err := primitive.ObjectIDFromHex(req.UserID)
	if err != nil {
		return &domain.BaseResponse{
			Success: false,
			Message: "Invalid user ID",
			Errors:  []string{err.Error()},
		}, nil
	}

	movie := &domain.Movie{
		Title:       req.Title,
		Description: req.Description,
		Poster:      req.Poster,
		Trailer:     req.Trailer,
		Actors:      req.Actors,
		Genres:      req.Genres,
		UserID:      userID,
	}

	if err := uc.movieRepo.Create(context.Background(), movie); err != nil {
		return nil, err
	}

	return &domain.BaseResponse{
		Success: true,
		Message: "Movie created successfully",
		Object:  movie,
	}, nil
}

func (uc *movieUsecase) SearchMovies(title string, page, size int) (*domain.PaginatedResponse, error) {
	movies, total, err := uc.movieRepo.SearchByTitle(context.Background(), title, page, size)
	if err != nil {
		return nil, err
	}

	return &domain.PaginatedResponse{
		Success:    true,
		Message:    "Movies retrieved successfully",
		Object:     movies,
		PageNumber: page,
		PageSize:   size,
		TotalSize:  total,
	}, nil
}

func (uc *movieUsecase) GetMovieByID(id string) (*domain.BaseResponse, error) {
	movie, err := uc.movieRepo.GetByID(context.Background(), id)
	if err != nil {
		return &domain.BaseResponse{
			Success: false,
			Message: "Movie not found",
		}, nil
	}

	return &domain.BaseResponse{
		Success: true,
		Message: "Movie retrieved successfully",
		Object:  movie,
	}, nil
}

func (uc *movieUsecase) UpdateMovie(id, userID string, req *domain.UpdateMovieRequest) (*domain.BaseResponse, error) {
	// Verify movie exists and belongs to user
	movie, err := uc.movieRepo.GetByID(context.Background(), id)
	if err != nil {
		return &domain.BaseResponse{
			Success: false,
			Message: "Movie not found",
		}, nil
	}

	// Check ownership
	if movie.UserID.Hex() != userID {
		return &domain.BaseResponse{
			Success: false,
			Message: "You are not authorized to update this movie",
		}, nil
	}

	// Update movie fields
	updatedMovie := &domain.Movie{
		Title:       req.Title,
		Description: req.Description,
		Poster:      req.Poster,
		Trailer:     req.Trailer,
		Actors:      req.Actors,
		Genres:      req.Genres,
		UserID:      movie.UserID, // Keep original user ID
	}

	err = uc.movieRepo.Update(context.Background(), id, updatedMovie)
	if err != nil {
		return nil, err
	}

	return &domain.BaseResponse{
		Success: true,
		Message: "Movie updated successfully",
		Object:  updatedMovie,
	}, nil
}

func (uc *movieUsecase) DeleteMovie(id, userID string) (*domain.BaseResponse, error) {
	// Verify movie exists and belongs to user
	movie, err := uc.movieRepo.GetByID(context.Background(), id)
	if err != nil {
		return &domain.BaseResponse{
			Success: false,
			Message: "Movie not found",
		}, nil
	}

	// Check ownership
	if movie.UserID.Hex() != userID {
		return &domain.BaseResponse{
			Success: false,
			Message: "You are not authorized to delete this movie",
		}, nil
	}

	err = uc.movieRepo.Delete(context.Background(), id)
	if err != nil {
		return nil, err
	}

	return &domain.BaseResponse{
		Success: true,
		Message: "Movie deleted successfully",
	}, nil
}
func (uc *movieUsecase) GetMovies(page, size int) (*domain.PaginatedResponse, error) {
	movies, total, err := uc.movieRepo.GetAll(context.Background(), page, size)
	if err != nil {
		return nil, err
	}

	return &domain.PaginatedResponse{
		Success:    true,
		Message:    "Movies retrieved successfully",
		Object:     movies,
		PageNumber: page,
		PageSize:   size,
		TotalSize:  total,
	}, nil
}