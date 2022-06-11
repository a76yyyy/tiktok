package command

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"

	"github.com/a76yyyy/tiktok/pkg/errno"

	"github.com/a76yyyy/tiktok/cmd/user/kitex_gen/user"

	"github.com/a76yyyy/tiktok/cmd/user/dal/db"
)

type CreateUserService struct {
	ctx context.Context
}

// NewCreateUserService new CreateUserService
func NewCreateUserService(ctx context.Context) *CreateUserService {
	return &CreateUserService{ctx: ctx}
}

// CreateUser create user info.
func (s *CreateUserService) CreateUser(req *user.DouyinUserRegisterRequest) error {
	users, err := db.QueryUser(s.ctx, req.Username)
	if err != nil {
		return err
	}
	if len(users) != 0 {
		return errno.ErrUserAlreadyExist
	}

	h := md5.New()
	if _, err = io.WriteString(h, req.Password); err != nil {
		return err
	}
	passWord := fmt.Sprintf("%x", h.Sum(nil))
	return db.CreateUser(s.ctx, []*db.User{{
		UserName: req.Username,
		Password: passWord,
	}})
}
