package model

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// IOSConfig ... iOS 元信息
type IOSConfig struct {
	CalendarURL              string          `json:"calendarUrl" bson:"calendarUrl"`                           // 校历
	FlashScreenURL           string          `json:"flashScreenUrl" bson:"flashScreenUrl"`                     // 闪屏
	ShowGuisheng             string          `json:"showGuisheng" bson:"showGuisheng"`                         // 历史兼容
	StartCountDayPreset      string          `json:"startCountDayPreset" bson:"startCountDayPreset"`           // 历史兼容
	StartCountDayPresetForV2 string          `json:"startCountDayPresetForV2" bson:"startCountDayPresetForV2"` // 学期开始日
	UpdateInfo               string          `json:"updateInfo" bson:"updateInfo"`                             // 更新说明
	Version                  string          `json:"version" bson:"version"`                                   // 当前最新版本
	ShouldPullCourse         bool            `json:"shouldPullCourse" bson:"shouldPullCourse"`                 // 自动更新课程开关
	FlashStartDay            string          `json:"flashStartDay" bson:"flashStartDay"`                       // 闪屏显示开始日期
	FlashEndDay              string          `json:"flashEndDay" bson:"flashEndDay"`                           // 闪屏显示结束日期
	GradeJSUrl               string          `json:"gradeJSUrl" bson:"gradeJSUrl"`                             // 历史兼容
	TableJSUrl               string          `json:"tableJSUrl" bson:"tableJSUrl"`                             // 历史兼容
	Rax                      []RaxConfigItem `json:"rax" bson:"rax"`                                           // Rax 热更新
}

// RaxConfigItem ... Rax 热更新配置项
type RaxConfigItem struct {
	Key     string `json:"key" bson:"key"`         // Rax 包名
	Version string `json:"version" bson:"version"` //  版本号
	URL     string `json:"url" bson:"url"`         // URL
}

func newTrue() *bool {
	b := true
	return &b
}

// Create creates a new user account.
func (u *IOSConfig) Update() error {
	collection := DB.Self.Database("ccnubox").Collection("config")

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

func GetConfig() (*IOSConfig, error) {
	collection := DB.Self.Database("ccnubox").Collection("config")
	var result IOSConfig

	err := collection.FindOne(context.TODO(), bson.D{{}}).Decode(&result)
	return &result, err
}
