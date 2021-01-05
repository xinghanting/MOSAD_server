package app

import (
	"context"
	"errors"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DailyService 日常事件服务
type DailyService struct {
	DB *mongo.Collection
}

// OnetimeService 一次性事件服务
type OnetimeService struct {
	DB *mongo.Collection
}

//UserService 针对用户的服务
type UserService struct {
	DB *mongo.Collection
}

// WeeklyService 周常事件服务
type WeeklyService struct {
	DB *mongo.Collection
}

// NewService 数据库所涉及的服务
type NewService struct {
	DB      *mongo.Database
	User    *UserService
	Daily   *DailyService
	Onetime *OnetimeService
	Weekly  *WeeklyService
}

//InitNewService 初始化一个新的服务
func (service *NewService) InitNewService() (err error) {
	//首先需要连接上数据库
	clientOptions := options.Client().ApplyURI("mongodb+srv://first_user:xht123456@mosad.j8azs.mongodb.net/MOSAD_test?retryWrites=true&w=majority") // opts
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// 检查是否成功连接数据库
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connected to MongoDB Successfully!")
	}
	//开始初始化各个服务
	service.DB = client.Database("MOSAD_test")
	service.User.DB = service.DB.Collection("User")
	service.Daily.DB = service.DB.Collection("Daily")
	service.Weekly.DB = service.DB.Collection("Weekly")
	service.Onetime.DB = service.DB.Collection("Onetime")
	fmt.Println("Connect to MongoBD sucessfully")
	return
}

// CreateDaily 创建一个日常事件
func (service *DailyService) CreateDaily(ownID string, begin int, end int, content string) (dayitem Daily) {
	//进行数据合理性判断
	//(code)
	dayitem.Index = primitive.NewObjectID()
	dayitem.Begin = begin
	dayitem.End = end
	dayitem.Content = content
	dayitem.OwnID = ownID
	return
}

//PostDaily 上传日常事件
func (service *DailyService) PostDaily(dayitem Daily) (err error) {
	if dayitem.OwnID == "" {
		return errors.New("该事件没有附加用户ID")
	}
	_, err = service.DB.InsertOne(context.Background(), &dayitem)
	return err
}

// GetDailyByUserID 根据用户的ID获取所有的日常事件
func (service *DailyService) GetDailyByUserID(id string) (Dailys []Daily, err error) {
	cursor, err := service.DB.Find(context.Background(), bson.D{{"ownID", id}})
	if err != nil {
		return
	}
	err = cursor.All(context.Background(), &Dailys)
	return
}

// DeleteDaily 删除某一个特定的日常事件
func (service *DailyService) DeleteDaily(id string) (err error) {
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}
	_, err = service.DB.DeleteOne(context.Background(), bson.D{{"index", ID}})
	return
}

// CreateOnetime 创建一个一次性事件
func (service *OnetimeService) CreateOnetime(ownID string, ddl int64, content string) (onetimeitem Onetime) {
	//进行数据合理性判断
	//(code)
	onetimeitem.Index = primitive.NewObjectID()
	onetimeitem.Content = content
	onetimeitem.DDL = ddl
	onetimeitem.OwnID = ownID
	return
}

//PostOnetime 上传一次性事件
func (service *OnetimeService) PostOnetime(onetimeitem Onetime) (err error) {
	if onetimeitem.OwnID == "" {
		return errors.New("该事件没有附加用户ID")
	}
	_, err = service.DB.InsertOne(context.Background(), &onetimeitem)
	return err
}

// GetOnetimeByUserID 根据用户的ID获取所有的一次性事件
func (service *OnetimeService) GetOnetimeByUserID(id string) (Onetimes []Onetime, err error) {
	cursor, err := service.DB.Find(context.Background(), bson.D{{"ownID", id}})
	if err != nil {
		return
	}
	err = cursor.All(context.Background(), &Onetimes)
	return
}

// DeleteOnetime 删除一个一次性事件
func (service *OnetimeService) DeleteOnetime(id string) (err error) {
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}
	_, err = service.DB.DeleteOne(context.Background(), bson.D{{"index", ID}})
	return
}

// CreateWeekly 创建一个周常事件
func (service *WeeklyService) CreateWeekly(ownID string, week int, begin int, end int, content string) (weekitem Weekly, err error) {
	//进行数据合理性判断
	//(code)
	weekitem.Index = primitive.NewObjectID()
	weekitem.Begin = begin
	weekitem.End = end
	weekitem.Week = week
	weekitem.OwnID = ownID
	weekitem.Content = content
	return
}

//PostWeekly 上传周常事件
func (service *WeeklyService) PostWeekly(weekitem Weekly) (err error) {
	if weekitem.OwnID == "" {
		return errors.New("该事件没有附加用户ID")
	}
	_, err = service.DB.InsertOne(context.Background(), &weekitem)
	return err
}

// GetWeeklyByUserID 根据用户的ID获取所有的周常事件
func (service *WeeklyService) GetWeeklyByUserID(id string) (weeklys []Weekly, err error) {
	cursor, err := service.DB.Find(context.Background(), bson.D{{"ownID", id}})
	if err != nil {
		return
	}
	// err = cursor.Decode(&weeklys)
	err = cursor.All(context.TODO(), &weeklys)
	return
}

// DeleteWeekly 删除某一个特定的周常事件
func (service *WeeklyService) DeleteWeekly(id string) (err error) {
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}
	_, err = service.DB.DeleteOne(context.Background(), bson.D{{"index", ID}})
	return
}

//CreateUser 创建用户
func (service *UserService) CreateUser(id string, name string, password string) (ID string, err error) {
	user := User{}
	user.ID = id
	user.Name = name
	user.Password = password
	err = service.DB.FindOne(context.Background(), bson.D{{"ID", id}}).Err()
	if err != mongo.ErrNoDocuments {
		return user.ID, errors.New("The UserID already exists.")
	}
	_, err = service.DB.InsertOne(context.Background(), &user)
	return user.ID, err
}

// GetUserByID 通过ID获取用户信息
func (service *UserService) GetUserByID(id string) (user User, err error) {
	err = service.DB.FindOne(context.Background(), bson.D{{"ID", id}}).Decode(&user)
	return
}

// LoginUser 用户登录函数
func (service *UserService) LoginUser(id string, password string) (name string, err error) {
	var user User
	err = service.DB.FindOne(context.Background(), bson.D{{"ID", id}}).Decode(&user)
	if err != nil {
		err = errors.New("No such userID.")
	} else if user.Password != password {
		err = errors.New("The password is incorrect.")
	}
	name = user.Name
	return
}
