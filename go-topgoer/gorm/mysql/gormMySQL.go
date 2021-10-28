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
	Title   string
	Content string
	Author  string `gorm:"not null"`
	Type    string
}

// 模型3：评论模型
type Comment struct {
	gorm.Model
	Content string
	Author  string `gorm:"not null"`
	PostId  int    `gorm:"index"`
}

// init函数用于程序执行前初始化，会在程序运行时自动调用、且只会被调用一次
// 初始化 mysql 客户端，基于 database/sql 连接，进行二次封装
func init() {

	// 在输出日志中添加文件名和方法信息
	// logrus.SetReportCaller(true)

	// 从配置文件中加载配置
	config.Conf()
	username := viper.GetString("database.username")
	password := viper.GetString("database.password")
	host := viper.GetString("database.host")
	port := viper.GetInt("database.port")
	dbname := viper.GetString("database.dbname")

	// 格式化MySQL连接字符串 - 用户表（users）
	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbname)
	// 基于 database/sql 连接，进行二次封装
	db, err := sql.Open("mysql", dns)
	if err != nil {
		panic(err)
	}
	// 连接实例
	dbConn, err = gorm.Open(mysql.New(mysql.Config{Conn: db, DefaultStringSize: 256}))
	if err != nil {
		panic(err)
	}

	// 自动迁移schema,(根据结构体创建或者更新schema)
	// 可以注释下面这行关闭自动迁移功能
	err = dbConn.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{}, &Post{}, &Comment{})
	if err != nil {
		panic(err)
	}

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
// 批量插入
func BatchInsert() {
	// 定义 User类型 结构体切片，作为插入的内容
	userRows := []User{
		{NickName: "路人甲", Age: 20, Phone: "20000000", MemberNumber: "A0002"},
		{NickName: "路人乙", Age: 22, Phone: "30000000", MemberNumber: "A0003"},
		{NickName: "路人丙", Age: 24, Phone: "40000000", MemberNumber: "A0004"},
	}
	result := dbConn.FirstOrCreate(&userRows)
	fmt.Println("插入记录错误：", result.Error)
	fmt.Println("成功插入记录条数：", result.RowsAffected)
	for _, row := range userRows {
		fmt.Println("插入ID：", row.ID)
	}
}

// 单条查询：First、Take、Last；【注意】: 如果表里没有数据，会报错result.Error返回ErrRecordNotFound，使用Find可以避免此类错误；软删除了的数据不会被获取到
/*
First:  获取第一条记录（主键升序）
Take: 获取一条记录，没有指定排序字段
Last: 获取最后一条记录（主键降序）
*/
func GetOne() {
	var firstUser User
	var takeUser User
	var lastUser User
	var result *gorm.DB
	// 获取第一条记录（主键升序），等价于: SELECT * FROM users ORDER BY id LIMIT 1;
	// 此处获取记录值；
	dbConn.First(&firstUser)
	fmt.Printf("First Result: id: %d nick_name: %s \n", firstUser.ID, firstUser.NickName)

	// 获取一条记录，没有指定排序字段，等价于: SELECT * FROM users LIMIT 1;
	// 此处获取记录值；
	dbConn.Take(&takeUser)
	fmt.Printf("Take Result: id: %d nick_name: %s \n", takeUser.ID, takeUser.NickName)

	// 获取最后一条记录（主键降序），等价于: SELECT * FROM users ORDER BY id DESC LIMIT 1;
	// 此处获取受影响条数；
	result = dbConn.Last(&lastUser)
	fmt.Printf("Last Result: id: %d nick_name: %s \n", lastUser.ID, lastUser.NickName)
	fmt.Printf("Last Result: %v\n", result.RowsAffected)
}

// 关联模型操作，参考：https://mp.weixin.qq.com/s/f6OxX4O-e0uVEM3TJQi2yw

// 多条记录查询：Find
func GetByFind() {
	var userList []User
	// 指定查询字段
	result := dbConn.Select("id", "nick_name").Find(&userList)
	for _, user := range userList {
		fmt.Printf("id: %d nick_name: %s \n", user.ID, user.NickName)
	}
	fmt.Println("查询错误：", result.Error)
}

// 条件查询
func GetByStringWhere() {
	// 定义对应结构体变量，存储查询结果
	var user User
	var userList []User
	var result *gorm.DB

	// 字符串条件查询 - 一条
	dbConn.Where("nick_name = ?", "John").First(&user)
	fmt.Println(user.NickName)

	// 字符串条件查询 - 多条
	dbConn.Where("nick_name <> ?", "John").Find(&userList)
	for _, v := range userList {
		fmt.Println(v.NickName)
	}

	// 多个条件查询 - 一条
	result = dbConn.Where("nick_name = ? and age >= ?", "John", 18).First(&user)
	fmt.Printf("ResultCount: %v err:%v \n", result.RowsAffected, result.Error)

}

// 使用 Struct 、Map 主键切片作为查询条件
// 当使用结构作为条件查询时，GORM 只会查询非零值字段。即：如果查询条件字段值为 0、''、false 或其他 零值，该字段会从查询条件中剔除、不会被用于构建查询规则
func GetByStructAndMap() {
	// 定义对应结构体变量，存储查询结果
	var user User
	var userList []User

	// 以 Struct 作为条件
	dbConn.Where(&User{NickName: "John", Age: 19}).First(&user)
	fmt.Println(user.NickName)

	// 以 map 作为条件
	dbConn.Where(map[string]interface{}{"age": 18}).Find(&userList)
	for _, v := range userList {
		fmt.Println(v.NickName)
	}

	// 以主键（ID）切片作为条件
	dbConn.Where([]int64{2, 3, 4, 5, 6}).Find(&userList)
	for _, v := range userList {
		fmt.Println(v.NickName)
	}

}

// 更新单个字段
func UpdateColumn() {
	var result *gorm.DB
	// 字符串条件更新
	// UPDATE users SET age=30, updated_at=当前时间 WHERE nick_name='John';
	result = dbConn.Model(&User{}).Where("nick_name = ?", "John").Update("age", 30)
	fmt.Printf("条件更新条数: %+v err:%v \n", result.RowsAffected, result.Error)

	// 当使用了 Model 方法，且实参对象主键有值，该值会被用于组成过滤条件 Where
	var user_m User
	dbConn.First(&user_m)
	dbConn.Model(&user_m).Update("nick_name", "赵麻子")

	// 结构体条件更新
	// UPDATE users SET age=25, updated_at=当前时间 WHERE member_number='A0002';
	result = dbConn.Model(&User{}).Where(&User{MemberNumber: "A0002"}).Update("age", 25)
	fmt.Printf("结构体条件更新条数: %+v err: %v \n", result.RowsAffected, result.Error)
}

// 更新多个字段
func MultipleColumn() {
	var result *gorm.DB
	// 使用 map 作为新数据
	updateMap := map[string]interface{}{
		"age":      32,
		"birthday": "1991-01-02",
	}
	// UPDATE users SET age=32,birthday='1991-01-02',updated_at=当前时间 WHERE id=5;
	result = dbConn.Model(&User{}).Where("id = ?", 5).Updates(updateMap)
	fmt.Printf("使用 map 更新指定记录的多个字段（受影响条数）: %+v, err: %v \n", result.RowsAffected, result.Error)

	// 使用 struct 作为新数据，不使用 Select
	updateUser := User{
		Birthday: sql.NullString{String: "1993-10-10", Valid: true},
		Age:      0,
	}
	// 【注意】：这里的 age = 0 不会更新到MySQL
	// UPDATE users SET birthday='1993-09-09',updated_at=当前时间 WHERE id=5;
	result = dbConn.Model(&User{}).Where("id=?", 5).Updates(updateUser)
	fmt.Printf("使用struct更新: %+v err:%v \n", result.RowsAffected, result.Error)
	// 使用 struct 作为新数据，通过 Select 指定更新的字段
	updateUser2 := User{
		Birthday: sql.NullString{String: "1993-09-03", Valid: true},
		Age:      0,
	}

	// 通过 Select 指定要更新的字段
	// UPDATE users SET birthday='1993-09-09',age=0,updated_at=当前时间 WHERE id=4;
	result = dbConn.Model(&User{}).Select("birthday", "age").Where("id = ?", 4).Updates(updateUser2)
	fmt.Printf("使用struct更新-v2: %+v err:%v \n", result.RowsAffected, result.Error)
}

// 软删除：不真正删除数据库中的数据，数据是可以恢复的，只是对用户来讲是不可见的
// 内置 gorm.DeletedAt 字段可用于软删除：将 DeletedAt 置为当前时间表示数据已删除
func SoftDelete() {
	var result *gorm.DB
	// 根据主键，删除一条记录
	result = dbConn.Delete(&User{}, 1)
	fmt.Printf("根据主键删除一条: %+v err:%v \n", result.RowsAffected, result.Error)

	// 根据主键切片，删除多条记录
	result = dbConn.Delete(&User{}, []int64{2, 3})
	fmt.Printf("根据主键切片删除多条: %+v err:%v \n", result.RowsAffected, result.Error)

	// 根据条件删除
	result = dbConn.Where("age = ?", 0).Delete(&User{})
	fmt.Printf("根据条件删除: %+v err:%v \n", result.RowsAffected, result.Error)
}

// 硬删除：永久删除，使用 Unscoped() 方法
func StrongDelete() {
	var result *gorm.DB
	result = dbConn.Unscoped().Delete(&User{}, 1)
	fmt.Printf("硬删除: %+v err:%v \n", result.RowsAffected, result.Error)
}

// 事务（Transaction）
// 自动事务：每一条数据操作语句都自动地成为一个事务，事务的开始是隐式的，事务的结束有明确的标记（commit 或者 rollback标志事务结束）；
func AutoTransaction() {
	err := dbConn.Transaction(func(tx *gorm.DB) error {
		//在事务中执行一些 db 操作
		user := User{NickName: "老王", Age: 48}
		if err := tx.Create(&user).Error; err != nil {
			// 回滚事务
			return err
		}
		return nil
	})

	if err != nil {
		panic(err)
	}
	logrus.Info("自动事务执行完成")
}

// 手动事务：显式定义其开始和结束的事务，当使用 start 和 commit语句时则表示发生显式事务
func ManualTransaction() {
	user := User{NickName: "小丽", Age: 19}
	// 开启事务
	tx := dbConn.Begin()
	// 新增：添加用户
	if err := tx.Create(&user).Error; err != nil {
		// 遇到错误时回滚事务
		fmt.Println("添加用户失败：", err)
		tx.Rollback()
	}

	// 提交事务
	tx.Commit()
	// fmt.Println("手动事务执行完成")
	logrus.Info("手动事务执行完成")
}

// 【关联查询】关联查询必须使用外键，不推荐在数据库中使用外键，故不推荐使用关联查询
// database中其他表的操作
func OtherCreate() {
	// 插入文章记录
	post := Post{Title: "GORM示例", Content: "基于 GORM 进行数据库增删改查", Author: "Alan"}
	result := dbConn.Create(&post)
	fmt.Println("文章表插入操作受影响记录条数：", result.RowsAffected)

	// 插入评论记录
	comment := Comment{Content: "GORM示例-评论", Author: "张三", PostId: 2}
	dbConn.Create(&comment)
}

// 主函数
func main() {
	// 插入记录：Create、FirstOrCreate
	// TestAddOne()

	// 批量插入记录：Create、FirstOrCreate
	// BatchInsert()

	// 单条查询：First、Take、Last
	GetOne()

	// 多条记录查询：Find
	// GetByFind()

	// 条件查询
	// GetByStringWhere()

	// 使用 Struct 、Map 主键切片作为查询条件
	// GetByStructAndMap()

	// 更新单个字段
	// UpdateColumn()

	// 更新多个字段
	// MultipleColumn()

	// 软删除
	// SoftDelete()

	// 硬删除
	// StrongDelete()

	// 自动执行事务
	// AutoTransaction()

	// 手动执行事务
	// ManualTransaction()

	// database中其他表的操作
	// OtherCreate()
}
