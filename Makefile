#需要安装的一些仓库
# migrate
install_migrate:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# migrate的mysql驱动
install_migrate_mysql:
	go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
#sqlc
install_sqlc:
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

# mysql数据库
lanch_mysql:
	docker run --name mysql-url -e MYSQL_ROOT_PASSWORD=123456 -e MYSQL_DATABASE=urldb -e MYSQL_USER=lang -e MYSQL_PASSWORD=123456 -p 3306:3306 -d mysql:latest

# redis
lanch_redis:
	docker run --name redis -p 6379:6379 -d redis

#本地不需要加密连接
# databaseURL="postgres://lang:123456@192.168.0.37:5432/urldb?sslmode=disable"
databaseURL="mysql://lang:123456@tcp(localhost:3306)/urldb"
# 数据库迁移
migrate_up:
	migrate -path="./database/migrate" -database=${databaseURL} up

migrate_down:
	migrate -path="./database/migrate" -database=${databaseURL} drop -f