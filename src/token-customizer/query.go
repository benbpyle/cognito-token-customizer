package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	log "github.com/sirupsen/logrus"
)

type QueryUser interface {
	GetUser(context.Context, string) (*User, error)
}

type UserService struct {
	Dao   *dynamodb.DynamoDB
	Table string
}

func (u *UserService) GetUser(ctx context.Context, username string) (*User, error) {
	log.WithFields(log.Fields{
		"key":   fmt.Sprintf("USER#%s", username),
		"table": u.Table,
	}).Info("Printing out before the query")
	key := fmt.Sprintf("USERPROFILE#%s", username)
	input := &dynamodb.GetItemInput{

		Key: map[string]*dynamodb.AttributeValue{
			"PK": {
				S: aws.String(key),
			},
			"SK": {
				S: aws.String(key),
			},
		},
		ReturnConsumedCapacity: nil,
		TableName:              aws.String(u.Table),
	}

	result, err := u.Dao.GetItemWithContext(ctx, input)

	if err != nil {
		return nil, err
	}

	user := &User{}

	_ = dynamodbattribute.UnmarshalMap(result.Item, user)

	log.WithFields(log.Fields{
		"user": user,
		"err":  err,
	}).Debug("printing out the user after being packaged from query")
	return user, err
}
