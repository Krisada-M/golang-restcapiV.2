package services

import (
	"api/models"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceFunc struct {
	usercollection *mongo.Collection
	ctx            context.Context
}

func NewUserService(usercollection *mongo.Collection, ctx context.Context) UserService {
	return &UserServiceFunc{
		usercollection: usercollection,
		ctx:            ctx,
	}
}

func (u *UserServiceFunc) CreateUser(user *models.User) error {
	_, err := u.usercollection.InsertOne(u.ctx, user)
	return err
}

func (u *UserServiceFunc) UpdateUser(user *models.User, name *string) error {
	filter := bson.D{bson.E{Key: "user_name", Value: name}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{
		bson.E{Key: "user_name", Value: user.Name},
		bson.E{Key: "user_age", Value: user.Age},
		bson.E{Key: "user_birthday", Value: user.Birthday},
		bson.E{Key: "user_address", Value: user.Address}}}}
	result, _ := u.usercollection.UpdateOne(u.ctx, filter, update)
	if result.MatchedCount != 1 {
		return errors.New("User Data Not Match Found For Update")
	}
	return nil
}

func (u *UserServiceFunc) DeleteUser(user *string) error {
	filter := bson.D{bson.E{Key: "user_name", Value: user}}
	result, _ := u.usercollection.DeleteOne(u.ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("User Data Not Match Found For Delete")
	}
	return nil
}

func (u *UserServiceFunc) GetOnlyUser(user *string) (*models.User, error) {
	var userdata *models.User
	query := bson.D{bson.E{Key: "user_name", Value: user}}
	err := u.usercollection.FindOne(u.ctx, query).Decode(&userdata)

	return userdata, err
}

func (u *UserServiceFunc) GetAllUser() ([]*models.User, error) {
	var users []*models.User
	data, err := u.usercollection.Find(u.ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for data.Next(u.ctx) {
		var user models.User
		err := data.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	if err := data.Err(); err != nil {
		return nil, err
	}
	data.Close(u.ctx)

	if len(users) == 0 {
		return nil, errors.New("Data Not Found")
	}

	return users, nil
}
