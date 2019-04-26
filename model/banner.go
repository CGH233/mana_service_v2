package model

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// BannerItem ... iOS 元信息
type BannerItem struct {
	Img string `json:"img" bson:"img"` // 图片 URL
	URL string `json:"url" bson:"url"` // URL
	Num int    `json:"num" bson:"num"` // 权重（排序用）
}

// BannerConfig ... iOS 元信息
type BannerConfig struct {
	List []BannerItem `json:"list" bson:"list"`
}

func (u *BannerConfig) Update() error {
	collection := DB.Self.Database("ccnubox").Collection("banner")

	// Update a document
	filter := bson.D{}

	// val, err := bson.Marshal(&u)
	// if err != nil {
	// 	return err
	// }

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

func GetBanner() (*BannerConfig, error) {
	collection := DB.Self.Database("ccnubox").Collection("banner")
	var result BannerConfig

	err := collection.FindOne(context.TODO(), bson.D{{}}).Decode(&result)
	return &result, err
}
