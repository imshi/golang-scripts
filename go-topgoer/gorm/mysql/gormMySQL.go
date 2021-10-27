// 在 GORM 中，所有的增删改查操作都可以通过 GORM 库提供的方法来实现，我们不用维护原生 SQL 语句了~~
package main

import (
	"database/sql"
	"fmt"
	"mysql/config"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 定义包内全局变量
var dbConn *gorm.DB

// 模型1：用户模型
type User struct {
	// gorm 内置了一个 gorm.Model 结构体，其包括字段 ID、CreatedAt、UpdatedAt、DeletedAt等字段，其他结构体直接匿名嵌套即可
	// gorm 默认使用ID作为主键，对应表中的id；
	// 字段标签 index：根据参数创建索引，多个字段使用相同的名称则创建复合索引
	gorm.Model
	NickName     string         `gorm:"type:varchar(20);not null;default:'';comment:昵称"`
	Age          uint8          `gorm:"size:4;comment:年龄"`
	Phone        string         `gorm:"type:char(11);index:un_phone;comment:手机号"`
	MemberNumber string         `gorm:"type:varchar(20);index:un_phone;comment:会员编号"`
	Birthday     sql.NullString `gorm:"type:varchar(10);comment:生日"`
	ActivatedAt  sql.NullTime   `gorm:"comment:激活时间"`
}

// 模型2：文章模型
type Post struct {
	gorm.Model
	Title    string
	Content  string
	Author   string `gorm:"not null"`
	Comments []Comment
}

// 模型3：文章内容模型
type Comment struct {
	gorm.Model
	Content string
	Author  string `gorm:"not null"`
	PostId  int    `gorm:"index:post_id"`
}

// init函数用于程序执行前初始化，会在程序运行时自动调用、且只会被调用一次
// 初始化 mysql 客户端，基于 database/sql 连接，进行二次封装
func init() {

	// 在输出日志中添加文件名和方法信息
	logrus.SetReportCaller(true)

	// 从配置文件中加载配置
	config.Conf()
	username := viper.GetString("database.username")
	password := viper.GetString("database.password")
	host := viper.GetString("database.host")
	port := viper.GetInt("database.port")
	dbname := viper.GetString("database.dbname")

	// 格式化MySQL连接字符串
	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbname)
	// 基于 database/sql 连接，进行二次封装
	db, err := sql.Open("mysql", dns)
	if err != nil {
		logrus.Error("[ERR001]:", err)
	}
	// 连接实例
	dbConn, err = gorm.Open(mysql.New(mysql.Config{Conn: db, DefaultStringSize: 256}))
	if err != nil {
		logrus.Error("[ERR002]:", err)
	}

	// 自动迁移schema,(根据结构体创建或者更新schema)
	// 可以注释下面这行关闭自动迁移功能
	dbConn.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{}, &Post{}, &Comment{})

}

// 单条插入
func TestAddOne() {
	// 初始化结构体
	userRow := User{
		NickName:     "John",
		Age:          18,
		Phone:        "1234000",
		MemberNumber: "A0001",
		Birthday:     sql.NullString{String: "1991-03-04", Valid: true},
		ActivatedAt:  sql.NullTime{Time: time.Now(), Valid: true},
	}
	// 通过数据指针 为 users 表插入记录入
	// Create():插入单条数据，id自增，NickName会重复
	// FirstOrCreate():获取第一个匹配的记录，或创建一个具有给定条件的新记录（仅适用于struct, map条件），此处保证NickName唯一
	// dbConn.Create(&userRow)
	dbConn.FirstOrCreate(&userRow, User{NickName: "John"})

	/*	获取详细返回信息
		result := dbConn.Create(&userRow)
		fmt.Println(result)
		fmt.Println("插入记录错误: ", result.Error)
	*/

}

// 其他增删改查、批量操作，参考：https://mp.weixin.qq.com/s/uQ3KOpRwtwNQoeAjNSne9A
// 关联模型操作，参考：https://mp.weixin.qq.com/s/f6OxX4O-e0uVEM3TJQi2yw

// 主函数
func main() {
	// 插入记录
	TestAddOne()
}
