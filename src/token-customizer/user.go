package main

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"strconv"
)

type UserRole struct {
	TenantId   int    `dynamodbav:"TenantId"`
	EntityType string `dynamodbav:"EntityType"`
	Roles      []Role `dynamodbav:"Roles"`
}

type Role struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type User struct {
	UserId         int64  `dynamodbav:"UserId"`
	EntityType     string `dynamodbav:"EntityType"`
	EmailAddress   string `dynamodbav:"EmailAddress"`
	FirstName      string `dynamodbav:"FirstName"`
	LastName       string `dynamodbav:"LastName"`
	AllowedTenants []int  `dynamodbav:"AllowedTenants"`
	CurrentTenant  int    `dynamodbav:"CurrentTenant"`
	UserRoles      []UserRole
}

func (u *User) mapToMap() map[string]string {
	m := make(map[string]string)

	m["firstName"] = u.FirstName
	m["lastName"] = u.LastName

	return m
}

func (r *Role) UnmarshalDynamoDBAttributeValue(value *dynamodb.AttributeValue) error {
	for k, kv := range value.M {
		if k == "id" {
			v, _ := strconv.ParseInt(*kv.N, 10, 32)
			r.Id = int(v)
		} else if k == "name" {
			r.Name = *kv.S
		}
	}

	return nil
}

func (ur *UserRole) UnmarshalDynamoDBAttributeValue(value *dynamodb.AttributeValue) error {
	for k, kv := range value.M {
		if k == "EntityType" {
			ur.EntityType = *kv.S
		} else if k == "TenantId" {
			v, _ := strconv.ParseInt(*kv.N, 10, 32)
			ur.TenantId = int(v)
		} else if k == "Roles" {
			for _, nkv := range kv.M {
				r := &Role{}
				err := dynamodbattribute.UnmarshalMap(nkv.M, r)
				if err != nil {
					return err
				}
				ur.Roles = append(ur.Roles, *r)
			}
		}
	}
	return nil
}
