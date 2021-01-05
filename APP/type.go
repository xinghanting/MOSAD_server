package app

import "go.mongodb.org/mongo-driver/bson/primitive"

//Daily 日常时间结构体
type Daily struct {
	//该日常事件的ID
	Index primitive.ObjectID `json:"index" bson:"index"`

	// 用户的ID
	OwnID string `json:"owanID" bson:"ownID"`

	// 日常事件的始时间 0-23
	Begin int `json:"begin" bson:"bgin"`

	// 日常事件的结束0-23
	End int `json:"end" bson:"end"`

	// 日常事件的描述
	Content string `json:"content" son:"content"`
}

//User 用户信息结构体
type User struct {
	ID string `json:"ID" bson:"ID"`

	Name string `json:"name" bson:"name"`

	Password string `json:"password" bson:"password"`
}

// Response 回复的数据结构
type Response struct {

	// true for success and false for failure
	State bool `json:"state"`

	// the return message of the interfaces
	Response interface{} `json:"response"`
}

//Weekly 周常事件结构体
type Weekly struct {
	//该周常事件的ID
	Index primitive.ObjectID `json:"index" bson:"index"`

	// 用户的ID
	OwnID string `json:"ownID" bson:"ownID"`

	//周常时间的星期数 (1-7)
	Week int `json:"week" bson:"week"`

	// 周常事件的开始时间 0-23
	Begin int `json:"begin" bson:"begin"`

	// 周常事件的结束时间 0-23
	End int `json:"end" bson:"end"`

	// 周常事件的描述
	Content string `json:"content" bson:"content"`
}

//Onetime 一次性事件结构体
type Onetime struct {
	//该一次性事件的ID
	Index primitive.ObjectID `json:"index" bson:"index"`

	// 用户的ID
	OwnID string `json:"ownID" bson:"ownID"`

	// DDL Of the Onetime
	DDL int64 `json:"commentTime" bson:"commentTime"`

	// Description of the Onetime
	Content string `json:"content" bson:"content"`
}
