> Go 基础中间件

---



### 1. 日志组件

> 概览

- `zerolog` ：[https://github.com/rs/zerolog](https://github.com/rs/zerolog)  

### 2. ` MySQL` 

#### 2.1 docker准备` mysql` 

```bash
# 拉取mysql容器
docker pull mysql

# 通过镜像启动一个mysql容器
docker run -itd --name mysql-test -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 mysql

# 进入mysql-test 容器
docker exec -it mysql-test bash

# 登录mysql
mysql -uroot -p
```



相关文章

1. [Go 中的 MySql 和 ORM](https://studygolang.com/articles/6992) 

### 3. `orm` 组件

> 概览

- 原生的方式：
- `gorm`：[https://gorm.io/zh_CN/docs/](https://gorm.io/zh_CN/docs/)  

