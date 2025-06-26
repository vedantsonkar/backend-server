package utils

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func PutItemWithTimestamps(client *dynamodb.Client, input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	now := time.Now().Format(time.RFC3339)

	if _, ok := input.Item["CreatedAt"]; !ok {
		input.Item["CreatedAt"] = &types.AttributeValueMemberS{Value: now}
	}
	input.Item["UpdatedAt"] = &types.AttributeValueMemberS{Value: now}

	return client.PutItem(context.TODO(), input)
}

// Adds UpdatedAt to UpdateItem expressions
func UpdateItemWithUpdatedAt(client *dynamodb.Client, input *dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error) {
	now := time.Now().Format(time.RFC3339)

	if input.ExpressionAttributeValues == nil {
		input.ExpressionAttributeValues = make(map[string]types.AttributeValue)
	}
	input.ExpressionAttributeValues[":updatedAt"] = &types.AttributeValueMemberS{Value: now}

	// Automatically append to the UpdateExpression
	if input.UpdateExpression != nil {
		*input.UpdateExpression += ", UpdatedAt = :updatedAt"
	} else {
		expr := "SET UpdatedAt = :updatedAt"
		input.UpdateExpression = &expr
	}

	return client.UpdateItem(context.TODO(), input)
}