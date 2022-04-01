package repository

import (
	"context"

	"github.com/readreceipt/api/model"
	"github.com/readreceipt/api/service/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UpsertReceipt(user model.User, receipt model.Receipt) error {
	uo := &options.UpdateOptions{}
	uo.SetUpsert(true)

	filter := bson.M{"email": user.Email}
	update := bson.M{
		"$set": bson.M{"email": user.Email},
		"$push": bson.M{
			"receipts": receipt,
		},
	}

	_, err := database.
		Database().
		Collection("users").
		UpdateOne(context.TODO(), filter, update, uo)

	return err
}

func IsReceiptRead(id string) (bool, error) {
	c, err := database.
		Database().
		Collection("users").
		Aggregate(
			context.TODO(),
			mongo.Pipeline{
				{{"$unwind", "$receipts"}},
				{{
					"$match", bson.D{
						{"receipts.id", id},
						{"receipts.isread", true},
					},
				}},
				{{"$count", "receipts"}},
			},
		)

	if err != nil {
		return false, err
	}

	var results []map[string]int

	c.All(context.TODO(), &results)

	return len(results) == 1, nil
}

func UpdateSetReceiptRead(id string) error {
	filter := bson.M{"receipts.id": id}
	update := bson.M{
		"$set": bson.M{"receipts.$.isread": true},
	}

	_, err := database.
		Database().
		Collection("users").
		UpdateOne(context.TODO(), filter, update)

	return err
}
