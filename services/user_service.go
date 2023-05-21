package services

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/prayogatriady/sawer-grpc/model"
	userPb "github.com/prayogatriady/sawer-grpc/model"
	"gorm.io/gorm"
)

type UserServiceInterface interface {
	GetUsers(ctx context.Context, void *empty.Empty) (*userPb.Users, error)
	GetUser(ctx context.Context, id *userPb.Id) (*userPb.User, error)
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

func (us *UserService) GetUser(ctx context.Context, id *userPb.Id) (*userPb.User, error) {

	// user, err := us.UserRepository.GetUser(ctx, int(id.GetId()))
	// if err != nil {
	// 	return nil, err
	// }
	// fmt.Printf("%v \n", user)

	var user *model.UserEntity
	if err := us.DB.WithContext(ctx).Table("users").Where("id =?", id.GetId()).Find(&user).Error; err != nil {
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
