package pack

import (
	"errors"
	"tiktok-backend/kitex_gen/user"
	"tiktok-backend/pkg/errno"
)

// BuildUserLoginResp build baseResp from error
func BuildUserLoginResp(err error) *user.DouyinUserLoginResponse {
	if err == nil {
		return userLoginResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return userLoginResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return userLoginResp(s)
}

func userLoginResp(err errno.ErrNo) *user.DouyinUserLoginResponse {
	return &user.DouyinUserLoginResponse{StatusCode: err.ErrCode, StatusMsg: err.ErrMsg}
}

// BuildUserRegisterResp build baseResp from error
func BuildUserRegisterResp(err error) *user.DouyinUserRegisterResponse {
	if err == nil {
		return userRegisterResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return userRegisterResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return userRegisterResp(s)
}

func userRegisterResp(err errno.ErrNo) *user.DouyinUserRegisterResponse {
	return &user.DouyinUserRegisterResponse{StatusCode: err.ErrCode, StatusMsg: err.ErrMsg}
}

// BuildUserInfoResp build baseResp from error
func BuildUserInfoResp(err error) *user.DouyinUserResponse {
	if err == nil {
		return userInfoResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return userInfoResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return userInfoResp(s)
}

func userInfoResp(err errno.ErrNo) *user.DouyinUserResponse {
	return &user.DouyinUserResponse{StatusCode: err.ErrCode, StatusMsg: err.ErrMsg}
}


