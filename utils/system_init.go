package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	PublicKey = "websocket"
)

var db *gorm.DB
var rdb *redis.Client

func InitMySql() {
	dsn := viper.Get("mysql.dns").(string)
	var err error
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Millisecond,
			LogLevel:      logger.Info,
			Colorful:      true,
		})
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		// SkipDefaultTransaction: false,
		// NamingStrategy: schema.NamingStrategy{
		// 	TablePrefix:   "",
		// 	SingularTable: false,
		// 	NoLowerCase:   false,
		// },
		Logger: newLogger,
	})
	if err != nil {
		panic("connect db error, error = " + err.Error())
	} else {
		fmt.Println("Connected to redis!")
	}
}

func InitRedis() {
	// fmt.Println("conn redis")
	addr, pwd, db, poolSize, minIdleConn :=
		viper.Get("redis.addr").(string),
		viper.Get("redis.password").(string),
		viper.Get("redis.DB").(int),
		viper.Get("redis.poolSize").(int),
		viper.Get("redis.minIdleConn").(int)
	rdb = redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     pwd, // 没有密码，默认值
		DB:           db,  // 默认DB 0
		PoolSize:     poolSize,
		MinIdleConns: minIdleConn,
	})
	ctx := context.Background()
	st := rdb.Ping(ctx)
	if st.Err() != nil {
		fmt.Println(st.Err().Error())
	}

}

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println("config app: ", viper.Get("dns"))
	fmt.Println("config mysql: ", viper.Get("mysql.dns"))
}

func GetDB() *gorm.DB {
	if db == nil {
		InitMySql()
	}
	// res := []*models.UserBasic{}
	// db.Find(res)
	// fmt.Println(res)

	return db
}

func GetRDB() *redis.Client {
	if rdb == nil {
		InitRedis()
	}
	return rdb
}

func Publish(ctx context.Context, channel string, msg string) error {
	rdb := GetRDB()
	err := rdb.Publish(ctx, channel, msg).Err()
	if err != nil {
		fmt.Println("error in subscribe", err.Error())
		return err
	}
	return err
}

func Subscribe(ctx context.Context, channel string) (string, error) {
	rdb := GetRDB()
    // fmt.Println(rdb.Ping(ctx).Val())
	sub := rdb.Subscribe(ctx, channel)
    fmt.Println("sub... ", ctx)
	// fmt.Println("Subscribe... from- ", channel, " msg- ", msg.Payload)
	msg, err := sub.ReceiveMessage(ctx)
	if err != nil {
		fmt.Println("error in subscribe", err.Error())
		return "", err
	}
	fmt.Println("Subscribe... from- ", channel, " msg- ", msg.Payload)
	return msg.Payload, err
}
