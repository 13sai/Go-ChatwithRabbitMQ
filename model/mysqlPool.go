package model

import (
    "fmt"
	"sync"
	"errors"
	"os"
	
    orm "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

type MySqlPool struct {}

var instance *MySqlPool
var once sync.Once

var DB *orm.DB
var err error 

// 单例模式
func GetInstance() *MySqlPool {
	once.Do(func() {
		instance = &MySqlPool{}
	})

	return instance
}

func (pool *MySqlPool) InitPool() (isSuc bool) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=true", 
		viper.GetString("db.username"), 
		viper.GetString("db.password"), 
		viper.GetString("db.host"), 
		viper.GetString("db.name"), 
		viper.GetString("db.charset"))
    DB, err = orm.Open("mysql", dsn)
    if err != nil {
		fmt.Println(errors.New("mysql连接失败"))
		os.Exit(1)
        return false
	}

	// 连接数配置也可以写入配置，在此读取
	DB.DB().SetMaxIdleConns(5)
    DB.DB().SetMaxOpenConns(10)
	if viper.GetInt("db.log") == 1 {
		DB.LogMode(true)
	}
    return true
}