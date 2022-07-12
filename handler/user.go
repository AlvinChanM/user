package handler

import (
	"context"
	"user/domain/model"
	"user/domain/service"
	user "user/proto/user"
)

type User struct{
	UserDataService service.IUserDataService
}


//注册
func (u *User) Register(ctx context.Context, req *user.RegisterRequest,
	res *user.RegisterResponse) error{
	user2 := &model.User{
		UserName:     req.UserName,
		FirstName:    req.FirstName,
		HashPassword: req.Pwd,
	}

	_, err := u.UserDataService.AddUser(user2)
	if err != nil {
		return err
	}
	res.Message = "注册成功！"
	return nil
}


func (u *User) Login(ctx context.Context, req user.LoginRequest, res user.LoginResponse) error {
	isOk, err := u.UserDataService.CheckPwd(req.UserName, req.Pwd)
	if err != nil{
		return err
	}
	res.IsSuccess = isOk
	return nil
}

// 查询用户信息
func (u *User) GetUserInfo(ctx context.Context, req user.UserInfoRequest,
	res *user.UserInfoResponse) error{
	userInfo, err := u.UserDataService.FindUserByName(req.UserName)
	if err != nil{
		return err
	}
	res = UserForResponse(userInfo)
	return nil
}


//类型转换
func UserForResponse(userModel *model.User) *user.UserInfoResponse{
	return &user.UserInfoResponse{
		UserId:               userModel.ID,
		UserName:             userModel.UserName,
		FirstName:            userModel.FirstName,
	}
}