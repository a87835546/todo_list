package logic

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

const (
	dbUser     string = "root"
	dbPassword string = "12345678"
	dbHost     string = "127.0.0.1"
	dbPort     int    = 3306
	dbName     string = "task"
)

var dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&loc=Local&parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)

var Db *gorm.DB
var MongoDB *mongo.Database
var Client *redis.Client

func InitDb() *gorm.DB {
	Db = connectDB()
	return Db
}

func InitMongoDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	err = client.Ping(ctx, readpref.Primary())
	db := client.Database("todo")
	MongoDB = db
	collection := db.Collection("user")
	filter := bson.M{"username": "zhansan"}
	var result bson.M
	if err := collection.FindOne(ctx, filter).Decode(&result); err != nil {
		log.Printf("mongodb find one err-->>> %s\n", err.Error())
	}
	log.Printf("mongodb find one result -->> %s\n", result)
}
func InitRedis() error {
	Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // host:port of the redis server
		Password: "",               // no password set
		DB:       0,                // use default DB
	})

	if err := Client.Ping().Err(); err != nil {
		log.Println("err--->>>", err.Error())
	} else {
		//Client.Set("test", "a", -1)
		log.Println("redis connect success")
	}
	return nil
}
func connectDB() *gorm.DB {
	var err error
	fmt.Println("dsn : ", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})

	if err != nil {
		fmt.Printf("Error connecting to database : error=%v\n", err)
		return nil
	} else {
		log.Println("db connect success")
	}
	return db
}
