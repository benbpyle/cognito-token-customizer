package main

import (
	"context"
	"os"
	"strconv"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	log "github.com/sirupsen/logrus"
)

var (
	svc QueryUser
)

func handler(ctx context.Context, e events.CognitoEventUserPoolsPreTokenGen) (events.CognitoEventUserPoolsPreTokenGen, error) {
	log.WithFields(log.Fields{
		"event": e,
	}).Debug("logging out the debug event")

	u, err := svc.GetUser(ctx, e.UserName)
	cod := events.ClaimsOverrideDetails{}
	if err == nil && u != nil {
		cod.ClaimsToAddOrOverride = u.mapToMap()
	} else if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("Error querying dynamodb")
	} else {
		log.Info("No error and nothing found")
	}

	resp := events.CognitoEventUserPoolsPreTokenGenResponse{
		ClaimsOverrideDetails: cod,
	}

	e.Response = resp
	return e, nil
}

func setLogLevel() {
	switch level := strings.ToLower(os.Getenv("LOG_LEVEL")); level {
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "debug":
		log.SetLevel(log.DebugLevel)
	default:
		log.SetLevel(log.DebugLevel)
	}
}

func main() {
	setLogLevel()
	isLocalEnv := os.Getenv("IS_LOCAL")
	table := os.Getenv("TABLE_NAME")
	localParse, _ := strconv.ParseBool(isLocalEnv)

	log.SetFormatter(&log.JSONFormatter{
		PrettyPrint: false,
	})

	db, e := NewDynamoDBClient(localParse)
	if e != nil {
		log.WithFields(log.Fields{"error": e}).Fatalf("can't start the lambda")
	}

	svc = &UserService{
		Dao:   db,
		Table: table,
	}

	lambda.Start(handler)
}
