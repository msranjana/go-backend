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
