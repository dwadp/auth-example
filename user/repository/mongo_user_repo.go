package repository

import (
	"context"

	"github.com/dwadp/auth-example/models"
	"github.com/dwadp/auth-example/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoUserRepository struct {
	ctx            context.Context
	userCollection *mongo.Collection
	primaryKey     string
}

func NewMongoUserRepository(
	db *mongo.Database,
	ctx context.Context) user.Repository {
	return &mongoUserRepository{
		ctx:            ctx,
		userCollection: db.Collection("users"),
		primaryKey:     "_id",
	}
}

func (m *mongoUserRepository) GetAll() ([]models.User, error) {
	users := []models.User{}

	cur, err := m.userCollection.Find(m.ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	defer cur.Close(m.ctx)

	for cur.Next(m.ctx) {
		user := models.User{}

		if err := cur.Decode(&user); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (m *mongoUserRepository) Store(user models.User) error {
	_, err := m.userCollection.InsertOne(m.ctx, bson.M{
		"name":     user.Name,
		"email":    user.Email,
		"password": user.Password,
	})

	return err
}

func (m *mongoUserRepository) GetByEmail(email string) (models.User, error) {
	return m.getBy(bson.M{
		"email": email,
	})
}

func (m *mongoUserRepository) GetByID(id string) (models.User, error) {
	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return models.User{}, nil
	}

	return m.getBy(bson.M{
		m.primaryKey: objectID,
	})
}

func (m *mongoUserRepository) getBy(filter bson.M) (models.User, error) {
	var result models.User

	cur := m.userCollection.FindOne(m.ctx, filter)

	if err := cur.Decode(&result); err != nil {
		return result, err
	}

	return result, nil
}
