package models

type CreateUserRequest struct {
	Name string `json:"name" validate:"required"`
	DOB  string `json:"dob"  validate:"required"`
}

type UpdateUserRequest struct {
	Name string `json:"name" validate:"required"`
	DOB  string `json:"dob"  validate:"required"`
}

type UserResponse struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
	DOB  string `json:"dob"`
}

type UserWithAgeResponse struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
	DOB  string `json:"dob"`
	Age  int    `json:"age"`
}

type PaginatedUsersResponse struct {
	Data       []UserWithAgeResponse `json:"data"`
	Total      int32                 `json:"total"`
	Page       int32                 `json:"page"`
	Limit      int32                 `json:"limit"`
	TotalPages int32                 `json:"total_pages"`
}