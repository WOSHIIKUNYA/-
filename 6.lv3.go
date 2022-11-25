// 时间不够了我也没做留言的回复，下次做吧
// 我的代码很多地方有缺陷，懒得改了
// MySQL的表名叫userbook有Name,Possword,question,Protect(密保信息),message5行
package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type people struct {
	Possword string `json:"Possword"`
	Name     string `json:"Name"`
}

// 用于服务登录和注册的
type kk struct {
	Message string `json:"Message"`
}

// 服务与消息
var user_book []people

// 用于服务登录和注册的
var Bool bool = false

// 用于登录
type number struct {
	Name     string `json:"Name"`
	Question string `json:"Question"`
	Protect  string `json:"Protect"`
}

// 密保的
type message struct {
	Name    string `json:"Name"`
	Massage string `json:"Massage"`
}

// 发消息
var SEEK []number

// 密保
var rf int = 0

// 留言
// 统计登录的人在切片的序号，这里我觉得应该用数组并且用并发编程实现多用户，下次一定
type rrq struct {
	Name string `json:"Name"`
}

// 查看留言
var zz rrq

// 查看留言
func liu(m *gin.Context) {
	//rf是用户登录后rf会自动变为10的，相当与一个判断
	if rf != 10 {
		m.String(200, "请先登录")
		return
	}
	var use message
	m.ShouldBind(&use)
	a, err := sql.Open("mysql", "root:w2002101500f@tcp(127.0.0.1:3306)/mine")
	if err != nil {
		fmt.Println(err)
	}
	err = a.Ping()
	if err != nil {
		fmt.Println(err)
	}
	a.Exec("update userbook set  Message=? where Name=? ", use.Massage, use.Name)
}

// 留言
func mes(m *gin.Context) {
	if rf != 10 {
		m.String(200, "请先登录")
		return
	}

	//rf是用户登录后rf会自动变为10的，相当与一个判断
	a, err := sql.Open("mysql", "root:w2002101500f@tcp(127.0.0.1:3306)/mine")
	if err != nil {
		fmt.Println(err)
	}
	rt := a.QueryRow("select Message from userbook where name=?", zz.Name)
	var mesr kk
	rt.Scan(&mesr.Message)
	m.String(200, mesr.Message)
}

// 查流言
func seek(m *gin.Context) {
	a, err := sql.Open("mysql", "root:w2002101500f@tcp(127.0.0.1:3306)/mine")
	x, err1 := a.Query("select Name,qusetion,protect from userbook")
	if err1 != nil {
		fmt.Println(err)
	}
	for x.Next() {
		var cc number
		x.Scan(&cc.Name, &cc.Question, &cc.Protect)
		SEEK = append(SEEK, cc)
	}

	var use number
	m.ShouldBind(&use)
	var cc int
	for xx := 0; xx < len(SEEK); xx++ {
		if SEEK[xx].Name == use.Name {
			m.String(200, "有该用户信息")
			cc = xx
			break
		}
	}
	if SEEK[cc].Protect == use.Protect {
		rt := a.QueryRow("select Possword from userbook where name=?", use.Name)
		var ff string
		rt.Scan(&ff)
		m.String(200, ff)
	} else {
		m.String(200, "错误密保")
	}

}

// 找密码函数
func protect(m *gin.Context) {
	var ww number
	a, err := sql.Open("mysql", "root:w2002101500f@tcp(127.0.0.1:3306)/mine")
	err = a.Ping()
	if err != nil {
		fmt.Println(err)
	}
	m.ShouldBind(&ww)
	a.Exec("update userbook set  qusetion=? where Name=? ", ww.Question, ww.Name)
	a.Exec("update userbook set  protect=? where name=?", ww.Protect, ww.Name)
	m.JSON(200, "已完成")
}

// 设置密保函数
func mine(m *gin.Context) {
	a, err := sql.Open("mysql", "root:w2002101500f@tcp(127.0.0.1:3306)/mine")
	if err != nil {
		fmt.Println(err)
	}
	err = a.Ping()
	if err != nil {
		fmt.Println(err)
	}
	var use people
	m.ShouldBind(&use)
	a.Exec("insert into userbook(Name,Possword)values (?,?)", use.Name, use.Possword)
	m.JSON(200, use)
}

// 注册函数
func logn_in(m *gin.Context) {
	Bool = false
	a, err := sql.Open("mysql", "root:w2002101500f@tcp(127.0.0.1:3306)/mine")
	if err != nil {
		fmt.Println(err)
	}
	b, err1 := a.Query("select Name,Possword from userbook")
	if err1 != nil {
		fmt.Println(err)
	}
	for b.Next() {
		var use people
		b.Scan(&use.Name, &use.Possword)
		user_book = append(user_book, use)
	}
	var use people
	m.ShouldBind(&use)
	m.String(200, use.Name)
	m.String(200, use.Possword)
	for xr := 0; xr < len(user_book); xr++ {
		if (user_book[xr].Name == use.Name) && (user_book[xr].Possword == use.Possword) {
			Bool = true
			m.String(200, "已登录")
			rf = 10
			zz.Name = use.Name
			break
		}
	}
	if Bool == false {
		m.String(200, "密码或用户名错误")
	}

}

// 登录函数

func main() {
	r := gin.Default()
	r.POST("cin", mine)
	r.POST("Get", logn_in)
	r.POST("protect", protect)
	r.POST("Seek", seek)
	r.POST("LiuYan", liu)
	r.POST("look", mes)
	r.Run()
}

//注释还是写了点，放过鼠鼠把
