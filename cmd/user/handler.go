package main

import (
	"context"
	"tiktok-backend/cmd/user/pack"
	"tiktok-backend/cmd/user/service"
	"tiktok-backend/kitex_gen/user"
	"tiktok-backend/pkg/errno"
	"tiktok-backend/pkg/jwt"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// UserLogin implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserLogin(ctx context.Context, req *user.DouyinUserLoginRequest) (resp *user.DouyinUserLoginResponse, err error) {
	resp = new(user.DouyinUserLoginResponse)

	if len(req.Username) == 0 || len(req.Password) == 0 {
		resp = pack.BuildUserLoginResp(errno.ParamErr)
		return resp, nil
	}

	uid, err := service.NewUserLoginService(ctx).UserLogin(req)
	if err != nil {
		resp = pack.BuildUserLoginResp(errno.ParamErr)
		return resp, nil
	}

	resp = pack.BuildUserLoginResp(errno.Success)
	resp.UserId = uid
	return resp, nil
}

// UserRegister implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserRegister(ctx context.Context, req *user.DouyinUserRegisterRequest) (resp *user.DouyinUserRegisterResponse, err error) {
	resp = new(user.DouyinUserRegisterResponse)

	if len(req.Username) == 0 || len(req.Password) == 0 {
		resp = pack.BuildUserRegisterResp(errno.ParamErr)
	}

	userId, err := service.NewUserRegisterService(ctx).UserRegister(req)
	if err != nil {
		resp = pack.BuildUserRegisterResp(err)
		return resp, nil
	}

	token, err := jwt.CreateToken(userId)
	if err != nil {
		resp = pack.BuildUserRegisterResp(err)
		return resp, nil
	}

	resp = pack.BuildUserRegisterResp(errno.Success)
	resp.UserId = userId
	resp.Token = token
	return resp, nil
}

// UserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserInfo(ctx context.Context, req *user.DouyinUserRequest) (resp *user.DouyinUserResponse, err error) {
	resp = new(user.DouyinUserResponse)

	userinfo, err := service.NewUserInfoService(ctx).UserInfo(req)
	if err != nil {
		resp = pack.BuildUserInfoResp(err)
		return resp, nil
	}

	resp = pack.BuildUserInfoResp(errno.Success)
	resp.User = userinfo
	return resp, nil
}
