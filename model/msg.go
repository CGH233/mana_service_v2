package model

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MessageItem ... 消息结构体
type MessageItem struct {
	OS      string `json:"os" bson:"os" binding:"required"`     // 操作系统 iOS/Android
	Page    string `json:"page" bson:"page" binding:"required"` // 页面id 如 com.muxistudio.main
	Msg     string `json:"msg" bson:"msg"`                      // 显示的消息
	Detail  string `json:"detail" bson:"detail"`                // 任意的字符串，比如在 kind 是 link 的时候，detail 就是点击消息跳转的 URL
	Time    string `json:"time" bson:"time"`                    // 时间 如 2018-04-27
	Version string `json:"version" bson:"version"`              // 版本号
	Kind    string `json:"kind" bson:"kind"`                    // 类别，如 text/link
}

// func (u *MessageItem) Create() error {
// 	collection := DB.Self.Database("ccnubox").Collection("msg")

// 	_, err := collection.InsertOne(context.TODO(), u)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// GetMsg ... 获取消息
func GetMsg(os, page string) (*MessageItem, error) {
	collection := DB.Self.Database("ccnubox").Collection("msg")
	var result MessageItem

	err := collection.FindOne(context.TODO(), bson.D{{
		"os", os,
	}, {
		"page", page,
	}}).Decode(&result)
	return &result, err
}

// Update ... 创建/更新消息
func (u *MessageItem) Update() error {
	collection := DB.Self.Database("ccnubox").Collection("msg")

	// Update a document
	filter := bson.D{
		{
			"os", u.OS,
		}, {
			"page", u.Page,
		},
	}

	update := bson.D{
		{"$set", u},
	}

	updateResult, err := collection.UpdateOne(context.TODO(), filter, update, &options.UpdateOptions{
		Upsert: newTrue(),
	})

	if err != nil {
		return err
	}

	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	return nil
}

// DeleteMsg ... 删除消息
func DeleteMsg(os, page string) error {
	collection := DB.Self.Database("ccnubox").Collection("msg")

	deleteResult, err := collection.DeleteOne(context.TODO(), bson.D{
		{
			"os", os,
		}, {
			"page", page,
		},
	})

	if err != nil {
		return err
	}

	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
	return nil
}
