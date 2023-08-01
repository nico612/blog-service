package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nico612/blog-service/global"
	"github.com/nico612/blog-service/internal/model"
	"github.com/nico612/blog-service/internal/routers"
	"github.com/nico612/blog-service/pkg/logger"
	"github.com/nico612/blog-service/pkg/setting"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"time"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}

	setupLogger()
}

func main() {
	gin.SetMode(global.ServerSetting.RunModel)
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		return
	}
}

// 读取配置文件
func setupSetting() error {
	// 初始化配置读取Setting
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}

	// 读取服务Setting
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

// db配置
func setupDBEngine() error {
	var err error
	// 注意不要写成了global.DBEngine, err := model.NewDBEngine(global.DatabaseSetting)
	// := 会重新声明并创建了左侧的新局部变量，因此在其它包中调用 global.DBEngine 变量时，它仍然是 nil，
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	return err
}
func setupLogger() {
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName,
		MaxSize:   600,  //	最大占用空间600M
		MaxAge:    10,   // 日志文件最大生存周期为10天
		LocalTime: true, //	日志文件名的时间格式为本地时间
	}, "", log.LstdFlags).WithCaller(2)
}
