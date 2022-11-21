package user

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
)

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
	AppName  = "User"
)

type Service interface {
	CreateUser(context.Context, *User) (*User, error)
	QueryUser(context.Context, *QueryUserRequest) (*UserSet, error)
	UpdateUser(context.Context, *User) (*User, error)
	DescribeUser(context.Context, *DescribeUserRequest) (*User, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*User, error)
}

func NewQueryUserRequestFromHTTP(r *http.Request) *QueryUserRequest {
	qs := r.URL.Query()

	ps := qs.Get("page_size")
	pn := qs.Get("page_number")
	kw := qs.Get("keywords")

	psUint64, _ := strconv.ParseUint(ps, 10, 64)
	pnUint64, _ := strconv.ParseUint(pn, 10, 64)

	if psUint64 == 0 {
		psUint64 = 20
	}
	if pnUint64 == 0 {
		pnUint64 = 1
	}
	return &QueryUserRequest{
		PageSize:   psUint64,
		PageNumber: pnUint64,
		Keywords:   kw,
	}
}

type QueryUserRequest struct {
	PageSize   uint64 `json:"page_size,omitempty"`
	PageNumber uint64 `json:"page_number,omitempty"`
	Keywords   string `json:"keywords"`
}

func (req *QueryUserRequest) OffSet() int64 {
	return int64(req.PageSize) * int64(req.PageNumber-1)
}

func NewDescribeUserRequestWithID(user_name string) *DescribeUserRequest {
	return &DescribeUserRequest{
		user_name: user_name,
	}
}

type DescribeUserRequest struct {
	user_name string `json:"user_name" validate:"required"`
}

func NewDeleteUserRequestWithID(user_name string) *DeleteUserRequest {
	return &DeleteUserRequest{user_name: user_name}
}

type DeleteUserRequest struct {
	user_name string `json:"user_name" validate:"required"`
}
