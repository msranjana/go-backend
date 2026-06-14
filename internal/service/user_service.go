package service

import (
	"context"
	"time"

	"user-dob-api/internal/models"
	"user-dob-api/internal/repository"
)

const dobLayout = "2006-01-02"

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func calcAge(dob time.Time) int {
	now := time.Now()
	age := now.Year() - dob.Year()
	if now.YearDay() < dob.YearDay() {
		age--
	}
	return age
}

func (s *UserService) Create(ctx context.Context, req models.CreateUserRequest) (models.UserResponse, error) {
	dob, err := time.Parse(dobLayout, req.DOB)
	if err != nil {
		return models.UserResponse{}, err
	}
	u, err := s.repo.Create(ctx, req.Name, dob)
	if err != nil {
		return models.UserResponse{}, err
	}
	return models.UserResponse{ID: u.ID, Name: u.Name, DOB: u.Dob.Format(dobLayout)}, nil
}

func (s *UserService) GetByID(ctx context.Context, id int32) (models.UserWithAgeResponse, error) {
	u, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return models.UserWithAgeResponse{}, err
	}
	return models.UserWithAgeResponse{
		ID:   u.ID,
		Name: u.Name,
		DOB:  u.Dob.Format(dobLayout),
		Age:  calcAge(u.Dob),
	}, nil
}

func (s *UserService) Update(ctx context.Context, id int32, req models.UpdateUserRequest) (models.UserResponse, error) {
	dob, err := time.Parse(dobLayout, req.DOB)
	if err != nil {
		return models.UserResponse{}, err
	}
	u, err := s.repo.Update(ctx, id, req.Name, dob)
	if err != nil {
		return models.UserResponse{}, err
	}
	return models.UserResponse{ID: u.ID, Name: u.Name, DOB: u.Dob.Format(dobLayout)}, nil
}

func (s *UserService) Delete(ctx context.Context, id int32) error {
	return s.repo.Delete(ctx, id)
}

func (s *UserService) List(ctx context.Context) ([]models.UserWithAgeResponse, error) {
	users, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]models.UserWithAgeResponse, 0, len(users))
	for _, u := range users {
		result = append(result, models.UserWithAgeResponse{
			ID:   u.ID,
			Name: u.Name,
			DOB:  u.Dob.Format(dobLayout),
			Age:  calcAge(u.Dob),
		})
	}
	return result, nil
}
