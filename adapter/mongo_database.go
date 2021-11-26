package adapter

import (
	"backend-wale/config"
	appmongo "backend-wale/database/mongo"
	"backend-wale/model"
	"context"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"math/rand"
	"time"
)

var Mongodb *mongo.Database

func NewMongoDatabaseAdapter(ctx context.Context, config config.Config) (*mongo.Database, error) {
	db, err := appmongo.NewDriver(ctx, appmongo.Config{
		URI:     config.MongoURI,
		Timeout: config.MongoTimeout,
	})
	if err != nil {
		return nil, err
	}
	Mongodb = db
	return Mongodb, nil
}

func CreateAccountNormal(ctx context.Context, email, password string) (bool, error) {
	salt := generateSalt()

	newAccount := model.CreateAccount{
		Email:        email,
		PasswordHash: computePasswordHash(password, salt),
		Salt:         salt,
		CreatedAt:    time.Now(),
	}

	filter := bson.D{{"email", email}} // pass email to filter query

	cursor, err := Mongodb.Collection("backend").Find(
		ctx, filter)

	if cursor.Decode(&newAccount) != nil {
		log.Println("user email address already taken")
		return false, errors.New("user email address already taken")
	}

	accountInsert, err := Mongodb.Collection("backend").InsertOne(ctx, newAccount)
	if err != nil {
		log.Println("the error with this: ", err)
		return false, err
	} else {
		log.Println("the insertId of the job insert is : ", accountInsert.InsertedID)
	}

	//return "Welcome " + email + "! You've been successfully logged-in" , err

	return true, nil

}

func GetAccountNormal(ctx context.Context, email, password string) (bool, error) {

	filter := bson.M{}
	filter["email"] = email

	cursor, err := Mongodb.Collection("backend").Find(
		ctx, filter)

	if cursor.ID() == 0 {
		return false, err
	}
	var result model.CreateAccount
	errr := cursor.Decode(&result)
	if errr != nil {
		return false, errr
	}

	if !(computePasswordHash(password, result.Salt) == result.PasswordHash) {
		return false, err
	}

	return true, nil

}

func generateSalt() string {
	rand.Seed(time.Now().Unix())
	result := ""

	for i := 0; i <= 8; i++ {
		result += fmt.Sprint('0' + rand.Intn(41))
	}

	return result
}

func computePasswordHash(password, salt string) string {
	hasher := sha512.New()
	hasher.Write([]byte(password + salt))
	return hex.EncodeToString(hasher.Sum(nil))
}
