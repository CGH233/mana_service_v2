package model

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// FeedbackItem ... 用户反馈
type FeedbackItem struct {
	Contact   string `json:"contact" bson:"contact"`                    // 联系方式
	Content   string `json:"content" bson:"content" binding:"required"` // 反馈内容
	Timestamp int64  `json:"timestamp" bson:"timestamp"`                // 时间戳
}

// Create ... 创建用户反馈
func (u *FeedbackItem) Create() error {
	collection := DB.Self.Database("ccnubox").Collection("feedback")

	u.Timestamp = time.Now().Unix()
	_, err := collection.InsertOne(context.TODO(), u)

	if err != nil {
		return err
	}

	return nil
}

// ListFeedback ... 获取用户反馈列表（分页）
func ListFeedback(offset, limit int) ([]*FeedbackItem, int64, error) {
	if limit == 0 {
		limit = LIST_DEFAULT_LIMIT
	}
	collection := DB.Self.Database("ccnubox").Collection("feedback")

	findOptions := options.Find()
	findOptions.SetLimit(int64(limit))
	findOptions.SetSkip(int64(offset))

	var results = make([]*FeedbackItem, 0)

	count, err := collection.CountDocuments(context.TODO(), bson.D{{}})
	if err != nil {
		return results, count, err
	}

	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		return results, count, err
	}

	for cur.Next(context.TODO()) {

		var elem FeedbackItem
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		return results, count, err
	}

	cur.Close(context.TODO())

	return results, count, nil
}
