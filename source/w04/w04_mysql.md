## Go 操作MySQL

---

### Docker准备MySQL环境

```bash
# 查看 mysql 库
docker search mysql 

# 普通docker 使用
docker pull mysql
# M1 芯片使用
docker pull mysql/mysql-server

# 通过镜像启动容器
docker run -itd --name mysql-test -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 [镜像名称或ID]

----
参数说明：
  -p 3306:3306 ：映射容器服务的 3306 端口到宿主机的 3306 端口
   外部主机可以直接通过 宿主机ip:3306 访问到 MySQL 的服务。
  - MYSQL_ROOT_PASSWORD=123456：设置 MySQL 服务 root 用户的密码。
  
  
# 进入mysql 容器
docker exec -it mysql-test bash

# 登录mysql
mysql -u root -p
```

### 准备数据库和表

#### 准备数据库

```mysql
# 创建数据库 userinfo
mysql> create database wxw_test;
Query OK, 1 row affected (0.04 sec)

# 查看存在的数据库
mysql> show databases;
+--------------------+
| Database           |
+--------------------+
| information_schema |
| mysql              |
| performance_schema |
| sys                |
| wxw_test           |
+--------------------+
5 rows in set (0.00 sec)
```

可以看到 wxw_test 数据库已经成功创建。接下来我们在 wxw_test 数据库创建两张表：userinfo 和 userdetail。首先切换到 wxw_test 数据库。我们先看看当前正在使用的数据库：

```mysql
## 查看正在使用的数据库
mysql> select database();
+------------+
| database() |
+------------+
| user_info   |
+------------+
1 row in set (0.00 sec)

# 切换数据库
mysql> use wxw_test;
Database changed
```

如上，我们就成功切换到 test 数据库中了，接下来我们创建两张表啦：

#### 准备表

```bash
## 创建 userinfo 表
CREATE TABLE `user_info` (
         `uid` INT(10) NOT NULL AUTO_INCREMENT,
         `username` VARCHAR(64) NULL DEFAULT NULL,
         `company` VARCHAR(64) NULL DEFAULT NULL,
         `createtime` DATE NULL DEFAULT NULL,
          PRIMARY KEY (`uid`)
    );
    
## 查看表是否创建成功
mysql> show tables;
+--------------------+
| Tables_in_wxw_test |
+--------------------+
| userinfo           |
+--------------------+

## 查看建表的表结构
mysql> desc use_rinfo;
+------------+-------------+------+-----+---------+----------------+
| Field      | Type        | Null | Key | Default | Extra          |
+------------+-------------+------+-----+---------+----------------+
| uid        | int         | NO   | PRI | NULL    | auto_increment |
| username   | varchar(64) | YES  |     | NULL    |                |
| company    | varchar(64) | YES  |     | NULL    |                |
| createtime | date        | YES  |     | NULL    |                |
+------------+-------------+------+-----+---------+----------------+
```

#### 增加数据

```mysql
## 增加一条记录
mysql> insert into user_info values('1','wxw','jd','2021-09-29');
Query OK, 1 row affected (0.01 sec)

## 查看这条记录
mysql> select * from user_info;
+-----+----------+---------+------------+
| uid | username | company | createtime |
+-----+----------+---------+------------+
|   1 | wxw      | jd      | 2021-09-29 |
+-----+----------+---------+------------+
```

上面就是 Mysql 数据表最基本的操作了，还有很多要学习，其他的话大家 Google 一下就好了。接下来我们讲述 Go 中如何操作 Mysql 数据库。

### 基于go-sql-driver操作 MySQL

Go没有内置的驱动支持任何的数据库，但是Go定义了database/sql接口，用户可以基于驱动接口开发相应数据库的驱动。这一节来尝试与喜爱 mysql 的驱动

```shell
## 获取 sql驱动的依赖
go get github.com/go-sql-driver/mysql
```

安装好 mysql 驱动之后，我们按照上述的步骤在终端创建：userinfo 。接下来看看 Go 官方提供的示例代码：

```go
/*
@Time: 2021/10/5 20:48
@Author: wxw
@File: demo
*/
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// 基于 sql-driver 操作数据库
func main() {
	db, err := sql.Open("mysql", "root:123456@/wxw_test?charset=utf8")
	checkErr(err)

	//2. 插入数据
	stmt, err := db.Prepare("INSERT user_info SET username=?,company=?,createtime=?")
	checkErr(err)
	res, err := stmt.Exec("w02", "WeChat", "2012-12-09")
	checkErr(err)
	// 插入的Id
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Printf("插入的Id = %d\n", id)

	//3. 更新数据
	stmt, err = db.Prepare("update user_info set username=? where uid=?")
	checkErr(err)
	res, err = stmt.Exec("w02", id)
	checkErr(err)
	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)

	//1. 查询数据
	rows, err := db.Query("SELECT * FROM user_info")
	checkErr(err)
	for rows.Next() {
		var uid int
		var username string
		var company string
		var createtime string
		err = rows.Scan(&uid, &username, &company, &createtime)
		checkErr(err)
		fmt.Println(uid,username,company,createtime)
	}

	//删除数据
	stmt, err = db.Prepare("delete from user_info where uid=?")
	checkErr(err)
	res, err = stmt.Exec(id)
	checkErr(err)
	affect, err = res.RowsAffected()
	checkErr(err)
	fmt.Printf("影响行数: %d\n",affect)

	// 关闭连接
	db.Close()

}

// 存在 error 进行 painc
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
```

打印结果：

```bash
插入的Id = 4
0
1 wxw jd 2021-09-29
2 w01 WeChat 2012-12-09
3 w01 WeChat 2012-12-09
4 w02 WeChat 2012-12-09
影响行数: 1
```

### 基于beego 操作MySQL

ORM 是一种对象关系的映射。接下来体验一下 Beego 中内部支持的 ORM。我们还是先来安装 beego 框架。

- beego 在线文档：
- beego github地址：https://github.com/astaxie/beego

> 安装beego依赖

```bash
## 获取远程依赖
go get github.com/astaxie/beego
```

> 基于beego的数据操作案例

```go
package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // 导入数据库驱动
)

// UserInfo model struct
type UserInfo struct {
	Uid        int64 `orm:"colunm(uid);pk"` // 设置主键
	Username   string
	Company    string
	Createtime string `default:"1994-01-01"` // 设置默认值
}

func init() {
	// 设置默认数据库
	orm.RegisterDataBase("default", "mysql", "root:123456@/wxw_test?charset=utf8", 30)
	// 注册定义的 model 可以同时注册多个 model
	orm.RegisterModel(new(UserInfo))
	//orm.RegisterModel(new(User), new(Profile), new(Post))

	// 如果不存在就创建表
	orm.RunSyncdb("default", false, true)
}

func main() {
	fmt.Println("beego开始操作...")
	// 基本的赋值
	user := UserInfo{Username: "slene", Createtime: "2021-08-12"}

	// 打印记录
	fmt.Println(user.Uid, user.Username, user.Company, user.Createtime)

	// 开启一个连接
	o := orm.NewOrm()

	// 插入表
	id, err := o.Insert(&user)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)

	// 更新表
	//user.Username = "w4"
	//num, err := o.Update(&user)
	//fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	// 读取 one
	u := UserInfo{Uid: user.Uid}
	err = o.Read(&u)
	fmt.Printf("ERR: %v\n", err)
	//
	//	// 删除表
	//	//num, err = o.Delete(&u)
	//fmt.Printf("NUM: %d, ERR: %v\n", num, err)
}
```





























