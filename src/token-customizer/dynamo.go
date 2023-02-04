package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// NewDynamoDBClient inits a DynamoDB session to be used throughout the services
func NewDynamoDBClient(isLocal bool) (*dynamodb.DynamoDB, error) {
	c := &aws.Config{
		Region: aws.String("us-west-2")}

	if isLocal {
		c.Endpoint = aws.String("http://docker.for.mac.host.internal:4566")
	}

	sess, err := session.NewSession(c)

	if err != nil {
		return nil, err
	}

	return dynamodb.New(sess), err
}
