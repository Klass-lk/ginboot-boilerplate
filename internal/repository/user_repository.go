package repository

import (
	"boilerplate/internal/model"
	"github.com/klass-lk/ginboot"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	*ginboot.MongoRepository[model.User]
}

func NewUserRepository(database *mongo.Database) *UserRepository {
	return &UserRepository{
		MongoRepository: ginboot.NewMongoRepository[model.User](database),
	}
}
