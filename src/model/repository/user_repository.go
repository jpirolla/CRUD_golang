package repository

import (
	"github.com/jpirolla/CRUD_golang/src/configuration/rest_err"
	"github.com/jpirolla/CRUD_golang/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserRepository(
	database *mongo.Database,
) userRepository {
	return *userRepository{
		database
	}
}

type userRepository struct {
	databaseConnection *mongo.Database
}

type UserRepository interface {
	CreateUser(
		userDomain model.UserDomainInterface,
	) (model.UserDomainInterface, *rest_err.RestErr)
}
