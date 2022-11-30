/*
【功能说明】：使用 gin + gjson + MySQL 使现 Http 接口和 Json 解析
【 InsertCustomer 接口说明】：
1. 用于记录和执行对数据库的客户表customer的Insert操作，记录包括操作的用户、时间和具体的customer信息；
2. 接口定义:InsertCustomer(customer, timestamp, params, token)
3. 请求方式：接口以http形式提供服务，接收来自客户端的post请求。
4. 接口参数（参数传递无顺序要求，请求时parmas中参数非空的必须提供, 有默认值的参数如不提供将按默认值处理）：
	1.customer: ‘用户名’ 参数类型: string
	2.timestamp: ‘时间戳’ 参数类型: int
	3.params: ‘具体的修改参数’ 参数类型:json
	4.token:’由加密串计算出的md5码’ 参数类型: string
5. params说明:
	Id: 客户ID, 非空, 参数类型: int
	Name: 客户名称, 非空, 参数类型: string
6. 返回结果：以json形式返回，包含执行是否成功的标志。
7. 请求示例（插入了一条客户Id为9999，名称为”vdncloud”的客户记录）:
curl -i -X POST –data ‘{“Customer”:”robert”,”Timestamp”:1515143723,”Params”:{“Id”:9999,”Name”:”vdncloud”},”Token”:”f4269bb22a371f5a5d7039f9622acb57”}’ http://192.168.99.100:8000/insertCustomer
8. 需要注意的地方是Json解析时，首字母必须为大写，否则会解析失败。
*/

package main

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tidwall/gjson"
)

var gMysqlIp string = "192.168.99.100"
var gSqlDb *sql.DB

type CustomerInfo struct {
	Id   int64  `json:"Id"`   //1
	Name string `json:"Name"` //2
	Del  int64  `json:"Del"`  //3
}

type DomainInfo struct {
	Id             int64  `json:"d_id"`               //1
	Name           string `json:"d_dname"`            //2
	CustId         int64  `json:"d_cust_id"`          //3
	SrcIp          string `json:"d_src_ip"`           //4
	LogFmt         int64  `json:"d_logfmt"`           //5
	LogInterval    int64  `json:"d_log_interval"`     //6
	LogWild        int64  `json:"d_log_wild"`         //7
	Type           int64  `json:"d_type"`             //8
	HType          int64  `json:"d_htype"`            //9
	LogLevel       int64  `json:"d_log_level"`        //10
	BitRate        int64  `json:"d_bit_rate"`         //11
	CostWithParent int64  `json:"d_cost_with_parent"` //12
	Del            int64  `json:"d_status"`           //13
}

type CustomerJson struct {
	Customer  string       `json:"Customer"`
	Timestamp int64        `json:"Timestamp"`
	Params    CustomerInfo `json:"Params"`
	Token     string       `json:"Token"`
}

type DomainJson struct {
	Customer  string     `json:"Customer"`
	Timestamp int64      `json:"Timestamp"`
	Params    DomainInfo `json:"Params"`
	Token     string     `json:"Token"`
}

// 是否存在该客户
func ExistsCustomer(id int64) (bool, error) {
	sql := fmt.Sprintf("SELECT c_id FROM server_conf.customer WHERE c_id=%v;", id)
	fmt.Println(sql)
	rows, err := gSqlDb.Query(sql)
	if err != nil {
		return false, err
	}

	if !rows.Next() {
		return false, nil
	}

	return true, nil
}

// 是否存在该域名
func ExistsDomain(id int64) (bool, error) {
	sql := fmt.Sprintf("SELECT d_id FROM server_conf.domain WHERE d_id=%v;", id)
	fmt.Println(sql)
	rows, err := gSqlDb.Query(sql)
	if err != nil {
		return false, err
	}

	if !rows.Next() {
		return false, nil
	}

	return true, nil
}

// 获取客户信息
func GetLastCustomer(id int64, cm *CustomerInfo) error {
	sql := fmt.Sprintf("SELECT c_id, c_name, c_if_del FROM server_conf.customer WHERE c_id=%v;", id)
	fmt.Println(sql)
	rows, err := gSqlDb.Query(sql)
	if err != nil {
		return err
	}

	for rows.Next() {
		//rows.Columns()
		err = rows.Scan(&cm.Id, &cm.Name, &cm.Del)
		if err != nil {
			fmt.Printf("Get customer info failed! [%s] \n", err)
			return err
		}

		fmt.Println(cm.Id, cm.Name, cm.Del)
	}

	return nil
}

// 获取域名信息
func GetLastDomain(id int64, dm *DomainInfo) error {
	sql := fmt.Sprintf("SELECT d_id, d_dname, d_cust_id, d_src_ip, d_logfmt, d_log_interval, d_log_wild, d_type, d_htype, d_log_level, d_bit_rate, d_cost_with_parent, d_status FROM server_conf.domain WHERE d_id=%v;", id)
	fmt.Println(sql)
	rows, err := gSqlDb.Query(sql)
	if err != nil {
		return err
	}

	for rows.Next() {
		//rows.Columns()
		err = rows.Scan(&dm.Id, &dm.Name, &dm.CustId, &dm.SrcIp, &dm.LogFmt, &dm.LogInterval, &dm.LogWild, &dm.Type,
			&dm.HType, &dm.LogLevel, &dm.BitRate, &dm.CostWithParent, &dm.Del)
		if err != nil {
			fmt.Printf("Get domain info failed! [%s] \n", err)
			return err
		}

		fmt.Println(dm.Id, dm.Name, dm.CustId, dm.SrcIp, dm.LogFmt, dm.LogInterval, dm.LogWild, dm.Type,
			dm.HType, dm.LogLevel, dm.BitRate, dm.CostWithParent, dm.Del)
	}

	return nil
}

// 插入客户信息
func InsertCustomer(cm *CustomerInfo) error {
	sql := "INSERT INTO server_conf.customer(c_id, c_name) VALUES("

	var values string = ""
	// ID
	if cm.Id != 0 {
		values += fmt.Sprintf("'%v',", cm.Id)
	}

	// Name
	if cm.Name != "None" {
		values += fmt.Sprintf("'%v',", cm.Name)
	} else {
		return fmt.Errorf("Error:Customer Name cannot be empty ")
	}

	values = strings.TrimRight(values, ",")
	if values == "" {
		return fmt.Errorf("Error:Customer info cannot be empty ")
	}

	sql += values + ");"
	fmt.Println(sql)

	_, err := gSqlDb.Query(sql)
	if err != nil {
		return err
	}

	return nil
}

// 插入域名信息
func InsertDomain(dm *DomainInfo) error {
	sql := "INSERT server_conf.domain(d_id, d_dname, d_cust_id, d_src_ip, d_logfmt, d_log_interval, d_log_wild, d_type, d_htype, d_log_level, d_bit_rate, d_cost_with_parent, d_status) VALUES("

	var values string = ""
	// ID
	if dm.Id != 0 {
		values += fmt.Sprintf("'%v',", dm.Id)
	}

	// Name
	if dm.Name != "None" {
		values += fmt.Sprintf("'%v',", dm.Name)
	} else {
		return fmt.Errorf("Error:Domain Name cannot be empty ")
	}

	// CustId
	if dm.CustId != -1 {
		values += fmt.Sprintf("'%v',", dm.CustId)
	} else {
		return fmt.Errorf("Error:Domain CustId cannot be empty ")
	}

	// SrcIp
	if dm.SrcIp != "None" {
		values += fmt.Sprintf("'%v',", dm.SrcIp)
	} else {
		values += fmt.Sprintf("'%v',", "")
	}

	// LogFmt
	if dm.LogFmt != -1 {
		values += fmt.Sprintf("'%v',", dm.LogFmt)
	} else {
		values += fmt.Sprintf("'%v',", 0)
	}

	// LogInterval
	if dm.LogInterval != -1 {
		values += fmt.Sprintf("'%v',", dm.LogInterval)
	} else {
		values += fmt.Sprintf("'%v',", 0)
	}

	// LogWild
	if dm.LogWild != -1 {
		values += fmt.Sprintf("'%v',", dm.LogWild)
	} else {
		values += fmt.Sprintf("'%v',", 0)
	}

	// Type
	if dm.Type != -1 {
		values += fmt.Sprintf("'%v',", dm.Type)
	} else {
		values += fmt.Sprintf("'%v',", 0)
	}

	// HType
	if dm.HType != -1 {
		values += fmt.Sprintf("'%v',", dm.HType)
	} else {
		values += fmt.Sprintf("'%v',", 0)
	}

	// LogLevel
	if dm.LogLevel != -1 {
		values += fmt.Sprintf("'%v',", dm.LogLevel)
	} else {
		values += fmt.Sprintf("'%v',", 0)
	}

	// BitRate
	if dm.BitRate != -1 {
		values += fmt.Sprintf("'%v',", dm.BitRate)
	} else {
		values += fmt.Sprintf("'%v',", 0)
	}

	// CostWithParent
	if dm.CostWithParent != -1 {
		values += fmt.Sprintf("'%v',", dm.CostWithParent)
	} else {
		values += fmt.Sprintf("'%v',", 0)
	}

	// Del
	if dm.Del != -1 {
		values += fmt.Sprintf("'%v',", dm.Del)
	} else {
		values += fmt.Sprintf("'%v',", 0)
	}

	values = strings.TrimRight(values, ",")
	if values == "" {
		return fmt.Errorf("Error:Domain info cannot be empty ")
	}
	fmt.Println("values=", values)

	sql += values + ");"
	fmt.Println(sql)

	_, err := gSqlDb.Query(sql)
	if err != nil {
		return err
	}

	return nil
}

// 更新客户信息
func UpdateCustomer(cm *CustomerInfo) error {
	sql := "UPDATE server_conf.customer SET "

	// Id
	if cm.Id == 0 {
		return fmt.Errorf("Error:Customer id can not be empty ")
	}

	last := &CustomerInfo{}
	GetLastCustomer(cm.Id, last)

	var values string = ""
	// Name
	if cm.Name != "None" && cm.Name != last.Name {
		values += fmt.Sprintf("c_name='%v',", cm.Name)
	}

	// Del
	if cm.Del != -1 && cm.Del != last.Del {
		values += fmt.Sprintf("c_if_del='%v',", cm.Del)
	}

	values = strings.TrimRight(values, ",")
	if values == "" {
		return fmt.Errorf("Error:Same as last Customer, nothing to update ")
	}
	fmt.Println("values=", values)

	sql += values + fmt.Sprintf(" WHERE c_id=%v;", cm.Id)
	fmt.Println(sql)

	_, err := gSqlDb.Query(sql)
	if err != nil {
		return err
	}

	return nil
}

// 更新域名信息
func UpdateDomain(dm *DomainInfo) error {
	sql := "UPDATE server_conf.domain SET "

	// Id
	if dm.Id == 0 {
		return fmt.Errorf("Error:Domain id can not be empty ")
	}

	last := &DomainInfo{}
	err := GetLastDomain(dm.Id, last)
	if err != nil {
		return fmt.Errorf("Error:Get last domain info failed ")
	}

	var values string = ""
	// Name
	if dm.Name != "None" && dm.Name != last.Name {
		values += fmt.Sprintf("d_dname='%v',", dm.Name)
	}

	// CustId
	if dm.CustId != -1 && dm.CustId != last.CustId {
		values += fmt.Sprintf("d_cust_id='%v',", dm.CustId)
	}

	// SrcIp
	if dm.SrcIp != "None" && dm.SrcIp != last.SrcIp {
		values += fmt.Sprintf("d_src_ip='%v',", dm.SrcIp)
	}

	// LogFmt
	if dm.LogFmt != -1 && dm.LogFmt != last.LogFmt {
		values += fmt.Sprintf("d_logfmt='%v',", dm.LogFmt)
	}

	// LogInterval
	if dm.LogInterval != -1 && dm.LogInterval != last.LogInterval {
		values += fmt.Sprintf("d_log_interval='%v',", dm.LogInterval)
	}

	// LogWild
	if dm.LogWild != -1 && dm.LogWild != last.LogWild {
		values += fmt.Sprintf("d_log_wild='%v',", dm.LogWild)
	}

	// Type
	if dm.Type != -1 && dm.Type != last.Type {
		values += fmt.Sprintf("d_type='%v',", dm.Type)
	}

	// HType
	if dm.HType != -1 && dm.HType != last.HType {
		values += fmt.Sprintf("d_htype='%v',", dm.HType)
	}

	// LogLevel
	if dm.LogLevel != -1 && dm.LogLevel != last.LogLevel {
		values += fmt.Sprintf("d_log_level='%v',", dm.LogLevel)
	}

	// BitRate
	if dm.BitRate != -1 && dm.BitRate != last.BitRate {
		values += fmt.Sprintf("d_bit_rate='%v',", dm.BitRate)
	}

	// CostWithParent
	if dm.CostWithParent != -1 && dm.CostWithParent != last.CostWithParent {
		values += fmt.Sprintf("d_cost_with_parent='%v',", dm.CostWithParent)
	}

	// Del
	if dm.Del != -1 && dm.Del != last.Del {
		values += fmt.Sprintf("d_status='%v',", dm.Del)
	}

	values = strings.TrimRight(values, ",")
	if values == "" {
		return fmt.Errorf("Error:Same as current domain, nothing to update ")
	}
	fmt.Println("values=", values)

	sql += values + fmt.Sprintf(" WHERE d_id=%v;", dm.Id)
	fmt.Println(sql)

	_, err = gSqlDb.Query(sql)
	if err != nil {
		return err
	}

	return nil
}

func UnmarCustomer(param string, cm *CustomerInfo) error {
	// 1
	value := gjson.Get(param, "Id")
	if value.Index != 0 {
		cm.Id = value.Int()
	} else {
		cm.Id = -1
		return fmt.Errorf("Error:Customer id can not be empty ")
	}

	// 2
	value = gjson.Get(param, "Name")
	if value.Index != 0 {
		cm.Name = value.String()
	} else {
		cm.Name = "None"
	}

	// 3
	value = gjson.Get(param, "Del")
	if value.Index != 0 {
		cm.Del = value.Int()
	} else {
		cm.Del = -1
	}

	return nil
}

func UnmarDomain(param string, dm *DomainInfo) error {
	// 1
	value := gjson.Get(param, "Id")
	if value.Index != 0 {
		dm.Id = value.Int()
	} else {
		dm.Id = -1
		return fmt.Errorf("Error:Domain id can not be empty ")
	}

	// 2
	value = gjson.Get(param, "Name")
	if value.Index != 0 {
		dm.Name = value.String()
	} else {
		dm.Name = "None"
	}

	// 3
	value = gjson.Get(param, "CustId")
	if value.Index != 0 {
		dm.CustId = value.Int()
	} else {
		dm.CustId = -1
	}

	// 4
	value = gjson.Get(param, "SrcIp")
	if value.Index != 0 {
		dm.SrcIp = value.String()
	} else {
		dm.SrcIp = "None"
	}

	// 5
	value = gjson.Get(param, "LogFmt")
	if value.Index != 0 {
		dm.LogFmt = value.Int()
	} else {
		dm.LogFmt = -1
	}

	// 6
	value = gjson.Get(param, "LogInterval")
	if value.Index != 0 {
		dm.LogInterval = value.Int()
	} else {
		dm.LogInterval = -1
	}

	// 7
	value = gjson.Get(param, "LogWild")
	if value.Index != 0 {
		dm.LogWild = value.Int()
	} else {
		dm.LogWild = -1
	}

	// 8
	value = gjson.Get(param, "Type")
	if value.Index != 0 {
		dm.Type = value.Int()
	} else {
		dm.Type = -1
	}

	// 9
	value = gjson.Get(param, "HType")
	if value.Index != 0 {
		dm.HType = value.Int()
	} else {
		dm.HType = -1
	}

	// 10
	value = gjson.Get(param, "LogLevel")
	if value.Index != 0 {
		dm.LogLevel = value.Int()
	} else {
		dm.LogLevel = -1
	}

	// 11
	value = gjson.Get(param, "BitRate")
	if value.Index != 0 {
		dm.BitRate = value.Int()
	} else {
		dm.BitRate = -1
	}

	// 12
	value = gjson.Get(param, "CostWithParent")
	if value.Index != 0 {
		dm.CostWithParent = value.Int()
	} else {
		dm.CostWithParent = -1
	}

	// 13
	value = gjson.Get(param, "Del")
	if value.Index != 0 {
		dm.Del = value.Int()
	} else {
		dm.Del = -1
	}

	return nil
}

func ResponseSuccess(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": gin.H{
			"code":    http.StatusOK,
			"message": "ok",
		},
	})
}

func ResponseError(c *gin.Context, err error) {
	fmt.Println(err)
	c.JSON(http.StatusOK, gin.H{
		"status": gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		},
	})
}

func GetToken(args string) string {
	var strKey string = "a8e6989a4c84tg65"
	args = args + strKey
	sum := md5.Sum([]byte(args))
	hexStr := fmt.Sprintf("%x", sum)
	fmt.Println(hexStr)
	return hexStr
}

func init() {
	fmt.Println("do init")

	var err error
	// 连接mysql
	mysqlConn := fmt.Sprintf("root:123456@tcp(%v:3600)/server_conf?charset=utf8", gMysqlIp)
	gSqlDb, err = sql.Open("mysql", mysqlConn)
	if err != nil {
		fmt.Printf("connect mysql failed! [%s]", err)
		return
	} else {
		fmt.Println("connect mysql ok ")
	}
}

func main() {
	router := gin.Default()

	router.POST("/insertCustomer", func(c *gin.Context) {
		bytes, err := c.GetRawData()
		if err != nil {
			ResponseError(c, err)
			return
		}
		data := string(bytes)

		var customer string
		value := gjson.Get(data, "Customer")
		if value.Index != 0 {
			customer = value.String()
		} else {
			ResponseError(c, fmt.Errorf("Error:Customer can not be empty "))
			return
		}

		var timestamp int64
		value = gjson.Get(data, "Timestamp")
		if value.Index != 0 {
			timestamp = value.Int()
		} else {
			ResponseError(c, fmt.Errorf("Error:Timestamp can not be empty "))
			return
		}

		var token string
		value = gjson.Get(data, "Token")
		if value.Index != 0 {
			token = value.String()
		} else {
			ResponseError(c, fmt.Errorf("Error:Token can not be empty "))
			return
		}

		var params string
		value = gjson.Get(data, "Params")
		if value.Index != 0 {
			params = value.String()
		} else {
			ResponseError(c, fmt.Errorf("Error:Params can not be empty "))
			return
		}

		//fmt.Println(customer, timestamp, params, token)
		args := customer + fmt.Sprintf("%v", timestamp) + params
		tk := GetToken(args)
		if tk != token {
			ResponseError(c, fmt.Errorf("Error:Token is invalid "))
			return
		}

		cmInfo := &CustomerInfo{}
		err = UnmarCustomer(params, cmInfo)
		if err != nil {
			ResponseError(c, err)
			return
		}

		exists, err := ExistsCustomer(cmInfo.Id)
		if err != nil {
			ResponseError(c, err)
			return
		}

		if exists {
			ResponseError(c, fmt.Errorf("Error:Customer [%v] is existed ", cmInfo.Id))
			return
		}

		err = InsertCustomer(cmInfo)
		if err != nil {
			ResponseError(c, err)
			return
		}

		ResponseSuccess(c)
	})

	router.POST("/insertDomain", func(c *gin.Context) {
		bytes, err := c.GetRawData()
		if err != nil {
			ResponseError(c, err)
			return
		}
		data := string(bytes)

		var customer string
		value := gjson.Get(data, "Customer")
		if value.Index != 0 {
			customer = value.String()
		} else {
			ResponseError(c, fmt.Errorf("Error:Customer can not be empty "))
			return
		}

		var timestamp int64
		value = gjson.Get(data, "Timestamp")
		if value.Index != 0 {
			timestamp = value.Int()
		} else {
			ResponseError(c, fmt.Errorf("Error:Timestamp can not be empty "))
			return
		}

		var token string
		value = gjson.Get(data, "Token")
		if value.Index != 0 {
			token = value.String()
		} else {
			ResponseError(c, fmt.Errorf("Error:Token can not be empty "))
			return
		}

		var params string
		value = gjson.Get(data, "Params")
		if value.Index != 0 {
			params = value.String()
		} else {
			ResponseError(c, fmt.Errorf("Error:Params can not be empty "))
			return
		}

		//fmt.Println(customer, timestamp, params, token)
		args := customer + fmt.Sprintf("%v", timestamp) + params
		tk := GetToken(args)
		if tk != token {
			ResponseError(c, fmt.Errorf("Error:Token is invalid "))
			return
		}

		dm := &DomainInfo{}
		err = UnmarDomain(params, dm)
		//err := json.Unmarshal([]byte(param), &dm)
		//if err != nil {
		// ResponseError(c, err)
		// return
		//}
		if err != nil {
			ResponseError(c, err)
			return
		}
		fmt.Println("dm=", dm)

		exists, err := ExistsDomain(dm.Id)
		if err != nil {
			ResponseError(c, err)
			return
		}

		if exists {
			ResponseError(c, fmt.Errorf("Error:Domain [%v] is existed ", dm.Id))
			return
		}

		err = InsertDomain(dm)
		if err != nil {
			ResponseError(c, err)
			return
		}

		ResponseSuccess(c)
	})

	router.POST("/updateCustomer", func(c *gin.Context) {
		bytes, err := c.GetRawData()
		if err != nil {
			ResponseError(c, err)
			return
		}
		data := string(bytes)

		var customer string
		value := gjson.Get(data, "Customer")
		if value.Index != 0 {
			customer = value.String()
		} else {
			ResponseError(c, fmt.Errorf("Error:Customer can not be empty "))
			return
		}

		var timestamp int64
		value = gjson.Get(data, "Timestamp")
		if value.Index != 0 {
			timestamp = value.Int()
		} else {
			ResponseError(c, fmt.Errorf("Error:Timestamp can not be empty "))
			return
		}

		var token string
		value = gjson.Get(data, "Token")
		if value.Index != 0 {
			token = value.String()
		} else {
			ResponseError(c, fmt.Errorf("Error:Token can not be empty "))
			return
		}

		var params string
		value = gjson.Get(data, "Params")
		if value.Index != 0 {
			params = value.String()
		} else {
			ResponseError(c, fmt.Errorf("Error:Params can not be empty "))
			return
		}

		//fmt.Println(customer, timestamp, params, token)
		args := customer + fmt.Sprintf("%v", timestamp) + params
		tk := GetToken(args)
		if tk != token {
			ResponseError(c, fmt.Errorf("Error:Token is invalid "))
			return
		}

		cm := &CustomerInfo{}
		err = UnmarCustomer(params, cm)
		if err != nil {
			ResponseError(c, err)
			return
		}

		err = UpdateCustomer(cm)
		if err != nil {
			ResponseError(c, err)
			return
		}

		ResponseSuccess(c)
	})

	router.POST("/updateDomain", func(c *gin.Context) {
		bytes, err := c.GetRawData()
		if err != nil {
			ResponseError(c, err)
			return
		}
		data := string(bytes)

		var customer string
		value := gjson.Get(data, "Customer")
		if value.Index != 0 {
			customer = value.String()
		} else {
			ResponseError(c, fmt.Errorf("Error:Customer can not be empty "))
			return
		}

		var timestamp int64
		value = gjson.Get(data, "Timestamp")
		if value.Index != 0 {
			timestamp = value.Int()
		} else {
			ResponseError(c, fmt.Errorf("Error:Timestamp can not be empty "))
			return
		}

		var token string
		value = gjson.Get(data, "Token")
		if value.Index != 0 {
			token = value.String()
		} else {
			ResponseError(c, fmt.Errorf("Error:Token can not be empty "))
			return
		}

		var params string
		value = gjson.Get(data, "Params")
		if value.Index != 0 {
			params = value.String()
		} else {
			ResponseError(c, fmt.Errorf("Error:Params can not be empty "))
			return
		}

		//fmt.Println(customer, timestamp, params, token)
		args := customer + fmt.Sprintf("%v", timestamp) + params
		tk := GetToken(args)
		if tk != token {
			ResponseError(c, fmt.Errorf("Error:Token is invalid "))
			return
		}

		dm := &DomainInfo{}
		err = UnmarDomain(params, dm)
		if err != nil {
			ResponseError(c, err)
			return
		}
		fmt.Println("dm=", dm)

		//err := json.Unmarshal([]byte(param), &dm)
		//if err != nil {
		// ResponseError(c, err)
		// return
		//}

		err = UpdateDomain(dm)
		if err != nil {
			ResponseError(c, err)
			return
		}

		ResponseSuccess(c)
	})

	router.Run(":8000")
	defer gSqlDb.Close()
}
