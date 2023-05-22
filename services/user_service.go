package services

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/prayogatriady/sawer-grpc/model"
	userPb "github.com/prayogatriady/sawer-grpc/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserServiceInterface interface {
	GetUsers(ctx context.Context, void *empty.Empty) (*userPb.Users, error)
	GetUser(ctx context.Context, userId *userPb.Id) (*userPb.User, error)
	CreateUser(ctx context.Context, userRequest *userPb.UserSignupRequest) (*userPb.Id, error)
	UpdateUser(ctx context.Context, userRequest *userPb.UserEditRequest) (*userPb.Status, error)
	DeleteUser(ctx context.Context, userId *userPb.Id) (*userPb.Status, error)
}

type UserService struct {
	userPb.UnimplementedUserServiceServer
	// UserRepository repository.UserRepoInterface
	DB *gorm.DB
}

// func NewUserService(userRepository repository.UserRepoInterface) UserServiceInterface {
// 	return &UserService{
// 		UserRepository: userRepository,
// 	}
// }

func (us *UserService) GetUsers(ctx context.Context, void *empty.Empty) (*userPb.Users, error) {

	// users, err := us.UserRepository.GetUsers(ctx)
	// if err != nil {
	// 	return nil, err
	// }
	// fmt.Printf("%v \n", users)

	var users []*model.UserEntity
	if err := us.DB.WithContext(ctx).Table("users").Find(&users).Error; err != nil {
		return nil, err
	}

	usersPb := &userPb.Users{Data: []*userPb.User{}}

	for _, user := range users {
		usersPb.Data = append(usersPb.Data, &userPb.User{
			Id:       uint64(user.ID),
			Username: user.Username,
			Password: user.Password,
			Balance:  int64(user.Balance),
		})
	}

	return usersPb, nil
}

func (us *UserService) GetUser(ctx context.Context, userId *userPb.Id) (*userPb.User, error) {

	// user, err := us.UserRepository.GetUser(ctx, int(id.GetId()))
	// if err != nil {
	// 	return nil, err
	// }
	// fmt.Printf("%v \n", user)

	user, err := GetUser(us.DB, ctx, int(userId.GetId()))
	if err != nil {
		return nil, err
	}

	userPb := &userPb.User{
		Id:       uint64(user.ID),
		Username: user.Username,
		Password: user.Password,
		Balance:  int64(user.Balance),
	}

	return userPb, nil
}

func (us *UserService) CreateUser(ctx context.Context, userRequest *userPb.UserSignupRequest) (*userPb.Id, error) {
	bytePassword, err := bcrypt.GenerateFromPassword([]byte(userRequest.GetPassword()), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	userEntity := &model.UserEntity{
		Username: userRequest.GetUsername(),
		Password: string(bytePassword),
	}

	if err := us.DB.WithContext(ctx).Table("users").Create(&userEntity).Error; err != nil {
		return nil, err
	}

	id := &userPb.Id{
		Id: uint64(userEntity.ID),
	}

	return id, nil
}

func (us *UserService) UpdateUser(ctx context.Context, userRequest *userPb.UserEditRequest) (*userPb.Status, error) {
	// get current entity
	userCurrentEntity, err := GetUser(us.DB, ctx, int(userRequest.GetId()))
	if err != nil {
		return nil, err
	}

	// hashing passsword
	bytePassword, err := bcrypt.GenerateFromPassword([]byte(userRequest.GetPassword()), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	userCurrentEntity.Password = string(bytePassword)
	userCurrentEntity.Balance = int(userRequest.GetBalance())

	var user *model.UserEntity
	if err := us.DB.WithContext(ctx).Table("users").Where("id =?", userRequest.GetId()).Updates(&userCurrentEntity).Find(&user).Error; err != nil {
		return nil, err
	}

	status := &userPb.Status{
		Status: 0,
	}

	return status, nil
}

func (us *UserService) DeleteUser(ctx context.Context, userId *userPb.Id) (*userPb.Status, error) {
	if err := us.DB.WithContext(ctx).Table("users").Where("id =?", userId.GetId()).Delete(&model.UserEntity{}).Error; err != nil {
		return nil, err
	}

	status := &userPb.Status{
		Status: 0,
	}

	return status, nil
}

func GetUser(db *gorm.DB, ctx context.Context, userId int) (*model.UserEntity, error) {
	var user *model.UserEntity
	if err := db.WithContext(ctx).Table("users").Where("id =?", userId).Find(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
