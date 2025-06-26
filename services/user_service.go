package services

import (
	"context"
	"fmt"

	"backend-server/config"
	"backend-server/utils"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

const TableName = "Users"

// CreateUser adds a new user to the Users table
func CreateUser(userID, email, name string) error {
	input := &dynamodb.PutItemInput{
		TableName: aws.String(TableName),
		Item: map[string]types.AttributeValue{
			"UserID": &types.AttributeValueMemberS{Value: userID},
			"Email":  &types.AttributeValueMemberS{Value: email},
			"Name":   &types.AttributeValueMemberS{Value: name},
		},
		ConditionExpression: aws.String("attribute_not_exists(UserID)"),
	}
	_, err := utils.PutItemWithTimestamps(config.DynamoClient, input)
	if err != nil {
		return fmt.Errorf("CreateUser failed: %w", err)
	}
	return nil
}

// GetUser retrieves a user by UserID
func GetUser(identifier string) (map[string]types.AttributeValue, error) {
	// Try primary key (UserID) lookup
	getInput := &dynamodb.GetItemInput{
		TableName: aws.String(TableName),
		Key: map[string]types.AttributeValue{
			"UserID": &types.AttributeValueMemberS{Value: identifier},
		},
	}

	result, err := config.DynamoClient.GetItem(context.TODO(), getInput)
	if err != nil {
		return nil, fmt.Errorf("GetUser by UserID failed: %w", err)
	}
	if result.Item != nil {
		return result.Item, nil
	}

	// Fallback to query via GSI on Email
	queryInput := &dynamodb.QueryInput{
		TableName:              aws.String(TableName),
		IndexName:              aws.String("EmailIndex"),
		KeyConditionExpression: aws.String("Email = :email"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":email": &types.AttributeValueMemberS{Value: identifier},
		},
		Limit: aws.Int32(1),
	}

	queryResult, err := config.DynamoClient.Query(context.TODO(), queryInput)
	if err != nil {
		return nil, fmt.Errorf("GetUser by Email failed: %w", err)
	}
	if len(queryResult.Items) == 0 {
		return nil, fmt.Errorf("user not found by userid or email")
	}

	return queryResult.Items[0], nil
}

// UpdateUser updates the Email and Name of an existing user
func UpdateUser(userID, email, name string) error {
	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(TableName),
		Key: map[string]types.AttributeValue{
			"UserID": &types.AttributeValueMemberS{Value: userID},
		},
		UpdateExpression: aws.String("SET Email = :email, Name = :name"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":email": &types.AttributeValueMemberS{Value: email},
			":name":  &types.AttributeValueMemberS{Value: name},
		},
	}
	_, err := utils.UpdateItemWithUpdatedAt(config.DynamoClient, input)
	if err != nil {
		return fmt.Errorf("UpdateUser failed: %w", err)
	}
	return nil
}

// DeleteUser removes a user by UserID
func DeleteUser(userID string) error {
	input := &dynamodb.DeleteItemInput{
		TableName: aws.String(TableName),
		Key: map[string]types.AttributeValue{
			"UserID": &types.AttributeValueMemberS{Value: userID},
		},
	}
	_, err := config.DynamoClient.DeleteItem(context.TODO(), input)
	if err != nil {
		return fmt.Errorf("DeleteUser failed: %w", err)
	}
	return nil
}
