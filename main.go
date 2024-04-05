package main

import (
	"database/sql"
	"fmt"
	device_handler "iot/internal/handler/rest/device"
	sensor_handler "iot/internal/handler/rest/sensor"
	device_repo "iot/internal/repo/device"
	sensor_repo "iot/internal/repo/sensor"
	device_uc "iot/internal/usecase/device"
	sensor_uc "iot/internal/usecase/sensor"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// SETUP DATABASE
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
	)

	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		panic("failed to open database")
	}
	defer sqlDB.Close()

	dbName := os.Getenv("DB_NAME")

	_, err = sqlDB.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)
	if err != nil {
		panic("failed to initiate database")
	}

	_, err = sqlDB.Exec("USE " + dbName)
	if err != nil {
		panic("failed to use database")
	}

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// REPO INITIALIZATION
	deviceRepo := device_repo.New(gormDB)
	sensorRepo := sensor_repo.New(gormDB)

	// USECASE INITIALIZATION
	deviceUsecase := device_uc.New(deviceRepo)
	sensorUsecase := sensor_uc.New(sensorRepo)

	// HANDLER INITIALIZATION
	deviceHandler := device_handler.New(deviceUsecase)
	sensorHandler := sensor_handler.New(sensorUsecase)

	// SETUP ROUTER
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Welcome to IoT project!",
		})
	})

	router.GET("/devices", deviceHandler.GetListDevices)
	router.GET("/devices/:id", deviceHandler.GetDevice)
	router.POST("/devices", deviceHandler.CreateDevice)
	router.PUT("/devices/:id", deviceHandler.UpdateDevice)
	router.DELETE("/devices/:id", deviceHandler.DeleteDevice)

	router.GET("/sensors", sensorHandler.GetListSensors)
	router.GET("/sensors/:id", sensorHandler.GetSensor)
	router.POST("/sensors", sensorHandler.CreateSensor)
	router.PUT("/sensors/:id", sensorHandler.UpdateSensor)
	router.DELETE("/sensors/:id", sensorHandler.DeleteSensor)

	serverPort := fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))
	log.Printf("server listening at %s", serverPort)

	http.ListenAndServe(serverPort, router)
}
