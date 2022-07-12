package service

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"user/domain/model"
	"user/domain/repository"
)

type IUserDataService interface {
	AddUser(*model.User)(int64, error)
	DeleteUser(int64)error
	UpdateUser(user *model.User, isChangePwd bool)(err error)
	FindUserByName(string)(*model.User, error)
	CheckPwd(userName string, pwd string)(isOk bool,err error)
}

//此处声明了UserDataService 实现了IUserDataService接口
func NewUserDataService(userRepository repository.UserRepository)IUserDataService{
	return &UserDataService{UserRepository: userRepository}
}
type UserDataService struct {
	UserRepository repository.IUserRepository
}

//加密用户密码
func GeneratePassword(userPassword string) ([]byte, error){
	return bcrypt.GenerateFromPassword([]byte(userPassword),bcrypt.DefaultCost)
}

func ValidatePassword(userPassword, hashed string)(isOk bool, err error) {
	if err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(userPassword));err!=nil{
		return false, errors.New("用户名或密码错误！")
	}
	return true, nil
}

//更新用户
func (u UserDataService) AddUser(user *model.User) (int64, error) {
	pwdByte, err := GeneratePassword(user.HashPassword)
	if err != nil {
		return user.ID, nil
	}
	user.HashPassword = string(pwdByte)
	return u.UserRepository.CreateUser(user)

}

//删除用户
func (u UserDataService) DeleteUser(userID int64) error {
	return u.UserRepository.DeleteUserByID(userID)
}

//更新用户
func (u UserDataService) UpdateUser(user *model.User, isChangePwd bool) (err error) {
	//判断是否更新了密码
	if isChangePwd{
		pwdByte, err := GeneratePassword(user.HashPassword)
		if err !=nil{
			return err
		}
		user.HashPassword = string(pwdByte)
		//TODO rabbitmq
	}
	return u.UserRepository.UpdateUser(user)

}

//根据名字查找用户
func (u UserDataService) FindUserByName(userName string) (*model.User, error) {
	return u.UserRepository.FindUserByName(userName)
}

//比对账号密码是否正确
func (u UserDataService) CheckPwd(userName string, pwd string) (isOk bool, err error) {
	user, err := u.UserRepository.FindUserByName(userName)
	if err != nil {
		return false, err
	}
	return ValidatePassword(pwd, user.HashPassword)
}
