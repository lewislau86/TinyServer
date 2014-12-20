///////////////////////////////////////
// TinyServer
// Golang Server Framework
//  Lewis	( lewislau86@gmail.com )
///////////////////////////////////////
package TinyServer

//////////////////////////////////////

import (
	"database/sql"
	"fmt"
	_ "github.com/Go-SQL-Driver/MySQL"
	//"os"
)

///////////////////////////////////////

type TinyDatabase struct {
	openString string
	db         *sql.DB
}

//////////////////////////////////////

func (mysql *TinyDatabase) Open(dataSource string) {
	var err error
	mysql.db, err = sql.Open("mysql", dataSource)
	if err != nil {
		fmt.Println("open sql error")
		fmt.Println(err)
	} else {
		fmt.Println("Open Databse Succsess")
	}
}

///////////////////////////////////////

func (mysql TinyDatabase) getUserInfo(name string) (UserUid int, UserName string, UserHash string, UserPwd string, UserFlag int, Expiry string) {
	// 检查登陆名、密码、机器id
	var (
		uid  int
		user string
		hash string
		pwd  string
		flag int
		data string
	)

	sql := fmt.Sprintf("select * from userinfo where user='%s' ", name)
	fmt.Println(sql)
	rows, err := mysql.db.Query(sql)
	if err != nil {
		fmt.Print(err)
	}
	for rows.Next() {
		err := rows.Scan(&uid, &user, &hash, &pwd, &flag, &data)
		if err != nil {
			fmt.Println(err)
		}
		return uid, user, hash, pwd, flag, data
	}
	return uid, user, hash, pwd, flag, data
}

///////////////////////////////////////

func (mysql TinyDatabase) CheckLogin(name string, passwd string) bool {
	//mysql.Open("root:lewislau@/test")
	uid, user, pwd, hash, flag, date := mysql.getUserInfo(name)

	if uid != 0 && user != "" && pwd != "" {
		if user == name && passwd == pwd {
			return true
		}
	}
	fmt.Println(uid)
	fmt.Println(user)
	fmt.Println(hash)
	fmt.Println(pwd)
	fmt.Println(flag)
	fmt.Println(date)
	mysql.db.Close()
	return false
}

///////////////////////////////////////
/*
func main() {
	var test TinyDatabse
	arg_num := len(os.Args)
	if arg_num != 3 {
		fmt.Println(arg_num)
		fmt.Println("参数错误")
		return
	}

	ok := test.CheckLogin(os.Args[1], os.Args[2])
	if ok {
		fmt.Println("Login ok")
	} else {
		fmt.Println("Login error")
	}

}
*/
///////////////////////////////////////
