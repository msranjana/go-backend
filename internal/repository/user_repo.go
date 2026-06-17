package repository

import (
	"context"
	"time"

	db "user-dob-api/db/sqlc"
)

type UserRepository struct {
	q *db.Queries
}

func NewUserRepository(q *db.Queries) *UserRepository {
	return &UserRepository{q: q}
}

func (r *UserRepository) Create(ctx context.Context, name string, dob time.Time) (db.User, error) {
	return r.q.CreateUser(ctx, name, dob)
}

func (r *UserRepository) GetByID(ctx context.Context, id int32) (db.User, error) {
	return r.q.GetUser(ctx, id)
}

func (r *UserRepository) Update(ctx context.Context, id int32, name string, dob time.Time) (db.User, error) {
	return r.q.UpdateUser(ctx, name, dob, id)
}

func (r *UserRepository) Delete(ctx context.Context, id int32) error {
	return r.q.DeleteUser(ctx, id)
}

func (r *UserRepository) List(ctx context.Context, limit, offset int32) ([]db.User, error) {
	return r.q.ListUsersPaginated(ctx, limit, offset)
}

func (r *UserRepository) Count(ctx context.Context) (int32, error) {
	return r.q.CountUsers(ctx)
}