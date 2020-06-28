Sqlxx 

（目前属于开发完善阶段，尚未发布正式版本）

项目里可以通过下面获取该仓库

    go get github.com/dk-sirius/sqlxx

Sqlxx 配置

方案一 使用yaml 文件配置

    postgres = &confpostgres.PostgresSql{}
    db      = postgres.ConnYAML("db.yml")

文件内容如下

    host: 127.0.0.1
    user: your_db_user
    password: your_db_password
    port: your_db_port
    name: your_db_name

方案二 内容配置

    postgre = &confpostgres.PostgresSql{
      // 数据库地址
    	// Host 
    	// 数据库名称
    	// Name
    	// 用户名称
    	// User
    	// 用户密码
    	// Password
    	// 端口号
    	// Port
    }

