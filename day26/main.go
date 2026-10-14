package main

import (
	"fmt"
	"go-learning/day26/handler"
	"go-learning/day26/repository"
	"go-learning/day26/service"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("读取.env文件失败:", err)
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user,
		password,
		host,
		port,
		dbName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("连接MYSQL失败:", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("获取数据库连接对象失败:", err)
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Fatal("MYSQL连通性检查失败:", err)
	}
	fmt.Println("MYSQL连接成功!")

	userRepo := repository.NewUserRepository(db)
	fmt.Println("UserRepository初始化成功!")

	userService := service.NewUserService(userRepo)
	fmt.Println("UserService 初始化成功!")

	userHandler := handler.NewUserHandler(userService)
	fmt.Println("UserHandler初始化成功!")

	r := gin.Default()
	r.GET("/users/:id", userHandler.GetUserByID)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Gin 服务启动失败:", err)
	}
}
