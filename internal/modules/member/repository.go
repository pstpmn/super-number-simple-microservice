package member

import (
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const DB_NAME = "members-service"
const COLL_NAME = "members"

type (
	IMemberRepository interface {
		InsertOne(name, user, pass, email string) (primitive.ObjectID, error)
		FindOneByKey(key, value string) (*MemberEntity, error)
		IsUnigue(user, email string) bool
	}

	repo struct {
		db *mongo.Client
	}
)

// IsUnigue implements IMemberRepository.
func (r *repo) IsUnigue(user string, email string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	coll := r.db.Database(DB_NAME).Collection(COLL_NAME)
	result := new(MemberEntity)
	if err := coll.FindOne(
		ctx,
		bson.M{"$or": []bson.M{
			{"username": user},
			{"email": email},
		}},
	).Decode(result); err != nil {
		log.Printf("Error: IsUniquePlayer: %s", err.Error())
		return true
	}
	return false
}

func NewRepository(client *mongo.Client) IMemberRepository {
	return &repo{
		db: client,
	}
}

// FindOneByKey implements IMemberRepository.
func (r *repo) FindOneByKey(key string, value string) (*MemberEntity, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	coll := r.db.Database(DB_NAME).Collection(COLL_NAME)
	result := new(MemberEntity)
	if err := coll.FindOne(ctx, bson.M{key: value}).Decode(result); err != nil {
		if err.Error() != "mongo: no documents in result" {
			log.Printf("Error: FindOneByKey: %s", err.Error())
			return nil, errors.New("error: " + key + " is invalid")
		}
	}
	return result, nil
}

// InsertOne implements IRepository.
func (r *repo) InsertOne(name string, user string, pass string, email string) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	coll := r.db.Database(DB_NAME).Collection(COLL_NAME)
	memberId, err := coll.InsertOne(ctx, MemberEntity{
		FullName:  name,
		Username:  user,
		Password:  pass,
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		log.Printf("Error: InsertOneMember: %s", err.Error())
		return primitive.NilObjectID, errors.New("error: insert one member failed")
	}
	return memberId.InsertedID.(primitive.ObjectID), nil
}
